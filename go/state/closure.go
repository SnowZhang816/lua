package state

import "main/binChunk"

type closure struct {
	proto *binChunk.Prototype
}

func newLuaClosure(proto *binChunk.Prototype) *closure {
	return &closure{proto: proto}
}