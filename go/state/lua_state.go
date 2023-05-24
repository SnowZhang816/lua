package state

import "main/api"
import "main/cLog"

type luaState struct {
	registry 		*luaTable
	stack 			*luaStack

	coStatus		int
	coCaller		*luaState
	coChan			chan int
}

func New() *luaState {
	cLog.Println("New luaState")

	ls := &luaState{}

	registry := newLuaTable(8,0)
	registry.put(api.LUA_RIDX_MAINTHREAD, ls)					//
	registry.put(api.LUA_RIDX_GLOBALS, newLuaTable(0,20))		//全局环境

	ls.registry = registry
	ls.pushLuaStack(newLuaStack(api.LUA_MINSTACK, ls))
	ls.registry.printTable()

	return ls
}

func (self *luaState) pushLuaStack(stack *luaStack) {
	cLog.Println("pushLuaStack")
	stack.prev = self.stack
	self.stack = stack
}

func (self *luaState) popLuaStack() {
	cLog.Println("popLuaStack")
	stack := self.stack
	self.stack = stack.prev
	stack.prev = nil
}

func (self *luaState) isMainThread() bool {
	return self.registry.get(api.LUA_RIDX_MAINTHREAD) == self
}

func (self *luaState) printStack(loop bool) {
	self.stack.printStack(1, loop)
}

func (self *luaState)printUpValues() {
	self.stack.printUpValues()
}

func (self *luaState) printRegister (){
	cLog.Print("registry: ")
	self.registry.printTable()
}

func (self *luaState) printLoadedTable (){
	t := self.registry.get("_LOADED")
	if tbl,ok := t.(*luaTable); ok {
		cLog.Print("_LOADED: ")
		tbl.printTable()
		cLog.Println()
	}

}

func (self *luaState) printGlobalTable (){
	t := self.registry.get(api.LUA_RIDX_GLOBALS)
	if tbl,ok := t.(*luaTable); ok {
		cLog.Print("_GLOBALS: ")
		tbl.printTable()
		cLog.Println()
	}
}