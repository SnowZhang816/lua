package vm

import "main/api"
import "fmt"

func getTabUp(i Instruction, vm api.LuaVM) {
	a, _, c := i.ABC()
	fmt.Println("getTabUp", a, c)
	a += 1

	vm.PushGlobalTable()
	vm.GetRK(c)
	vm.GetTable(-2)
	vm.Replace(a)
	vm.Pop(1)
}