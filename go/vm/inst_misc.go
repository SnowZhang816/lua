package vm

import "main/api"
import "main/cLog"

func move(i Instruction, vm api.LuaVM) {
	a,b,_ := i.ABC()
	cLog.Println("move", a, b)
	a += 1
	b += 1
	vm.Copy(b, a)
}

func jmp(i Instruction, vm api.LuaVM)  {
	a,sbx := i.AsBx()
	vm.AddPC(sbx)
	if a != 0 {
		vm.CloseUpValues(a)
	}
}