package state

import (
	"main/api"
	"fmt"
)

func (self *luaState) PushNil() {
	self.stack.push(nil)
}

func (self *luaState) PushBoolean(b bool) {
	self.stack.push(b)
}

func (self *luaState) PushInteger(n int64) {
	self.stack.push(n)
}

func (self *luaState) PushNumber(n float64) {
	self.stack.push(n)
}

func (self *luaState) PushString(s string) {
	self.stack.push(s)
}

func (self *luaState) PushFString(fmtStr string, a ...interface{}) {
	s := fmt.Sprintf(fmtStr, a...)
	self.stack.push(s)
}

func (self *luaState) PushGoFunction(f api.GoFunction) {
	gClosure := newGoClosure(f, 0)
	self.stack.push(gClosure)
}

func (self *luaState) PushGoClosure(f api.GoFunction, n int) {
	gClosure := newGoClosure(f, n)
	for i := n; i > 0; i-- {
		val := self.stack.pop()		
		gClosure.upValues[n - 1] = &upValue{&val}
	}
	self.stack.push(gClosure)	
}

func (self *luaState) PushGlobalTable() {
	global := self.registry.get(api.LUA_RIDX_GLOBALS)
	self.stack.push(global)
}

func (self *luaState) PushThread() bool {
	self.stack.push(self)
	return self.isMainThread()
}