package state

// import "main/binChunk"

type luaState struct {
	stack 	*luaStack
}

func New() *luaState {
	return &luaState{
		stack: 	newLuaStack(20),
	}
}

func (self *luaState) pushLuaStack(stack *luaStack) {
	stack.prev = self.stack
	self.stack = stack
}

func (self *luaState) popLuaStack() {
	stack := self.stack
	self.stack = stack.prev
	stack.prev = nil
}

func (self *luaState) printStack() {
	self.stack.printStack(1)
}