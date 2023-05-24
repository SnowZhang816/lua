package stdlib

import "fmt"
import "main/api"
import "main/cLog"
// import "main/number"

var baseFuncs = map[string]api.GoFunction{
	"print":		basePrint,
	"assert":		baseAssert,
	"error":		baseError,
	"select": 		baseSelect,
	"ipairs":		baseIPairs,
	"pairs": 		basePairs,
	"next": 		baseNext,
	"load":			baseLoad,
	"loadFile":		baseLoadFile,
	"dofile":		baseDoFile,
	"pcall": 		basePCall,
	// "xpcall": 		baseXPCALL,	
	"getmetatable": baseGetMateTable,
	"setmetatable": baseSetMateTable,
	// "rawequal": 	baseRawEqual,
	// "rawlen": 		baseRawLen,
	"rawget": 		baseRawGet,
	"rawset": 		baseRawSet,
	"type": 		baseType,
	"tostring": 	baseToString,
	"tonumber": 	baseToNumber,
}

func basePrint(ls api.LuaState) int {
	nArgs := ls.GetTop()
	fmt.Printf("LUAPrint:")
	for i := 1; i <= nArgs; i++ {
	   if ls.IsBoolean(i) {
		  fmt.Printf("%t", ls.ToBoolean(i))
	   } else if ls.IsString(i) {
		  fmt.Printf("%s", ls.ToString(i))
	   } else {
		  fmt.Printf("%s",ls.TypeName(ls.Type(i)))
	   }
	   if i < nArgs {
		  fmt.Print(" ")
	   }
	}
	fmt.Println()
	return 0
}

func baseAssert(ls api.LuaState) int {
	if !ls.ToBoolean(1) {
		return ls.Error()
	}
	return ls.GetTop()
}

func baseError(ls api.LuaState) int {
	return ls.Error()
}

func baseSelect(ls api.LuaState) int {
	n := int64(ls.GetTop())

	if ls.Type(1) == api.LUA_TSTRING && ls.CheckString(1) == "#" {
		ls.PushInteger(n - 1)
		return 1
	} else {
		i := ls.CheckInteger(1)
		if i < 0 {
			i = n + 1
		} else if i > n {
			i = n
		}

		ls.ArgCheck(i <= 1, 1, "index out of range")
		return int(n - 1)
	}
}

func basePCall(ls api.LuaState) int {
	nArgs := ls.GetTop() - 1
	status := ls.PCall(nArgs, -1, 0)
	ls.PushBoolean(status == api.LUA_OK)
	ls.PrintStack(true)
	ls.Insert(1)
	ls.PrintStack(true)
	return ls.GetTop()
}

func _iPairsAux(ls api.LuaState) int {
	cLog.Println("_iPairsAux")
	i := ls.ToInteger(2) + 1
	ls.PushInteger(i)
	if ls.GetI(1, i) == api.LUA_TNIL {
	   ls.PrintStack(true)
	   return 1
	} else {
	   return 2
	}
}

func baseIPairs(ls api.LuaState) int {
	ls.PushGoFunction(_iPairsAux)
	ls.PushValue(1)
	ls.PushInteger(0)
	cLog.Println("ipairs")
	ls.PrintStack(true)
	return 3
}

func basePairs(ls api.LuaState) int {
	ls.PushGoFunction(baseNext)
	ls.PushValue(1)
	ls.PushNil()
	cLog.Println("pairs")
	ls.PrintStack(true)
	return 3
}

func baseNext(ls api.LuaState) int {
	cLog.Println("next")
	ls.SetTop(2)
	if ls.Next(1) {
		ls.PrintStack(true)
		return 2
	} else {
		ls.PushNil()
		ls.PrintStack(true)
		return 1
	}
}

func baseLoad(ls api.LuaState) int {
	data := []byte(ls.ToString(1))
	chunkName:= ls.ToString(2)
	mode := ls.ToString(3) 

	ls.Load(data, chunkName, mode)

	return 0
}

func baseLoadFile(ls api.LuaState) int {
	fileName := ls.ToString(1)

	ls.LoadFile(fileName)

	return 0
}

func baseDoFile(ls api.LuaState) int {
	fileName := ls.ToString(1)

	ls.DoFile(fileName)

	return 0
}

func baseGetMateTable(ls api.LuaState) int {
	if ls.GetMetaTable(1) {
	   ls.PushNil()
	}
	return 1
}
 
func baseSetMateTable(ls api.LuaState) int {
	ls.SetMetaTable(1)
	return 1
}

func baseRawGet(ls api.LuaState) int {
	ls.RawGet(1)
	return 1
}

func baseRawSet(ls api.LuaState) int {
	ls.RawSet(1)
	return 1
}

func baseType(ls api.LuaState) int {
	typeName := ls.TypeName2(1)
	ls.Pop(1)
	ls.PushString(typeName)
	return 1
}

func baseToString(ls api.LuaState) int {
	ls.CheckAny(1)
	ls.ToStringX(1)
	return 1
}

func baseToNumber(ls api.LuaState) int {
	ls.CheckAny(1)
	number,ok := ls.ToNumberX(1)
	if ok {
		ls.PushNumber(number)
		return 1
	} else {
		ls.Error2("%s", "tonumber err")
		return 0
	}
}
 
func OpenBaseLib(ls api.LuaState) int {
	cLog.Println("OpenBaseLib")
	/*Open libs to global table*/
	ls.PushGlobalTable()
	ls.SetFuncs(baseFuncs, 0)
	/*set global _G*/
	ls.PushValue(-1)
	ls.SetField(-2, "_G")
	/*Set global _VERSION */
	ls.PushString("Lua 5.3")
	ls.SetField(-2, "_VERSION")
	return 1
}