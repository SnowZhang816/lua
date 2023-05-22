package state

import "main/binChunk"
import "main/api"
import "main/cLog"

type upValue struct {
	val *luaValue
}

type closure struct {
	proto *binChunk.Prototype		//lua closure
	goFunc api.GoFunction				//go closure
	upValues []*upValue
}

func newLuaClosure(proto *binChunk.Prototype) *closure {
	cLog.Println("newLuaClosure")
	c := &closure{proto: proto}
	if nUpValue := len(proto.UpValues); nUpValue > 0 {
		c.upValues = make([]*upValue, nUpValue)
	}
	return c
}

func newGoClosure(f api.GoFunction, nUpValue int) *closure {
	c := &closure{goFunc : f}
	if nUpValue > 0 {
		c.upValues = make([]*upValue, nUpValue)
	}
	return c
}