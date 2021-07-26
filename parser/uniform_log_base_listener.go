// Code generated from /Users/lemonhx/Desktop/Go/ti-uniform-log/parser/UNIFORM_LOG.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // UNIFORM_LOG

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseUNIFORM_LOGListener is a complete listener for a parse tree produced by UNIFORM_LOGParser.
type BaseUNIFORM_LOGListener struct{}

var _ UNIFORM_LOGListener = &BaseUNIFORM_LOGListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseUNIFORM_LOGListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseUNIFORM_LOGListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseUNIFORM_LOGListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseUNIFORM_LOGListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLogs is called when production logs is entered.
func (s *BaseUNIFORM_LOGListener) EnterLogs(ctx *LogsContext) {}

// ExitLogs is called when production logs is exited.
func (s *BaseUNIFORM_LOGListener) ExitLogs(ctx *LogsContext) {}

// EnterLog is called when production log is entered.
func (s *BaseUNIFORM_LOGListener) EnterLog(ctx *LogContext) {}

// ExitLog is called when production log is exited.
func (s *BaseUNIFORM_LOGListener) ExitLog(ctx *LogContext) {}

// EnterDate_time is called when production date_time is entered.
func (s *BaseUNIFORM_LOGListener) EnterDate_time(ctx *Date_timeContext) {}

// ExitDate_time is called when production date_time is exited.
func (s *BaseUNIFORM_LOGListener) ExitDate_time(ctx *Date_timeContext) {}

// EnterLevel is called when production level is entered.
func (s *BaseUNIFORM_LOGListener) EnterLevel(ctx *LevelContext) {}

// ExitLevel is called when production level is exited.
func (s *BaseUNIFORM_LOGListener) ExitLevel(ctx *LevelContext) {}

// EnterLocation is called when production location is entered.
func (s *BaseUNIFORM_LOGListener) EnterLocation(ctx *LocationContext) {}

// ExitLocation is called when production location is exited.
func (s *BaseUNIFORM_LOGListener) ExitLocation(ctx *LocationContext) {}

// EnterCustom_message is called when production custom_message is entered.
func (s *BaseUNIFORM_LOGListener) EnterCustom_message(ctx *Custom_messageContext) {}

// ExitCustom_message is called when production custom_message is exited.
func (s *BaseUNIFORM_LOGListener) ExitCustom_message(ctx *Custom_messageContext) {}

// EnterMessage is called when production message is entered.
func (s *BaseUNIFORM_LOGListener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production message is exited.
func (s *BaseUNIFORM_LOGListener) ExitMessage(ctx *MessageContext) {}

// EnterKnown_location is called when production known_location is entered.
func (s *BaseUNIFORM_LOGListener) EnterKnown_location(ctx *Known_locationContext) {}

// ExitKnown_location is called when production known_location is exited.
func (s *BaseUNIFORM_LOGListener) ExitKnown_location(ctx *Known_locationContext) {}

// EnterKv_pair is called when production kv_pair is entered.
func (s *BaseUNIFORM_LOGListener) EnterKv_pair(ctx *Kv_pairContext) {}

// ExitKv_pair is called when production kv_pair is exited.
func (s *BaseUNIFORM_LOGListener) ExitKv_pair(ctx *Kv_pairContext) {}

// EnterKey is called when production key is entered.
func (s *BaseUNIFORM_LOGListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseUNIFORM_LOGListener) ExitKey(ctx *KeyContext) {}

// EnterValue is called when production value is entered.
func (s *BaseUNIFORM_LOGListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseUNIFORM_LOGListener) ExitValue(ctx *ValueContext) {}
