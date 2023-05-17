package state

import "fmt"
import "main/api"

type luaStack struct {
	slots 		[]luaValue
	top 		int
	prev		*luaStack
	closure		*closure
	varargs		[]luaValue
	pc			int
}

func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots : make([]luaValue, size),
		top : 0,
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
	if idx >= 0 {
		return idx
	}

	return idx + self.top + 1
}

func (self *luaStack) isValid(idx int) bool {
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		return true
	}
	return false
}

func (self *luaStack) get(idx int) luaValue {
	absIdx := self.absIndex(idx)
	// fmt.Println("stack get absInx", idx, absIdx, self.top)
	if absIdx > 0 && absIdx <= self.top {
		return self.slots[absIdx -1]
	}
	return nil
}

func (self *luaStack) set(idx int, val luaValue) {
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		self.slots[absIdx -1] = val
		return
	}
	fmt.Println("invalid index", absIdx, self.top)
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
	fmt.Println("pushN", vals)
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

func (self *luaStack)printStack(i int)  {
	fmt.Printf("[%d] size[%d] top[%d] stack", i, len(self.slots), self.top)
	for i := 0; i < self.top; i++ {
		val := self.slots[i]
	   	t :=  typeOf(val)
	   	switch t {
	   	case api.LUA_TBOOLEAN:      	fmt.Printf("[%t]", convertToBoolean(val))
	   	case api.LUA_TNUMBER:
			g,_ := convertToFloat(val)    	
			fmt.Printf("[%g]", g)
	   	case api.LUA_TSTRING:       	fmt.Printf("[%q]", _toString(val))
	   	default:                		fmt.Printf("[%s]", _typeName(t))
	   	}
	}
	fmt.Println("\n")

	if self.prev != nil {
		i += 1
		self.prev.printStack(i)
	}
}