package state

import "fmt"
import "main/binChunk"
import "main/aUtil"
import "main/vm"
import "main/api"

func (self *luaState) Load(chunk []byte, chunkName, mode string) int {
	proto := binChunk.UnDump(chunk)
	aUtil.PrintProto(proto)

	fmt.Println("\n\n\n luaState Load")
	c := newLuaClosure(proto)

	self.stack.push(c)
	if len(proto.UpValues) > 0{
		env := self.registry.get(api.LUA_RIDX_GLOBALS)
		c.upValues[0] = &upValue{&env}
		fmt.Println("Load upValues", c.upValues)
	}

	self.printStack()

	return 0
}

func (self *luaState) runLuaClosure() {
	for {
		pc := self.PC()

		inst := vm.Instruction(self.Fetch())

		// fmt.Printf("[%02d] %s start \n", pc+1, inst.OpName())
		inst.Execute(self)
		fmt.Printf("[%02d] %s end\n", pc+1, inst.OpName())
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

	newStack := newLuaStack(nRegs + api.LUA_MINSTACK, self)
	newStack.closure = c

	funcAndArgs := self.stack.popN(nArgs + 1)

	newStack.pushN(funcAndArgs[1:], nParams)
	newStack.printStack(1)
	newStack.top = nRegs
	if nArgs > nParams && isVarArg {
		newStack.varargs = funcAndArgs[nParams + 1:]
		fmt.Println("callLuaClosure set varargs\n", newStack.varargs)
	}

	self.pushLuaStack(newStack)
	self.printStack()

	self.runLuaClosure()
	
	self.popLuaStack()
	self.printStack()

	if nResults != 0 {
		fmt.Println("callLuaClosure nResults", nResults, newStack.top, nRegs, c.proto.Source, c.proto.LineDefined, c.proto.LastLineDefined)
		results := newStack.popN(newStack.top - nRegs)
		self.stack.check(len(results))
		self.stack.pushN(results, nResults)
		self.printStack()
	}
}

func (self *luaState) callGoClosure(nArgs, nResults int, c *closure) {
	newStack := newLuaStack(nArgs + api.LUA_MINSTACK, self)
	newStack.closure = c

	args := self.stack.popN(nArgs)
	newStack.pushN(args, nArgs)

	self.pushLuaStack(newStack)
	r := c.goFunc(self)
	self.popLuaStack()

	if nResults != 0 {
		fmt.Println("callGoClosure nResults", r)
		results := newStack.popN(r)
		self.stack.check(len(results))
		self.stack.pushN(results, nResults)
		self.printStack()
	}
}

func (self *luaState) Call(nArgs, nResults int) {
	val := self.stack.get(-(nArgs + 1))
	if c,ok := val.(*closure); ok {
		if c.proto != nil {
			fmt.Printf("Call %s<%d,%d>\n", c.proto.Source, c.proto.LineDefined, c.proto.LastLineDefined)
			self.callLuaClosure(nArgs, nResults, c)
		} else {
			fmt.Printf("GoFunc\n")
			self.callGoClosure(nArgs, nResults, c)
		}
	} else {
		panic("not function!")
	}
}
