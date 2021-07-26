// Code generated from /Users/lemonhx/Desktop/Go/ti-uniform-log/parser/UNIFORM_LOG.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // UNIFORM_LOG

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by UNIFORM_LOGParser.
type UNIFORM_LOGVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by UNIFORM_LOGParser#logs.
	VisitLogs(ctx *LogsContext) interface{}

	// Visit a parse tree produced by UNIFORM_LOGParser#log.
	VisitLog(ctx *LogContext) interface{}

	// Visit a parse tree produced by UNIFORM_LOGParser#date_time.
	VisitDate_time(ctx *Date_timeContext) interface{}

	// Visit a parse tree produced by UNIFORM_LOGParser#level.
	VisitLevel(ctx *LevelContext) interface{}

	// Visit a parse tree produced by UNIFORM_LOGParser#location.
	VisitLocation(ctx *LocationContext) interface{}

	// Visit a parse tree produced by UNIFORM_LOGParser#custom_message.
	VisitCustom_message(ctx *Custom_messageContext) interface{}

	// Visit a parse tree produced by UNIFORM_LOGParser#message.
	VisitMessage(ctx *MessageContext) interface{}

	// Visit a parse tree produced by UNIFORM_LOGParser#known_location.
	VisitKnown_location(ctx *Known_locationContext) interface{}

	// Visit a parse tree produced by UNIFORM_LOGParser#kv_pair.
	VisitKv_pair(ctx *Kv_pairContext) interface{}

	// Visit a parse tree produced by UNIFORM_LOGParser#key.
	VisitKey(ctx *KeyContext) interface{}

	// Visit a parse tree produced by UNIFORM_LOGParser#value.
	VisitValue(ctx *ValueContext) interface{}
}
