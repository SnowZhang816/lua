package vm

import "main/api"
import "fmt"

func closure(i Instruction, vm api.LuaVM) {
	a, bx := i.ABx()
	fmt.Println("closure", a, bx)
	a += 1

	vm.LoadProto(bx)
	vm.Replace(a)
}

func _fixStack(a int, vm api.LuaVM) {
	fmt.Println("_fixStack", a)
	x := int(vm.ToInteger(-1))
	vm.Pop(1)

	vm.CheckStack(x - a)
	for i := a; i < x; i++ {
		vm.PushValue(i)
	}
	vm.PrintStack()
	vm.Rotate(vm.RegisterCount() + 1, x - a)
}

func _pushFuncAndArgs(a,b int, vm api.LuaVM) int {
	fmt.Println("_pushFuncAndArgs", a, b)
	if b >= 1 {
		vm.CheckStack(b)
		for i := a; i < a + b; i++ {
			vm.PushValue(i)
		}
		return b - 1
	} else {
		_fixStack(a, vm)
		nArgs := vm.GetTop() - vm.RegisterCount() - 1
		fmt.Println("_pushFuncAndArgs nArgs", nArgs, vm.GetTop(), vm.RegisterCount())
		return nArgs
	}
}

func _popResults(a,c int, vm api.LuaVM) {
	fmt.Println("_popResults", a, c)
	if c == 1 {
		//no results
	} else if c > 1 {
		for i := a + c - 2; i >= a; i-- {
			vm.Replace(i)
		}
		vm.PrintStack()
	} else {
		vm.CheckStack(1)
		vm.PushInteger(int64(a))
	}
}

func call(i Instruction, vm api.LuaVM) {
	a,b,c := i.ABC()
	fmt.Println("call", a,b,c)
	a += 1

	nArgs := _pushFuncAndArgs(a, b, vm)
	vm.PrintStack()
	vm.Call(nArgs, c - 1)
	_popResults(a, c, vm)
}

func _return(i Instruction, vm api.LuaVM) {
	a,b,_ := i.ABC()
	fmt.Println("_return", a, b)
	a += 1

	if b == 1 {

	} else if b > 1 {
		vm.CheckStack(b - 1)
		for i := a; i <= a + b - 2; i++{
			vm.PushValue(i)
		}
	} else {
		_fixStack(a, vm)
	}
}

func vararg(i Instruction, vm api.LuaVM) {
	a,b,_ := i.ABC()
	fmt.Println("vararg", a,b)
	a += 1

	if b != 1 {
		vm.LoadVararg(b - 1)
		_popResults(a, b, vm)
	}
}

func tailCall(i Instruction, vm api.LuaVM) {
	a,b,_ := i.ABC()
	a += 1
	c := 0

	nArgs := _pushFuncAndArgs(a, b, vm)
	vm.Call(nArgs, c - 1)
	_popResults(a, c, vm)
}

func self(i Instruction, vm api.LuaVM) {
	a,b,c := i.ABC()
	fmt.Println("self", a,b,c)
	a += 1
	b += 1

	vm.Copy(b, a + 1)
	vm.GetRK(c)
	vm.GetTable(b)
	vm.Replace(a)
	vm.PrintStack()
}

func tForCall(i Instruction, vm api.LuaVM) {
	a, _, c := i.ABC()
	fmt.Println("tForCall", a, c)
	a += 1

	_pushFuncAndArgs(a, 3, vm)
	vm.Call(2,c)
	_popResults(a + 3, c + 1, vm)
}

func tForLoop(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	fmt.Println("tForLoop", a, sBx)
	a += 1

	if !vm.IsNil(a + 1) {
		vm.Copy(a + 1, a)
		vm.AddPC(sBx)
	}
}