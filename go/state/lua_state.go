package state

import "fmt"
import "main/api"

type luaState struct {
	registry 		*luaTable
	stack 			*luaStack
}

func New() *luaState {
	registry := newLuaTable(0,0)
	registry.put(api.LUA_RIDX_GLOBALS, newLuaTable(0,0))		//全局环境

	ls :=  &luaState{registry : registry,}
	ls.pushLuaStack(newLuaStack(api.LUA_MINSTACK, ls))

	return ls
}

func (self *luaState) pushLuaStack(stack *luaStack) {
	fmt.Println("pushLuaStack")
	stack.prev = self.stack
	self.stack = stack
}

func (self *luaState) popLuaStack() {
	fmt.Println("popLuaStack")
	stack := self.stack
	self.stack = stack.prev
	stack.prev = nil
}

func (self *luaState) printStack() {
	self.stack.printStack(1)
}