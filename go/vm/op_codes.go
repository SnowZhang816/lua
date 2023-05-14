package vm

const (
	IABC 	= 	iota
	IABx
	IAsBx
	IAx
)

const (
	OP_MOVE		= iota
	OP_LOADK
	OP_LOADKX
	OP_LOADBOOL
	OP_LOADNIL
	OP_GETUPVAL
	OP_GETTABUP
	OP_GETTABLE
	OP_SETTABUP
	OP_SETUPVAL
	OP_SETTABLE
	OP_NEWTABLE
	OP_SELF
	OP_ADD
	OP_SUB
	OP_MUL
	OP_MOD
	OP_POW
	OP_DIV
	OP_IDIV
	OP_BAND
	OP_BOR
	OP_BXOR
	OP_SHL
	OP_SHR
	OP_UNM
	OP_BNOT
	OP_NOT
	OP_LEN
	OP_CONCAT
	OP_JMP
	OP_EQ
	OP_LT
	OP_LE
	OP_TEST
	OP_TESTSET
	OP_CALL
	OP_TAILCALL
	OP_RETURN
	OP_FORLOOP
	OP_FORPREP
	OP_TFORCALL
	OP_TFORLOOP
	OP_SETLIST
	OP_CLOSURE
	OP_VARARG
	OP_EXTRAARG
)

const (
	OpArgN = iota		//argument is not used(不使用)
	OpArgU				//argument is used(布尔值、整数值、upvalue值、子函数索引)
	OpArgR				//argument is a register or a jump offset(IABC模式表示寄存器索引，iAsBx模式表示跳转偏移)
	OPArgK				//argument is a constant or register/constant(表示常量索引/寄存器索引，B、C操作数的最高位为1，则表示常量索引表，否则表示寄存器索引表)
)

type opcode struct {
	testFlag		byte	//operator is a test(next instruction must be a jump)
	setAFlag		byte	//instruction set register A 
	argBMode		byte	//B arg code
	argCMode		byte	//C arg code
	opMode			byte	//op mode
	name 			string	//
}

var opcodes = []opcode{
	/* 	   T 	A      B        C  	   mode     name        */
	opcode{0,	1,	OpArgR,  OpArgN,   IABC,    "MOVE	  	"},
	opcode{0,	1,	OPArgK,  OpArgN,   IABx,    "LOADK	   	"},
	opcode{0,	1,	OpArgN,  OpArgN,   IABx,    "LOADKX	   	"},
	opcode{0,	1,	OpArgU,  OpArgU,   IABC,    "LOADBOOL	"},
	opcode{0,	1,	OpArgU,  OpArgN,   IABC,    "LOADNIL	"},
	opcode{0,	1,	OpArgU,  OpArgN,   IABC,    "GETUPVAL	"},
	opcode{0,	1,	OpArgU,  OPArgK,   IABC,    "GETTABUP	"},
	opcode{0,	1,	OpArgR,  OPArgK,   IABC,    "GETTABLE	"},
	opcode{0,	0,	OPArgK,  OPArgK,   IABC,    "SETTABUP	"},
	opcode{0,	0,	OpArgU,  OpArgN,   IABC,    "SETUPVAL	"},
	opcode{0,	0,	OPArgK,  OPArgK,   IABC,    "SETTABLE	"},
	opcode{0,	1,	OpArgU,  OpArgU,   IABC,    "NEWTABLE	"},
	opcode{0,	1,	OpArgR,  OPArgK,   IABC,    "SELF		"},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "ADD		"},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "SUB		"},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "MUL		"},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "MOD		"},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "POW		"},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "DIV		"},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "IDIV		"},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "BAND		"},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "BOR		"},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "BXOR		"},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "SHL		"},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "SHR		"},
	opcode{0,	1,	OpArgR,  OpArgN,   IABC,    "UNM		"},
	opcode{0,	1,	OpArgR,  OpArgN,   IABC,    "BNOT		"},
	opcode{0,	1,	OpArgR,  OpArgN,   IABC,    "NOT		"},
	opcode{0,	1,	OpArgR,  OpArgN,   IABC,    "LEN		"},
	opcode{0,	1,	OpArgR,  OpArgR,   IABC,    "CONCAT		"},
	opcode{0,	0,	OpArgR,  OpArgN,   IAsBx,   "JMP		"},
	opcode{1,	0,	OPArgK,  OPArgK,   IABC,    "EQ			"},
	opcode{1,	0,	OPArgK,  OPArgK,   IABC,    "LT			"},
	opcode{1,	0,	OPArgK,  OPArgK,   IABC,    "LE			"},
	opcode{1,	0,	OpArgN,  OpArgU,   IABC,    "TEST		"},
	opcode{1,	1,	OpArgR,  OpArgU,   IABC,    "TESTSET	"},
	opcode{0,	1,	OpArgU,  OpArgU,   IABC,    "CALL		"},
	opcode{0,	1,	OpArgU,  OpArgU,   IABC,    "TAILCALL	"},
	opcode{0,	0,	OpArgU,  OpArgN,   IABC,    "RETURN		"},
	opcode{0,	1,	OpArgR,  OpArgN,   IAsBx,   "FORLOOP	"},
	opcode{0,	1,	OpArgR,  OpArgN,   IAsBx,   "FORPREP	"},
	opcode{0,	0,	OpArgN,  OpArgU,   IABC,    "TFORCALL	"},
	opcode{0,	1,	OpArgR,  OpArgN,   IAsBx,   "TFORLOOP	"},
	opcode{0,	0,	OpArgU,  OpArgU,   IABC,    "SETLIST	"},
	opcode{0,	1,	OpArgU,  OpArgN,   IABC,    "CLOSURE	"},
	opcode{0,	1,	OpArgU,  OpArgN,   IABC,    "VARARG		"},
	opcode{0,	0,	OpArgU,  OpArgU,   IAx,     "EXTRAARG	"},
}