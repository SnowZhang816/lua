package state

import "main/binChunk"
import "main/aUtil"
import "main/vm"
import "main/api"
import "main/cLog"
import "fmt"
import "io/ioutil"

func (self *luaState) Load(chunk []byte, chunkName, mode string) int {
	proto := binChunk.UnDump(chunk)
	aUtil.PrintProto(proto)

	cLog.Println("\n\n\n luaState Load")
	c := newLuaClosure(proto)

	self.stack.push(c)
	if len(proto.UpValues) > 0{
		env := self.registry.get(api.LUA_RIDX_GLOBALS)
		c.upValues[0] = &upValue{&env}
		cLog.Println("Load upValues", c.upValues)
	}

	self.printStack()

	return api.LUA_OK
}

func (self *luaState) LoadString(s string) int {
	return self.Load([]byte(s), s, "bt")
}

func (self *luaState) LoadFileX(filename, mode string) int {
	if data, err := ioutil.ReadFile(filename); err == nil {
		return self.Load(data, "@" + fileName, mode)
	}
}

func (self *luaState) DoString(str string) bool {
	return self.LoadString(str) == LUA_OK && self.PCall(0, LUA_MULTRET, 0) == LUA_OK
}

func (self *luaState) LoadFile(filename string) int {
	return self.LoadFileX(fileName, "bt")
}

func (self *luaState) DoFile(filename string) bool {
	return self.LoadFile(filename) == LUA_OK && self.PCall(0, LUA_MULTRET, 0) == LUA_OK
}

func (self *luaState) runLuaClosure() {
	for {
		pc := self.PC()

		inst := vm.Instruction(self.Fetch())

		// cLog.Printf("[%02d] %s start \n", pc+1, inst.OpName())
		inst.Execute(self)
		cLog.Printf("[%02d] %s end\n", pc+1, inst.OpName())
		self.printStack()

		if inst.Opcode() == vm.OP_RETURN {
			break
		}
	}
}

func (self *luaState) callLuaClosure(nArgs, nResults int, c *closure) {
	cLog.Println("callLuaClosure", nArgs, nResults, &c.proto)

	nRegs := int(c.proto.MaxStackSize)
	nParams := int(c.proto.NumParams)
	isVarArg := c.proto.IsVarArg == 1

	cLog.Println("callLuaClosure", nRegs, nParams, isVarArg)

	newStack := newLuaStack(nRegs + api.LUA_MINSTACK, self)
	newStack.closure = c

	funcAndArgs := self.stack.popN(nArgs + 1)

	newStack.pushN(funcAndArgs[1:], nParams)
	newStack.printStack(1)
	newStack.top = nRegs
	if nArgs > nParams && isVarArg {
		newStack.varargs = funcAndArgs[nParams + 1:]
		cLog.Println("callLuaClosure set varargs\n", newStack.varargs)
	}

	self.pushLuaStack(newStack)
	self.printStack()

	self.runLuaClosure()
	
	self.popLuaStack()
	self.printStack()

	if nResults != 0 {
		cLog.Println("callLuaClosure nResults", nResults, newStack.top, nRegs, c.proto.Source, c.proto.LineDefined, c.proto.LastLineDefined)
		results := newStack.popN(newStack.top - nRegs)
		self.stack.check(len(results))
		self.stack.pushN(results, nResults)
		self.printStack()
	}
}

func (self *luaState) callGoClosure(nArgs, nResults int, c *closure) {
	cLog.Println("callGoClosure", nArgs, nResults, c.goFunc)

	newStack := newLuaStack(nArgs + api.LUA_MINSTACK, self)
	newStack.closure = c

	funcAndArgs := self.stack.popN(nArgs + 1)
	cLog.Println("callGoClosure funcAndArgs", funcAndArgs)
	newStack.pushN(funcAndArgs[1:], nArgs)

	self.pushLuaStack(newStack)
	self.printStack()
	r := c.goFunc(self)
	self.popLuaStack()

	if nResults != 0 {
		cLog.Println("callGoClosure nResults", r)
		results := newStack.popN(r)
		self.stack.check(len(results))
		self.stack.pushN(results, nResults)
		self.printStack()
	}
}

func (self *luaState) Call(nArgs, nResults int) {
	val := self.stack.get(-(nArgs + 1))
	c,ok := val.(*closure)
	cLog.Println("Call", c, ok)
	if !ok {
		if mf := getMetaField(val, "__call", self); mf != nil {
			if c,ok = mf.(*closure); ok {
				self.stack.push(val)
				self.Insert(-(nArgs + 2))
				nArgs += 1
			}
		}
	}

	if ok {
		if c.proto != nil {
			cLog.Printf("Call %s<%d,%d>\n", c.proto.Source, c.proto.LineDefined, c.proto.LastLineDefined)
			self.callLuaClosure(nArgs, nResults, c)
		} else {
			cLog.Printf("GoFunc\n")
			self.callGoClosure(nArgs, nResults, c)
		}
	} else {
		panic("not function!")
	}
}

func (self *luaState) PCall(nArgs, nResults, msgh int) (status int) {
	caller := self.stack
	status = api.LUA_ERRRUN

	//catch err
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("PCall", err)
			for self.stack != caller {
				self.popLuaStack()
			}
			self.stack.push(err)
		}
	}()

	self.Call(nArgs, nResults)
	status = api.LUA_OK
	return
}

