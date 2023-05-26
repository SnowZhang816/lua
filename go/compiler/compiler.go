package compiler

import "main/binChunk"
import "main/compiler/codegen"
import "main/compiler/parse"

func Compile(chunk, chunkName string) *binChunk.Prototype {
	ast := parse.Parse(chunk, chunkName)
	proto := codegen.GenProto(ast)
	setSource(proto, chunkName)
	return proto
}

func setSource(proto *binChunk.Prototype, chunkName string) {
	proto.Source = chunkName
	for _, f := range proto.Protos {
		setSource(f, chunkName)
	}
}
