package vm

import "main/api"
import "fmt"

func loadNil(i Instruction, vm api.LuaVM) {
	a,b,_ := i.ABC()
	fmt.Println("loadNil", a,b)
	a += 1

	vm.PushNil()
	for i := a; i <= a + b; i++ {
		vm.Copy(-1, i)
	}
	vm.Pop(1)
}

func loadBoolean(i Instruction, vm api.LuaVM) {
	a,b,c := i.ABC()
	a += 1
	vm.PushBoolean(b != 0)
	vm.Replace(a)
	if c != 0 {
		vm.AddPC(1)
	}
}

func loadK(i Instruction, vm api.LuaVM) {
	a, bx := i.ABx()
	fmt.Println("loadK", a, bx)
	a += 1
	vm.GetConst(bx)
	vm.Replace(a)
}

func loadKx(i Instruction, vm api.LuaVM) {
	a, _ := i.ABx()
	a += 1
	ax := Instruction(vm.Fetch()).Ax()

	vm.GetConst(ax)
	vm.Replace(a)
}