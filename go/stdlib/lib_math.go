package stdlib

import "math"
// import "math/rand"
import "main/api"
import "main/cLog"

var mathLib = map[string]api.GoFunction{
	"ceil":			mathCeil,
	"floor":		mathFloor,
}

func mathCeil(ls api.LuaState) int {
	val := ls.CheckNumber(1)
	val = math.Ceil(val)
	ls.Pop(1)
	ls.PushNumber(val)
	return 1
}

func mathFloor(ls api.LuaState) int {
	val := ls.CheckNumber(1)
	val = math.Floor(val)
	ls.Pop(1)
	ls.PushNumber(val)
	return 1
}

func OpenMathLib(ls api.LuaState) int {
	cLog.Println("OpenMathLib")
	ls.NewLib(mathLib)
	ls.PushNumber(math.Pi)
	ls.SetField(-2, "pi")
	ls.PushNumber(math.Inf(1))
	ls.SetField(-2, "huge")
	ls.PushInteger(math.MaxInt64)
	ls.SetField(-2, "maxinteger")
	ls.PushInteger(math.MinInt64)
	ls.SetField(-2, "mininteger")
	return 1
}