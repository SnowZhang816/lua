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

func getUpValue(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	fmt.Println("getUpValue", a, b)
	a += 1
	b += 1

	vm.Copy(LuaUpValueIndex(b), a)
}

func setUpValue(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	fmt.Println("setUpValue", a, b, c)
	a += 1
	b += 1

	vm.GetRK(c)
	vm.GetTable(LuaUpValueIndex(b))
	vm.Replace(a)
}