package codegen

import "main/binChunk"

func toProto(fi *funcInfo) *binChunk.Prototype {
	proto := &binChunk.Prototype{
		LineDefined:     uint32(fi.line),
		LastLineDefined: uint32(fi.lastLine),
		NumParams:       byte(fi.numParams),
		MaxStackSize:    byte(fi.maxRegs),
		Code:            fi.insts,
		Constants:       getConstants(fi),
		UpValues:        getUpvalues(fi),
		Protos:          toProtos(fi.subFuncs),
		LineInfo:        fi.lineNums,
		LocVars:         getLocVars(fi),
		UpValuesNames:    getUpvalueNames(fi),
	}

	if fi.line == 0 {
		proto.LastLineDefined = 0
	}
	if proto.MaxStackSize < 2 {
		proto.MaxStackSize = 2 // todo
	}
	if fi.isVararg {
		proto.IsVarArg = 1 // todo
	}

	return proto
}

func toProtos(fis []*funcInfo) []*binChunk.Prototype {
	protos := make([]*binChunk.Prototype, len(fis))
	for i, fi := range fis {
		protos[i] = toProto(fi)
	}
	return protos
}

func getConstants(fi *funcInfo) []interface{} {
	consts := make([]interface{}, len(fi.constants))
	for k, idx := range fi.constants {
		consts[idx] = k
	}
	return consts
}

func getLocVars(fi *funcInfo) []binChunk.LocVar {
	locVars := make([]binChunk.LocVar, len(fi.locVars))
	for i, locVar := range fi.locVars {
		locVars[i] = binChunk.LocVar{
			VarName: locVar.name,
			StartPC: uint32(locVar.startPC),
			EndPC:   uint32(locVar.endPC),
		}
	}
	return locVars
}

func getUpvalues(fi *funcInfo) []binChunk.UpValue {
	upvals := make([]binChunk.UpValue, len(fi.upvalues))
	for _, uv := range fi.upvalues {
		if uv.locVarSlot >= 0 { // instack
			upvals[uv.index] = binChunk.UpValue{1, byte(uv.locVarSlot)}
		} else {
			upvals[uv.index] = binChunk.UpValue{0, byte(uv.upvalIndex)}
		}
	}
	return upvals
}

func getUpvalueNames(fi *funcInfo) []string {
	names := make([]string, len(fi.upvalues))
	for name, uv := range fi.upvalues {
		names[uv.index] = name
	}
	return names
}
