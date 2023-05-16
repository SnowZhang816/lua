package state

func (self *luaState) PC() int {
	return self.stack.pc
}

func (self *luaState) AddPC(n int) {
	self.stack.pc += n
}

func (self *luaState) Fetch() uint32 {
	i := self.stack.closure.proto.Code[self.stack.pc]
	self.pc++
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

func (self *luaState) LoadProto(n int) {
	proto := self.stack.closure.proto.Protos[n]
	self.stack.push(proto)
}