package vm

import "main/api"
import "fmt"

func move(i Instruction, vm api.LuaVM) {
	a,b,_ := i.ABC()
	fmt.Println("move", a, b)
	a += 1
	b += 1
	vm.Copy(b, a)
}

func jmp(i Instruction, vm api.LuaVM)  {
	a,sbx := i.AsBx()
	vm.AddPC(sbx)
	if a != 0 {
		panic("todo!")
	}
}