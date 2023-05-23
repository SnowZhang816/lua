package stdlib

import "main/api"
import "main/cLog"
import "strings"

var stringLib = map[string]api.GoFunction{
	"lower":		strLower,
}

func strLower(ls api.LuaState) int {
	s := ls.CheckString(1)
	ls.PushString(strings.ToLower(s))
	return 1
}

func createMetaTable(ls api.LuaState) {
	cLog.Println("stringLib createMetaTable")
	ls.CreateTable(0,1)
	ls.PrintStack(false)
	ls.PushString("dummy")
	ls.PrintStack(false)
	ls.PushValue(-2)
	ls.PrintStack(false)
	ls.SetMetaTable(-2)
	ls.PrintRegister()
	ls.PrintStack(false)
	ls.Pop(1)
	ls.PrintStack(false)
	ls.PushValue(-2)
	ls.PrintStack(false)
	ls.SetField(-2, "__index")
	ls.PrintStack(false)
	ls.Pop(1)
	ls.PrintStack(false)
}

func OpenStringLib(ls api.LuaState) int {
	cLog.Println("OpenStringLib")
	ls.NewLib(stringLib)
	createMetaTable(ls)
	return 1
}