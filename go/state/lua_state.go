package state

import "fmt"

type luaState struct {
	stack 	*luaStack
}

func New() *luaState {
	return &luaState{
		stack: 	newLuaStack(20),
	}
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