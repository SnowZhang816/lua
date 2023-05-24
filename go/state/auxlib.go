package state

import "main/api"
import "main/stdlib"
import "main/cLog"
import "io/ioutil"

func (self *luaState) Len2(idx int) int64 {
	self.Len(idx)
	i, isNum := self.ToIntegerX(-1)
	if !isNum {
		self.Error2("object length is not a integer!")
	}
	self.Pop(1)
	return i
}

func (self *luaState) Error2(fmt string, a ...interface{}) int {
	self.PushFString(fmt, a...)
	return self.Error()
}

func (self *luaState) CheckStack2(sz int, msg string) {
	if !self.CheckStack(sz) {
		if msg != "" {
			self.Error2("stack overflow (%s)", msg)
		} else {
			self.Error2("stack overflow")
		}
	}
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
		"_G": 			stdlib.OpenBaseLib,
		"math":			stdlib.OpenMathLib,
		"table":		stdlib.OpenTableLib,
		"string":		stdlib.OpenStringLib,
		"utf8":			stdlib.OpenUtf8Lib,
		"os":			stdlib.OpenOsLib,
		"package":		stdlib.OpenPackageLib,
		"coroutine":	stdlib.OpenCoroutineLib,
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
		self.PushGoFunction(openf)
		self.PushString(modename)
		self.Call(1,1)
		self.PushValue(-1)
		self.SetField(-3, modename)
	}
	cLog.Println("RequireF End", modename)

	self.Remove(-2)
	if glb {
		self.PushValue(-1)
		self.SetGlobal(modename)
	}
}

func (self *luaState) SetFuncs(l api.FuncReg, nup int) {
	cLog.Println("SetFuncs", nup, l)
	self.CheckStack(nup)	
	/* package global package*/
	for name, fun := range l {
		cLog.Println("SetFuncs", name)
		for i := 0; i < nup; i++ {
			self.PushValue(-nup)	/* package global package package*/
		}
		self.PushGoClosure(fun, nup)	/* package global package GoFuncion*/
		self.SetField(-(2 + nup), name)
	}
	self.Pop(nup)
}

func (self *luaState) NewLib(l api.FuncReg) {
	self.CreateTable(0, len(l))
	self.SetFuncs(l, 0)
}

func (self *luaState) IsFunction(idx int) bool {
	val := self.stack.get(idx)
	if _, ok := val.(*closure); ok {
		return true
	}
	return false
}

func (self *luaState) LoadFileX(filename, mode string) int {
	if data, err := ioutil.ReadFile(filename); err == nil {
		return self.Load(data, "@" + filename, mode)
	}
	return api.LUA_ERRFILE
}

func (self *luaState) DoString(str string) bool {
	return self.LoadString(str) == api.LUA_OK && self.PCall(0, api.LUA_MULTRET, 0) == api.LUA_OK
}

func (self *luaState) LoadFile(filename string) int {
	return self.LoadFileX(filename, "bt")
}

func (self *luaState) DoFile(filename string) bool {
	return self.LoadFile(filename) == api.LUA_OK && self.PCall(0, api.LUA_MULTRET, 0) == api.LUA_OK
}

func (self *luaState) LoadString(s string) int {
	return self.Load([]byte(s), s, "bt")
}

func (self *luaState) TypeName2(idx int) string {
	return self.TypeName(self.Type(idx))
}


func (self *luaState) GetSubTable(idx int, fname string) bool {
	cLog.Println("GetSubTable", idx, fname)
	if self.GetField(idx, fname) == LUA_TTABLE {
		return true
	}
	self.Pop(1)
	idx = self.stack.absIndex(idx)
	self.NewTable()
	self.PushValue(-1)
	self.SetField(idx, fname)
	return false
}
