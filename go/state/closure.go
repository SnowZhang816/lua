package state

import "main/binChunk"
import "main/api"

type closure struct {
	proto *binChunk.Prototype		//lua closure
	goFunc api.GoFunction				//go closure
}

func newLuaClosure(proto *binChunk.Prototype) *closure {
	return &closure{proto: proto}
}

func newGoClosure(f api.GoFunction) *closure {
	return &closure{goFunc : f}
}