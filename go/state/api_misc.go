package state

// import "fmt"
import "main/api"
import "main/stdlib"
import "main/cLog"

func (self *luaState) Len(idx int) {
	val := self.stack.get(idx)
	if s, ok := val.(string); ok {
		self.stack.push(int64(len(s)))
	} else if result, ok := callMetaMethod(val, val, "__len", self); ok { 
		self.stack.push(result)
	} else if t, ok := val.(*luaTable); ok {
		self.stack.push(int64(t.len()))
	} else {
		panic("length error!")
	}
}

func (self *luaState) Len2(idx int) int64 {
	self.Len(idx)
	i, isNum := self.ToIntegerX(-1)
	if !isNum {
		self.Error2("object length is not a integer!")
	}
	self.Pop(1)
	return i
}

func (self *luaState) RawLen(idx int) {
	val := self.stack.get(idx)
	if s, ok := val.(string); ok {
		self.stack.push(int64(len(s)))
	} else if t, ok := val.(*luaTable); ok {
		self.stack.push(int64(t.len()))
	} else {
		panic("length error!")
	}
}

func (self *luaState) Concat(n int) {
	if n == 0 {
		self.stack.push("")
	} else if n >= 2 {
		for i := 1; i < n; i++{
			if self.IsString(-1) && self.IsString(-2) {
				s2 := self.ToString(-1)
				s1 := self.ToString(-2)
				self.stack.pop()
				self.stack.pop()
				self.stack.push(s1 + s2)
				continue
			}

			b := self.stack.pop()
			a := self.stack.pop()
			if result,ok := callMetaMethod(a, b, "__concat", self); ok {
				self.stack.push(result)
				continue
			}

			panic("concatenation error!")
		}
	}
}

func (self *luaState) Next(idx int) bool {
	val := self.stack.get(idx)
	if t, ok := val.(*luaTable); ok {
		key := self.stack.pop()
		if nextKey := t.nextKey(key); nextKey != nil {
			self.stack.push(nextKey)
			self.stack.push(t.get(nextKey))
			return true
		}
		return false
	}

	panic("table expected!")
}

func (self *luaState) Error() int {
	err := self.stack.pop()
	panic(err)
}

func (self *luaState) Error2(fmt string, a ...interface{}) int {
	self.PushFString(fmt, a...)
	return self.Error()
}

func (self *luaState) ArgError(arg int, extraMsg string) int {
	return self.Error2("bad argument #%d (%s)", arg, extraMsg)
}

func (self *luaState) ArgCheck(cond bool, arg int, extraMsg string) {
	if !cond {
		self.ArgError(arg, extraMsg)
	}
}

func (self *luaState) CheckAny(arg int) {
	if self.Type(arg) == LUA_TNONE {
		self.ArgError(arg, "value expected")
	}
}

func (self *luaState) tagError(arg int, t api.LuaType) {
	self.Error2("%s expected, got %s", self.TypeName(t), self.TypeName2(arg))
}

func (self *luaState) CheckType(arg int, t api.LuaType) {
	if self.Type(arg) != t {
		self.tagError(arg, t)
	}
}

func (self *luaState) CheckInteger(arg int) int64 {
	i, ok := self.ToIntegerX(arg)
	if !ok {
		self.tagError(arg, api.LUA_TNUMBER)
	}
	return i
}

func (self *luaState) CheckNumber(arg int) float64 {
	f, ok := self.ToNumberX(arg)
	if !ok {
		self.tagError(arg, api.LUA_TNUMBER)
	}
	return f
}

func (self *luaState) CheckString(arg int) string {
	s, ok := self.ToStringX(arg)
	if !ok {
		self.tagError(arg, api.LUA_TSTRING)
	}
	return s
}

func (self *luaState) OptInteger(arg int, d int64) int64 {
	if self.IsNoneOrNil(arg) {
		return d
	}
	return self.CheckInteger(arg)
}

func (self *luaState) OptNumber(arg int, f float64) float64 {
	if self.IsNoneOrNil(arg) {
		return f
	}
	return self.CheckNumber(arg)
}

func (self *luaState) OptString(arg int, s string) string {
	if self.IsNoneOrNil(arg) {
		return s
	}
	return self.CheckString(arg)
}

func (self *luaState) OpenLibs() {
	libs := map[string]api.GoFunction{
		"_G" : stdlib.OpenBaseLib,
	}

	for name, fun := range libs {
		self.RequireF(name, fun, true)
		self.Pop(1)
	}
}

func (self *luaState) RequireF(modename string, openf api.GoFunction, glb bool) {
	cLog.Println("RequireF", modename)
	self.GetSubTable(api.LUA_REGISTRY_INDEX, "_LOADED")
	self.GetField(-1, modename)			/*LOADED[modename]*/
	if !self.ToBoolean(-1) {
		self.Pop(1)
		self.PushGoFunction(openf, 0)
		self.PushString(modename)
		self.Call(1,1)
		self.PushValue(-1)
		self.SetField(-3, modename)
	}
	cLog.Println("RequireF End")

	self.Remove(-2)
	if glb {
		self.PushValue(-1)
		self.GetGlobal(modename)
	}
}

func (self *luaState) SetFuncs(l api.FuncReg, nup int) {
	cLog.Println("SetFuncs")
	self.CheckStack(nup)

	for name, fun := range l {
		cLog.Println("SetFuncs", name)
		for i := 0; i < nup; i++ {
			self.PushValue(-nup)
		}

		self.PushGoFunction(fun, nup)
		self.SetField(-2, name)
	}
}
