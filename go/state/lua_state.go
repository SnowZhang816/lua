package state

import "main/api"
import "main/cLog"

type luaState struct {
	registry 		*luaTable
	stack 			*luaStack
}

func New() *luaState {
	cLog.Println("New luaState")
	registry := newLuaTable(0,0)
	registry.put(api.LUA_RIDX_GLOBALS, newLuaTable(0,0))		//全局环境

	ls :=  &luaState{registry : registry,}

	ls.registry.printTable()

	ls.pushLuaStack(newLuaStack(api.LUA_MINSTACK, ls))

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

func (self *luaState) printStack() {
	self.stack.printStack(1)
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