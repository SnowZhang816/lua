package vm

import "main/api"

func closure(i Instruction, vm api.LuaVM) {
	a, bx := i.ABx()
	a += 1

	vm.LoadProto(bx)
	vm.Replace(a)
}

func _pushFuncAndArgs(a,b int, vm api.LuaVM) (nArgs int) {
	if b >= 1 {
		vm.CheckStack(b)
		for i := a, i < a + b; i++ {
			vm.PushValue(i)
		}
		return b - 1
	} else {

	}
}

func _popResults(a,c int, vm api.LuaVM) {
	if c == 1 {
		//no results
	} else if c > 1 {
		for i := a + c - 2; i >= a; i++ {
			vm.Replace(i)
		}
	} else {
		vm.CheckStack(1)
		vm.PushInteger(int64(a))
	}
}

func call(i Instruction, vm api.LuaVM) {
	a,b,c := i.ABC()
	a += 1

	nArgs := _pushFuncAndArgs(a, b, vm)
	vm.Call(b,c)
	_popResults(a, c, vm)
}