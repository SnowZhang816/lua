package vm

import "main/api"
import "main/cLog"

const LFIELDS_PER_FLUSH = 50

func newTable(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1

	cLog.Println("vm newTable", a, b, c)

	vm.CreateTable(Fb2int(b), Fb2int(c))
	vm.Replace(a)
}

func getTable(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	cLog.Println("getTable", a, b, c)
	a += 1
	b += 1

	vm.GetRK(c)
	vm.PrintStack(true)
	vm.GetTable(b)
	vm.PrintStack(true)
	vm.Replace(a)
}

func setTable(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1

	vm.GetRK(b)
	vm.GetRK(c)
	vm.SetTable(a)
}

func setList(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	cLog.Println("setList", a, b, c)
	a += 1

	bIsZero := b == 0
	if bIsZero {
		b = int(vm.ToInteger(-1)) - a - 1
		vm.Pop(1)
	}

	if c > 0 {
		c = c - 1
	} else {
		c = Instruction(vm.Fetch()).Ax()
	}

	idx := int64(c * LFIELDS_PER_FLUSH)
	for j := 1; j <= b; j++ {
		idx++
		vm.PushValue(a + j)
		vm.PrintStack(true)
		cLog.Println(a, idx)
		vm.SetI(a, idx)
	}

	if bIsZero {
		for j := vm.RegisterCount() + 1; j <= vm.GetTop(); j++{
			idx++
			vm.PushValue(j)
			vm.SetI(a, idx)
		}
		vm.SetTop(vm.RegisterCount())		//clear stack
	}

	vm.PrintTable(a)
}