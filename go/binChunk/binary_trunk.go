package binChunk

import "main/cLog"

const (
	LUA_SIGNATURE      	= 0x61754C1B//0x1B4C7561 //"\x1bLua"
	LUAC_VERSION		= 0x53
	LUAC_FORMAT			= 0
	LUAC_DATA			= "\x19\x93\r\n\x1a\n"
	CINT_SIZE  			= 4
	CSIZET_SIZE			= 8
	INSTRUCTION_SIZE	= 4
	LUA_INTEGER_SIZE 	= 8
	LUA_NUMBER_SIZE		= 8
	LUAC_INT			= 0x5678
	LUAC_NUM			= 370.5
)

const (
	TAG_NIL				= 0x00
	TAG_BOOLEAN			= 0x01
	TAG_NUMBER			= 0x03
	TAG_INTEGER			= 0x13
	TAG_SHORT_STR		= 0x04
	TAG_LONG_STR		= 0x14
)

type header struct {
	signature 			[4]byte
	version				byte
	format 				byte
	luacData 			[6]byte
	cintData 			byte
	sizetSize 			byte
	instructionSize 	byte
	luaIntegerSize 		byte
	luaNumberSize 		byte
	luacInt 			int64
	luacNum 			float64
}

type UpValue struct {
	InStack				byte
	Idx					byte
}

type LocVar struct {
	VarName				string
	StartPC				uint32
	EndPC				uint32
} 

type Prototype struct {
	Source 				string
	LineDefined			uint32
	LastLineDefined		uint32
	NumParams			byte
	IsVarArg			byte
	MaxStackSize		byte
	Code				[]uint32
	Constants			[]interface{}
	UpValues			[]UpValue
	Protos				[]*Prototype
	LineInfo			[]uint32
	LocVars				[]LocVar
	UpValuesNames		[]string			
}

type binaryChunk struct {
	header					//头部
	sizeUpValues byte		//upValue数量
	mainFunc *Prototype		//函数原型
}

func UnDump(data []byte) *Prototype {
	reader := &reader{data}
	cLog.Println("checkHeader===========")
	reader.checkHeader()
	cLog.Println("readByte===========")
	reader.readByte()
	cLog.Println("readProto===========")
	return reader.readProto("")
}

