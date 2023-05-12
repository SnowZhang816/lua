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
	OP_pow
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
	OP_TFORLOOP
	OP_SETLIST
	OP_CLOSURE
	OP_VARARG
	OP_EXTRAARG
)

const (
	OpArgN = iota		//argument is not used
	OpArgU				//argument is used
	OpArgR				//argument is a register or a jump offset
	OPArgK				//argument is a constant or register/constant
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
	opcode{0,	1,	OpArgR,  OpArgN,   IABC,    "MOVE"		},
}