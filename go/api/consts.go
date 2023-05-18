package api

const LUA_MINSTACK = 20
const LUA_MAXSTACK = 1000000
const LUA_REGISTRY_INDEX = -LUA_MAXSTACK - 1000
const LUA_RIDX_GLOBALS int64 = 2

const (
	LUA_TNONE = iota - 1
	LUA_TNIL
	LUA_TBOOLEAN
	LUA_TLIGHTUSRDATA
	LUA_TNUMBER
	LUA_TSTRING
	LUA_TTABLE
	LUA_TFUNCTION
	LUA_TUSERDATA
	LUA_TTHREAD
)

const (
	LUA_OPADD = iota	//	+
	LUA_OPSUB			//	-
	LUA_OPMUL			//	*
	LUA_OPMOD			//	%
	LUA_OPPOW			//	^
	LUA_OPDIV			//	/
	LUA_OPIDIV			//	//
	LUA_OPAND			//	&
	LUA_OPOR			//	|
	LUA_OPBXOR			//	~
	LUA_OPSHL			//	<<
	LUA_OPSHR			//	>>
	LUA_OPUNM			//	-
	LUA_OPBONT			//	~
)

const (
	LUA_OPEQ	= iota	//	==
	LUA_OPLT			//	<
	LUA_OPLE			//	<=
)		