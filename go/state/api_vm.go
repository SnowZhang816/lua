package state

import "main/cLog"

func (self *luaState) PC() int {
	return self.stack.pc
}

func (self *luaState) AddPC(n int) {
	self.stack.pc += n
}

func (self *luaState) Fetch() uint32 {
	i := self.stack.closure.proto.Code[self.stack.pc]
	self.stack.pc++
	return i
}

func (self *luaState) GetConst(idx int) {
	c := self.stack.closure.proto.Constants[idx]
	self.stack.push(c)
}

func (self *luaState) GetRK(rk int) {
	if rk > 0xFF {
		self.GetConst(rk & 0xFF)
	} else {
		self.PushValue(rk + 1)
	}
}

func (self *luaState) LoadProto(idx int) {
	proto := self.stack.closure.proto.Protos[idx]
	cLog.Printf("LoadProto[%d] %s<%d-%d>\n", idx, proto.Source, proto.LineDefined, proto.LastLineDefined)
	cLog.Println(proto.UpValues)
	closure := newLuaClosure(proto)
	self.stack.push(closure)

	for i, uvInfo := range proto.UpValues {
		uvIdx := int(uvInfo.Idx)
		if uvInfo.InStack == 1 {
			if self.stack.openuvs == nil {
				self.stack.openuvs = map[int]*upValue{}
			}
			if openuv,found := self.stack.openuvs[uvIdx]; found {
				closure.upValues[i] = openuv
			} else {
				closure.upValues[i] = &upValue{&self.stack.slots[uvIdx]}
				self.stack.openuvs[uvIdx] = closure.upValues[i]
			}
		} else {
			closure.upValues[i] = self.stack.closure.upValues[uvIdx]
		}
	}
}

func (self *luaState) RegisterCount() int {
	return int(self.stack.closure.proto.MaxStackSize)
}

func (self *luaState) LoadVararg(n int) {
	if n < 0 {
		n = len(self.stack.varargs)
	}
	self.stack.check(n)
	self.stack.pushN(self.stack.varargs, n)
}

func (self *luaState) CloseUpValues(a int) {
	for i, openuv := range self.stack.openuvs {
		if i >= a - 1 {
			val := *openuv.val
			openuv.val = &val
			delete(self.stack.openuvs, i)
		}
	}
}

func (self *luaState) PrintStack(loop bool) {
	self.printStack(loop)
}

func (self *luaState) PrintRegister() {
	self.printRegister()
}

func (self *luaState) PrintUpValues()  {
	self.printUpValues()
}

func (self *luaState) PrintLoadedTable()  {
	self.printLoadedTable()
}

func (self *luaState) PrintGlobalTable()  {
	self.printGlobalTable()
}