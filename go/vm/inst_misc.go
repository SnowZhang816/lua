package vm

import "main/api"

func move(i Instruction, vm api.LuaVM) {
	a,b,_ := i.ABC()
	a += 1
	b += 1
	vm.Copy(b, a)
}

func jmp(i Instruction, vm api.LuaVM)  {
	a,sbx := AsBx()
	vm.AddPC(sbx)
	if a != 0 {
		panic("todo!")
	}
}