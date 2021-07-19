// Code generated from /home/lemonhx/Desktop/Go/LogParser/UNIFORM_LOG.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // UNIFORM_LOG

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseUNIFORM_LOGVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseUNIFORM_LOGVisitor) VisitLogs(ctx *LogsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUNIFORM_LOGVisitor) VisitLog(ctx *LogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUNIFORM_LOGVisitor) VisitDate_time(ctx *Date_timeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUNIFORM_LOGVisitor) VisitLevel(ctx *LevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUNIFORM_LOGVisitor) VisitLocation(ctx *LocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUNIFORM_LOGVisitor) VisitMessage(ctx *MessageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUNIFORM_LOGVisitor) VisitKv_pair(ctx *Kv_pairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUNIFORM_LOGVisitor) VisitKey(ctx *KeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUNIFORM_LOGVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}
