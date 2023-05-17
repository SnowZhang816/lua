package state

import "fmt"
import "main/binChunk"
import "main/aUtil"
import "main/vm"

func (self *luaState) Load(chunk []byte, chunkName, mode string) int {
	proto := binChunk.UnDump(chunk)
	aUtil.PrintProto(proto)

	fmt.Println("\n\n\nstart new closure!")
	c := newLuaClosure(proto)

	self.stack.push(c)

	self.printStack()

	return 0
}

func (self *luaState) runLuaClosure() {
	for {
		pc := self.PC()

		inst := vm.Instruction(self.Fetch())
		inst.Execute(self)

		fmt.Printf("[%02d] %s \n", pc+1, inst.OpName())
		self.printStack()

		if inst.Opcode() == vm.OP_RETURN {
			break
		}
	}
}

func (self *luaState) callLuaClosure(nArgs, nResults int, c *closure) {
	fmt.Println("callLuaClosure", nArgs, nResults, c)

	nRegs := int(c.proto.MaxStackSize)
	nParams := int(c.proto.NumParams)
	isVarArg := c.proto.IsVarArg == 1

	fmt.Println("callLuaClosure", nRegs, nParams, isVarArg)

	newStack := newLuaStack(nRegs + 20)
	newStack.closure = c

	funcAndArgs := self.stack.popN(nArgs + 1)

	newStack.pushN(funcAndArgs[1:], nParams)

	newStack.top = nRegs
	if nArgs > nParams && isVarArg {
		newStack.varargs = funcAndArgs[nParams + 1:]
	}

	self.pushLuaStack(newStack)
	self.printStack()

	self.runLuaClosure()
	
	self.popLuaStack()
	self.printStack()

	if nResults != 0 {
		results := newStack.popN(newStack.top - nRegs)
		self.stack.check(len(results))
		self.stack.pushN(results, nResults)
	}
}

func (self *luaState) Call(nArgs, nResults int) {
	val := self.stack.get(-(nArgs + 1))
	if c,ok := val.(*closure); ok {
		fmt.Printf("Call %s<%d,%d>\n", c.proto.Source, c.proto.LineDefined, c.proto.LastLineDefined)
		self.callLuaClosure(nArgs, nResults, c)
	} else {
		panic("not function!")
	}
}
