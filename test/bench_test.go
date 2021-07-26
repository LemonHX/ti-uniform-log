package test

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
	"unified_log/parser"
	"unified_log/transformer"
)

func BenchmarkLogParser(b *testing.B) {
	b.ResetTimer()
	input, _ := antlr.NewFileStream("../test.log")
	lexer := parser.NewUNIFORM_LOGLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewUNIFORM_LOGParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	ctx := p.Logs()
	logs := ctx.GetChildren()
	for i := 0; i < ctx.GetChildCount(); i += 1 {
		log := logs[i]
		switch log.GetPayload().(type) {
		case *antlr.BaseParserRuleContext:
			payload := log.GetPayload().(*antlr.BaseParserRuleContext).GetChildren()
			transformer.ToJson(payload)
			break
		default:
			//println(t)
			continue
		}
	}
}
