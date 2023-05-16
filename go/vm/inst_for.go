package vm

import "main/api"
import "fmt"

func printStack(vm api.LuaVM) {
	top := vm.GetTop()
	for i := 1; i <= top; i++ {
	   t := vm.Type(i)
	   switch t {
	   case api.LUA_TBOOLEAN:      fmt.Printf("[%t]", vm.ToBoolean(i))
	   case api.LUA_TNUMBER:       fmt.Printf("[%g]", vm.ToNumber(i))
	   case api.LUA_TSTRING:       fmt.Printf("[%q]", vm.ToString(i))
	   default:                    fmt.Printf("[%s]", vm.TypeName(t))
	   }
	}
	fmt.Println("\n")
 }

func forPrep(i Instruction, vm api.LuaVM) {
	a, sbx := i.AsBx()
	a += 1

	//R(A) -= R(A + 2)
	vm.PushValue(a)
	// printStack(vm)
	vm.PushValue(a + 2)
	// printStack(vm)
	vm.Arith(api.LUA_OPSUB)
	// printStack(vm)
	vm.Replace(a)
	// printStack(vm)
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