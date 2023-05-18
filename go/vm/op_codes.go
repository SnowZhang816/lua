package vm

import "main/api"

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
	action			func(i Instruction, vm api.LuaVM)
}

var opcodes = []opcode{
	/* 	   T 	A      B        C  	   mode     name            action	*/
	opcode{0,	1,	OpArgR,  OpArgN,   IABC,    "MOVE       ",	move  		},
	opcode{0,	1,	OPArgK,  OpArgN,   IABx,    "LOADK      ",	loadK		},
	opcode{0,	1,	OpArgN,  OpArgN,   IABx,    "LOADKX     ",	loadKx		},
	opcode{0,	1,	OpArgU,  OpArgU,   IABC,    "LOADBOOL   ",	loadBoolean	},
	opcode{0,	1,	OpArgU,  OpArgN,   IABC,    "LOADNIL    ",	loadNil		},
	opcode{0,	1,	OpArgU,  OpArgN,   IABC,    "GETUPVAL   ",	getUpValue	},
	opcode{0,	1,	OpArgU,  OPArgK,   IABC,    "GETTABUP   ",	getTabUp	},
	opcode{0,	1,	OpArgR,  OPArgK,   IABC,    "GETTABLE   ",	getTable	},
	opcode{0,	0,	OPArgK,  OPArgK,   IABC,    "SETTABUP   ",	setTabUp	},
	opcode{0,	0,	OpArgU,  OpArgN,   IABC,    "SETUPVAL   ",	setUpValue	},
	opcode{0,	0,	OPArgK,  OPArgK,   IABC,    "SETTABLE   ",	setTable	},
	opcode{0,	1,	OpArgU,  OpArgU,   IABC,    "NEWTABLE   ",	newTable	},
	opcode{0,	1,	OpArgR,  OPArgK,   IABC,    "SELF       ",	self		},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "ADD        ",	add			},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "SUB        ",	sub			},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "MUL        ",	mul			},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "MOD        ",	mod			},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "POW        ",	pow			},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "DIV        ",	div			},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "IDIV       ",	idiv		},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "BAND       ",	band		},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "BOR        ",	bor			},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "BXOR       ",	bxor		},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "SHL        ",	shl			},
	opcode{0,	1,	OPArgK,  OPArgK,   IABC,    "SHR        ",	shr			},
	opcode{0,	1,	OpArgR,  OpArgN,   IABC,    "UNM        ",	unm			},
	opcode{0,	1,	OpArgR,  OpArgN,   IABC,    "BNOT       ",	bnot		},
	opcode{0,	1,	OpArgR,  OpArgN,   IABC,    "NOT        ",	not			},
	opcode{0,	1,	OpArgR,  OpArgN,   IABC,    "LEN        ",	len			},
	opcode{0,	1,	OpArgR,  OpArgR,   IABC,    "CONCAT     ",	concat		},
	opcode{0,	0,	OpArgR,  OpArgN,   IAsBx,   "JMP        ",	jmp			},
	opcode{1,	0,	OPArgK,  OPArgK,   IABC,    "EQ         ",	eq			},
	opcode{1,	0,	OPArgK,  OPArgK,   IABC,    "LT         ",	lt			},
	opcode{1,	0,	OPArgK,  OPArgK,   IABC,    "LE         ",	le			},
	opcode{1,	0,	OpArgN,  OpArgU,   IABC,    "TEST       ",	test		},
	opcode{1,	1,	OpArgR,  OpArgU,   IABC,    "TESTSET    ",	testSet		},
	opcode{0,	1,	OpArgU,  OpArgU,   IABC,    "CALL       ",	call		},
	opcode{0,	1,	OpArgU,  OpArgU,   IABC,    "TAILCALL   ",	tailCall	},
	opcode{0,	0,	OpArgU,  OpArgN,   IABC,    "RETURN     ",	_return		},
	opcode{0,	1,	OpArgR,  OpArgN,   IAsBx,   "FORLOOP    ",	forLoop		},
	opcode{0,	1,	OpArgR,  OpArgN,   IAsBx,   "FORPREP    ",	forPrep		},
	opcode{0,	0,	OpArgN,  OpArgU,   IABC,    "TFORCALL   ",	nil			},
	opcode{0,	1,	OpArgR,  OpArgN,   IAsBx,   "TFORLOOP   ",	nil			},
	opcode{0,	0,	OpArgU,  OpArgU,   IABC,    "SETLIST    ",	setList		},
	opcode{0,	1,	OpArgU,  OpArgN,   IABx,    "CLOSURE    ",	closure	  	},
	opcode{0,	1,	OpArgU,  OpArgN,   IABC,    "VARARG     ",	vararg		},
	opcode{0,	0,	OpArgU,  OpArgU,   IAx,     "EXTRAARG   ",	nil			},
}		