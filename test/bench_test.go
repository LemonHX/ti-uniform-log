package test

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
	"unified_log/parser"
)

func BenchmarkLogParser(b *testing.B) {
	b.ResetTimer()
	input, _ := antlr.NewFileStream("../test.log")
	lexer := parser.NewUNIFORM_LOGLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewUNIFORM_LOGParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	p.Logs()
}
