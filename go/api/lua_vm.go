package api

type LuaVM interface {
	LuaState
	PC() int					//返回当前PC
	AddPC(n int)				//修改PC
	Fetch() uint32				//取出当前指令，将PC指向下一条指令
	GetConst(idx int)			//将指定常量推入栈顶
	GetRK(rk int)				//将指定常量或栈值推入栈顶
	LoadProto(idx int)
	RegisterCount() int			//寄存器数量
	LoadVararg(n int)			//加载vararg参数
	CloseUpValues(a int)


	PrintUpValues()
}