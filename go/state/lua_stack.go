package state

import "main/api"
import "main/cLog"
import "fmt"

type luaStack struct {
	slots 		[]luaValue
	top 		int
	prev		*luaStack
	closure		*closure
	varargs		[]luaValue
	pc			int
	state 		*luaState
	openuvs		map[int]*upValue
}

func newLuaStack(size int, state *luaState) *luaStack {
	return &luaStack{
		slots : make([]luaValue, size),
		top : 0,
		state : state,
	}
}

func (self *luaStack) check(n int) {
	free := len(self.slots) - self.top
	for i := free; i < n; i++ {
		self.slots = append(self.slots, nil)
	}
}

func (self *luaStack) push(val luaValue) {
	if self.top == len(self.slots) {
		panic("stack overflow!")
	}

	self.slots[self.top] = val
	self.top++
}

func (self *luaStack) pop() luaValue {
	if self.top < 1 {
		panic("stack underflow!")
	}

	self.top--
	val := self.slots[self.top]
	self.slots[self.top] = nil
	return val
}

func (self *luaStack) absIndex(idx int) int {
	if idx <= api.LUA_REGISTRY_INDEX {
		return idx
	}
	if idx >= 0 {
		return idx
	}

	return idx + self.top + 1
}

func (self *luaStack) isValid(idx int) bool {
	if idx < api.LUA_REGISTRY_INDEX {
		uvIdx := api.LUA_REGISTRY_INDEX - idx - 1
		c := self.closure
		return c != nil && uvIdx < len(c.upValues)
	}
	if idx == api.LUA_REGISTRY_INDEX {
		return true
	}
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		return true
	}
	return false
}

func (self *luaStack) get(idx int) luaValue {
	// cLog.Println("luaStack get", idx)
	if idx < api.LUA_REGISTRY_INDEX {
		uvIdx := api.LUA_REGISTRY_INDEX - idx - 1
		c := self.closure
		if c == nil || uvIdx >= len(c.upValues) {
			return nil
		}
		// cLog.Println("luaStack get uvIdx", uvIdx)
		return *(c.upValues[uvIdx].val)
	}
	if idx == api.LUA_REGISTRY_INDEX {
		return self.state.registry
	}
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		return self.slots[absIdx -1]
	}
	return nil
}

func (self *luaStack) set(idx int, val luaValue) {
	if idx < api.LUA_REGISTRY_INDEX {
		uvIdx := api.LUA_REGISTRY_INDEX - idx - 1
		c := self.closure
		if c != nil && uvIdx < len(c.upValues) {
			*(c.upValues[uvIdx].val) = val
		}
		return
	}

	if idx == api.LUA_REGISTRY_INDEX {
		self.state.registry = val.(*luaTable)
		return
	}
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		self.slots[absIdx -1] = val
		return
	}
	cLog.Println("invalid index", absIdx, self.top)
	panic("invalid index")
}

func (self *luaStack) reverse(from, to int) {
	slots := self.slots

	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}

func (self *luaStack) pushN(vals []luaValue, n int) {
	cLog.Println("pushN", vals, n)
	nVals := len(vals)
	if n < 0 {
		n = nVals
	}
	for i := 0; i < n; i++ {
		if i < nVals {
			self.push(vals[i])
		} else {
			self.push(nil)
		}
	}
}

func (self *luaStack) popN(n int) ([]luaValue) {
	vals := make([]luaValue, n)
	for i := n - 1; i >= 0; i-- {
		vals[i] = self.pop()
	}
	return vals
}

func _toString(val luaValue) string {
	switch x := val.(type) {
	case string:				return x
	case int64, float64:		
		s := fmt.Sprintf("%v", x)
		return s
	default:					return ""
	}
}

func _typeName(val api.LuaType) string {
	switch val {
	case LUA_TNONE:					return "no value"
	case LUA_TNIL:					return "nil"
	case LUA_TBOOLEAN:				return "boolean"
	case LUA_TLIGHTUSRDATA:			return "light userdata"
	case LUA_TNUMBER:				return "number"
	case LUA_TSTRING:				return "string"
	case LUA_TTABLE:				return "table"
	case LUA_TFUNCTION:				return "function"
	case LUA_TUSERDATA:				return "userdata"
	case LUA_TTHREAD:				return "thread"
	default:						return ""
	}
}

func _printLuaValue(val luaValue) {
	t := typeOf(val)
		switch t {
		case api.LUA_TBOOLEAN:      	
			cLog.Printf("[%t]", convertToBoolean(val))
		case api.LUA_TNUMBER:
			g,_ := convertToFloat(val)    	
			cLog.Printf("[%g]", g)
		case api.LUA_TSTRING:       	
			cLog.Printf("[%q]", _toString(val))
		case LUA_TFUNCTION:
			closure := val.(*closure)
			if closure.proto != nil {
				cLog.Print("[LF(")
				cLog.Print(&closure.proto)
				cLog.Print(")]")
			} else {
				cLog.Print("[GF(")
				cLog.Print(closure.goFunc)
				cLog.Print(")]")
			}			
		default:                		
			cLog.Printf("[%s]", _typeName(t))
	}
}

func (self *luaStack)printStack(i int, loop bool)  {
	cLog.Printf("[%d] size[%d] top[%d] stack", i, len(self.slots), self.top)
	for i := 0; i < self.top; i++ {
		val := self.slots[i]
		_printLuaValue(val)
	}
	cLog.Println("\n")

	if loop {
		if self.prev != nil {
			i += 1
			self.prev.printStack(i, loop)
		}
	}
}

func (self *luaStack)printUpValues() {
	upValuesCount := len(self.closure.upValues)
	cLog.Printf("upValues: size[%d] values", upValuesCount)
	for i := 0; i < upValuesCount; i++ {
		val := *(self.closure.upValues[i].val)
		_printLuaValue(val)
	}
	cLog.Println()

	openuvsCount := 0
	if self.openuvs != nil {
		openuvsCount = len(self.openuvs)
	}
	cLog.Printf("openuvs: size[%d] ", openuvsCount)
	for i, openuv := range self.openuvs {
		val := *openuv.val
		cLog.Printf("[%d]-", i)
		_printLuaValue(val)
	}
	cLog.Println()
}