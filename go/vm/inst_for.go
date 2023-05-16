package vm

import "main/api"

func forPrep(i Instruction, vm api.LuaVM) {
	a, sbx := i.AsBx()
	a += 1

	//R(A) -= R(A + 2)
	vm.PushValue(a)
	// api.PrintStack(vm)
	vm.PushValue(a + 2)
	// api.PrintStack(vm)
	vm.Arith(api.LUA_OPSUB)
	// api.PrintStack(vm)
	vm.Replace(a)
	// api.PrintStack(vm)
	vm.AddPC(sbx)
}

func forLoop(i Instruction, vm api.LuaVM) {
	a, sbx := i.AsBx()
	a += 1

	//R(A) += R(A + 2)
	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(api.LUA_OPADD)
	vm.Replace(a)

	//R(A) <?= R(A + 1)
	isPositiveStep := vm.ToNumber(a + 2) >= 0
	if isPositiveStep && vm.Compare(a, a + 1, api.LUA_OPLE) || !isPositiveStep && vm.Compare(a + 1, a, api.LUA_OPLE) {
		vm.AddPC(sbx)			//pc+=sbx
		vm.Copy(a, a + 3)		//R(A + 3) = R(A)
	}
}