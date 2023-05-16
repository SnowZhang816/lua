package api

import "fmt"

type LuaVM interface {
	LuaState
	PC() int					//返回当前PC
	AddPC(n int)				//修改PC
	Fetch() uint32				//取出当前指令，将PC指向下一条指令
	GetConst(idx int)			//将指定常量推入栈顶
	GetRK(rk int)				//将指定常量或栈值推入栈顶
	LoadProto(n int)
}

func PrintStack(vm LuaVM) {
	top := vm.GetTop()
	for i := 1; i <= top; i++ {
	   t := vm.Type(i)
	   switch t {
	   case LUA_TBOOLEAN:      fmt.Printf("[%t]", vm.ToBoolean(i))
	   case LUA_TNUMBER:       fmt.Printf("[%g]", vm.ToNumber(i))
	   case LUA_TSTRING:       fmt.Printf("[%q]", vm.ToString(i))
	   default:                fmt.Printf("[%s]", vm.TypeName(t))
	   }
	}
	fmt.Println("\n")
 }