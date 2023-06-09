package codegen

import "main/binChunk"
import "main/compiler/ast"

func GenProto(chunk *ast.Block) *binChunk.Prototype {
	fd := &ast.FuncDefExp{
		LastLine: chunk.LastLine,
		IsVararg: true,
		Block:    chunk,
	}

	fi := newFuncInfo(nil, fd)
	fi.addLocVar("_ENV", 0)
	cgFuncDefExp(fi, fd, 0)
	return toProto(fi.subFuncs[0])
}
