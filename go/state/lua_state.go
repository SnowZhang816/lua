package state

import "main/binChunk"

type luaState struct {
	stack 	*luaStack

	proto 	*binChunk.Prototype
	pc 		int
}

func New(stackSize int, proto *binChunk.Prototype) *luaState {
	return &luaState{
		stack: 	newLuaStack(stackSize),
		proto: 	proto,
		pc: 	0,
	}
}