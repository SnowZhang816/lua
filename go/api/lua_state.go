package api

type LuaType = int
type ArithOp = int
type CompareOP = int

type BasicAPI interface {
	/* basic stack manipulation */
	GetTop() int
	AbsIndex(idx int) int
	CheckStack(n int) bool
	Pop(n int)
	Copy(formIdx, toIdx int)
	PushValue(idx int)
	Replace(idx int)
	Insert(idx int)
	Remove(idx int)
	Rotate(idx, n int)
	SetTop(inx int)
	/* access functions (Stack -> Go) */
	TypeName(tp LuaType) string
	Type(idx int) LuaType
	IsNone(idx int) bool
	IsNil(idx int) bool
	IsNoneOrNil(idx int) bool
	IsBoolean(idx int) bool
	IsInteger(idx int) bool
	IsNumber(idx int) bool
	IsString(idx int) bool
	ToBoolean(idx int) bool
	ToInteger(idx int) int64
	ToIntegerX(idx int) (int64, bool)
	ToNumber(idx int) float64
	ToNumberX(idx int) (float64, bool)
	ToString(idx int) string
	ToStringX(idx int) (string, bool)
	/* push functions (Go -> Stack) */
	PushNil()
	PushBoolean(b bool)
	PushInteger(n int64)
	PushNumber(n float64)
	PushString(s string)
	PushFString(fmt string, a ...interface{})
	/* */
	Arith(op ArithOp)
	Compare(idx1, idx2 int, op CompareOP) bool
	Len(idx int)
	Concat(n int)
	/* get functions (lua -> Stack) */
	NewTable()
	CreateTable(nArr, nRec int)
	GetTable(idx int) LuaType
	GetField(idx int, k string) LuaType
	GetI(idx int, i int64) LuaType
	/* set functions (Stack -> Lua) */
	SetTable(idx int)
	SetField(idx int, k string)
	SetI(idx int, n int64)
	/*function load an call*/
	Load(chunk []byte, chunkName, mode string) int
	Call(nArgs, nResults int)
	/*Go Function*/
	PushGoFunction(f GoFunction, n int)
	IsGoFunction(idx int) bool
	ToGoFunction(idx int) GoFunction
	/*Global table operator*/
	PushGlobalTable()
	SetGlobal(name string)
	GetGlobal(name string) LuaType
	Register(name string, f GoFunction)
	/*meta Method*/
	GetMetaTable(idx int) bool
	SetMetaTable(idx int)
	RawLen(idx int)
	RawEqual(idx1,idx2 int) bool
	RawGet(idx int) LuaType
	RawSet(idx int)
	RawGetI(idx int, i int64) LuaType
	RawSetI(idx int, i int64)
	/**/
	Next(idx int) bool
	/**/
	Error() int
	PCall(nArgs, nResults, msgh int) int
	/**/
	PrintTable(idx int)
	PrintStack()
	PrintLoadedTable()
	PrintGlobalTable()
}

type GoFunction func(LuaState) int

func LuaUpValueIndex(i int) int {
	return LUA_REGISTRY_INDEX - i
}

type LuaState interface {
	BasicAPI
	AuxLib
}