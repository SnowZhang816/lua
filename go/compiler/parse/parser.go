package parse

import "main/compiler/ast"
import "main/compiler/lexer"

/* recursive descent parser */

func Parse(chunk, chunkName string) *ast.Block {
	lex := lexer.NewLexer(chunk, chunkName)
	block := parseBlock(lex)
	lex.NextTokenOfKind(lexer.TOKEN_EOF)
	return block
}
