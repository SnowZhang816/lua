package api

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
	LUA_OPMUl			//	*
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