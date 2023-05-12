package binChunk

import "encoding/binary"
import "math"
import (
	"fmt"
	"bytes"
)

type reader struct {
	data []byte
}

func (self *reader) readByte() byte {
	b := self.data[0]
	self.data = self.data[1:]
	return b
}

func (self *reader) readUint32() uint32 {
	// i := binary.BigEndian.Uint32(self.data)
	i := binary.LittleEndian.Uint32(self.data)
	self.data = self.data[4:]
	return i
}

func (self *reader) readUint64() uint64 {
	// i := binary.BigEndian.Uint64(self.data)
	i := binary.LittleEndian.Uint64(self.data)
	self.data = self.data[8:]
	return i
}

func (self *reader) readLuaInteger() int64 {
	return int64(self.readUint64())
}

func (self *reader) readLuaNumber() float64 {
	return math.Float64frombits(self.readUint64())
}

func (self *reader) readBytes(n uint) []byte {
	b := self.data[:n]
	self.data = self.data[n:]
	return b
}

func (self *reader) readString() string {
	size := uint(self.readByte())
	fmt.Println("readString", size)
	if size == 0 {		//NULL空字符串
		return ""
	}
	if size == 0xFF {
		size = uint(self.readUint64())
	}
	b := self.readBytes(size - 1)
	return string(b)
}

func (self *reader) checkHeader() {
	fmt.Println("LUA_SIGNATURE", LUA_SIGNATURE)
	b := self.readBytes(4)
	fmt.Println("bytes", b)
	var x int32
	bytesBuffer := bytes.NewBuffer(b)
    binary.Read(bytesBuffer, binary.LittleEndian, &x)
	fmt.Println("x", x)
	// x := self.readUint32()
	// fmt.Println("x", x)

	if x != LUA_SIGNATURE {
		panic("not a precomiled chunk!")
	} else if self.readByte() != LUAC_VERSION {
		panic("version mismatch")
	} else if self.readByte() != LUAC_FORMAT {
		panic("format mismatch")
	} else if string(self.readBytes(6)) != LUAC_DATA {
		panic("corrupted")
	} else if self.readByte() != CINT_SIZE {
		panic("int size mismatch")
	} else if self.readByte() != CSIZET_SIZE {
		panic("size_t size mismatch")
	} else if self.readByte() != INSTRUCTION_SIZE {
		panic("instruction size mismatch")
	} else if self.readByte() != LUA_INTEGER_SIZE {
		panic("lua_integer size mismatch")
	} else if self.readByte() != LUA_NUMBER_SIZE {
		panic("lua_number size mismatch")
	} else {
		fmt.Println("LUAC_INT", LUAC_INT)
		// integer := self.readLuaInteger() 
		// fmt.Println("integer", integer)

		b := self.readBytes(8)
		fmt.Println("b", b)
		var x1 int64
		bytesBuffer := bytes.NewBuffer(b)
		binary.Read(bytesBuffer, binary.LittleEndian, &x1)
		fmt.Println("x1", x1)

		if x1 != LUAC_INT {
			panic("endianness mismatch")
		} else if self.readLuaNumber() != LUAC_NUM {
			panic("float format mismatch")
		}
	} 
}

func (self *reader) readCode() []uint32 {
	code := make([]uint32, self.readUint32())
	for i := range code {
		code[i] = self.readUint32()
	}
	return code
}

func (self *reader) readConstant() interface{} {
	tag := self.readByte()

	fmt.Println("tag", tag)

	switch tag {
	case TAG_NIL:			return nil
	case TAG_BOOLEAN:		return self.readByte() != 0
	case TAG_INTEGER:		return self.readLuaInteger()
	case TAG_NUMBER:		return self.readLuaNumber()
	case TAG_SHORT_STR:		return self.readString()
	case TAG_LONG_STR:		return self.readString()
	default:				panic("corrupted!")
	}
}

func (self *reader) readConstants() []interface{} {
	constants := make([]interface{}, self.readUint32())

	fmt.Println("readConstants=======", constants)

	for i := range constants {
		constants[i] = self.readConstant()
	}
	return constants
}

func (self *reader) readUpValue() UpValue {
	return UpValue{
		InStack: 	self.readByte(),
		Idx: 		self.readByte(),
	}
}

func (self *reader) readUpValues() []UpValue {
	upValues := make([]UpValue, self.readUint32())
	for i := range upValues {
		upValues[i] = self.readUpValue()
	}
	return upValues
}

func (self *reader) readLineInfo() []uint32 {
	lineInfo := make([]uint32, self.readUint32())
	for i := range lineInfo {
		lineInfo[i] = self.readUint32()
	}
	return lineInfo
}

func (self *reader) readLocVars() []LocVar {
	locVars := make([]LocVar, self.readUint32())
	for i := range locVars {
		locVars[i] = LocVar{
			VarName: 	self.readString(),
			StartPC: 	self.readUint32(),
			EndPC: 		self.readUint32(),
		}
	}
	return locVars
}

func (self *reader) readUpValuesNames() []string {
	upValuesNames := make([]string, self.readUint32())
	for i := range upValuesNames {
		upValuesNames[i] = self.readString()
	}
	return upValuesNames
}

func (self *reader) readProtos(parentSource string) []*Prototype {
	protos := make([]*Prototype, self.readUint32())
	for i := range protos {
		protos[i] = self.readProto(parentSource)
	}
	return protos
}

func (self *reader) readLineDefined() uint32 {
	lineDefined := self.readUint32()
	fmt.Println("readLineDefined=======", lineDefined)
	return lineDefined
}


func (self *reader) readLastLineDefined() uint32 {
	lastLineDefined := self.readUint32()
	fmt.Println("readLastLineDefined=======", lastLineDefined)
	return lastLineDefined
}


func (self *reader) readNumParams() byte {
	numParams := self.readByte()
	fmt.Println("readNumParams=======", numParams)
	return numParams
}

func (self *reader) readIsVarArg() byte {
	isVarArg := self.readByte()
	fmt.Println("readIsVarArg=======", isVarArg)
	return isVarArg
}

func (self *reader) readMaxStackSize() byte {
	maxStackSize := self.readByte()
	fmt.Println("readMaxStackSize=======", maxStackSize)
	return maxStackSize
}

func (self *reader) readProto(parentSource string) *Prototype {
	source := self.readString()
	if source == "" {
		source = parentSource
	}

	return &Prototype{
		Source: 			source,
		LineDefined: 		self.readLineDefined(),
		LastLineDefined: 	self.readLastLineDefined(),
		NumParams: 			self.readNumParams(),
		IsVarArg: 			self.readIsVarArg(),
		MaxStackSize: 		self.readMaxStackSize(),
		Code:				self.readCode(),
		Constants: 			self.readConstants(),
		UpValues: 			self.readUpValues(),
		Protos:				self.readProtos(source),
		LineInfo:			self.readLineInfo(),
		LocVars:			self.readLocVars(),
		UpValuesNames:		self.readUpValuesNames(),
	}
} 


