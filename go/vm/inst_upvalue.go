package vm

import "main/api"
import "fmt"

func getTabUp(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	fmt.Println("getTabUp", a, b, c)
	a += 1
	b += 1

	vm.GetRK(c)
	vm.GetTable(api.LuaUpValueIndex(b))
	vm.Replace(a)
}

func setTabUp(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	fmt.Println("getTabUp", a, b, c)
	a += 1

	vm.GetRK(b)
	vm.GetRK(c)
	vm.SetTable(api.LuaUpValueIndex(a))
}


func getUpValue(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	fmt.Println("getUpValue", a, b)
	a += 1
	b += 1

	vm.Copy(api.LuaUpValueIndex(b), a)
}

func setUpValue(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	fmt.Println("setUpValue", a, b)
	a += 1
	b += 1

	vm.Copy(a, api.LuaUpValueIndex(b))
}