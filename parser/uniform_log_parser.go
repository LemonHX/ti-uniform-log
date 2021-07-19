// Code generated from /home/lemonhx/Desktop/Go/LogParser/UNIFORM_LOG.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // UNIFORM_LOG

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 74, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 7, 2, 24, 10, 2, 12,
	2, 14, 2, 27, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 35, 10,
	3, 12, 3, 14, 3, 38, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 54, 10, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 64, 10, 7, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 2, 4, 3, 2, 15, 16, 6, 2, 8, 8, 10, 10, 12, 12, 15, 16, 2, 68, 2, 20,
	3, 2, 2, 2, 4, 30, 3, 2, 2, 2, 6, 39, 3, 2, 2, 2, 8, 43, 3, 2, 2, 2, 10,
	53, 3, 2, 2, 2, 12, 63, 3, 2, 2, 2, 14, 65, 3, 2, 2, 2, 16, 69, 3, 2, 2,
	2, 18, 71, 3, 2, 2, 2, 20, 25, 5, 4, 3, 2, 21, 22, 7, 18, 2, 2, 22, 24,
	5, 4, 3, 2, 23, 21, 3, 2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2,
	25, 26, 3, 2, 2, 2, 26, 28, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 28, 29, 7,
	18, 2, 2, 29, 3, 3, 2, 2, 2, 30, 31, 5, 6, 4, 2, 31, 32, 5, 8, 5, 2, 32,
	36, 5, 10, 6, 2, 33, 35, 5, 12, 7, 2, 34, 33, 3, 2, 2, 2, 35, 38, 3, 2,
	2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 5, 3, 2, 2, 2, 38, 36,
	3, 2, 2, 2, 39, 40, 7, 3, 2, 2, 40, 41, 7, 9, 2, 2, 41, 42, 7, 4, 2, 2,
	42, 7, 3, 2, 2, 2, 43, 44, 7, 3, 2, 2, 44, 45, 7, 13, 2, 2, 45, 46, 7,
	4, 2, 2, 46, 9, 3, 2, 2, 2, 47, 54, 7, 5, 2, 2, 48, 49, 7, 3, 2, 2, 49,
	50, 7, 14, 2, 2, 50, 51, 7, 6, 2, 2, 51, 52, 7, 10, 2, 2, 52, 54, 7, 4,
	2, 2, 53, 47, 3, 2, 2, 2, 53, 48, 3, 2, 2, 2, 54, 11, 3, 2, 2, 2, 55, 56,
	7, 3, 2, 2, 56, 57, 5, 18, 10, 2, 57, 58, 7, 4, 2, 2, 58, 64, 3, 2, 2,
	2, 59, 60, 7, 3, 2, 2, 60, 61, 5, 14, 8, 2, 61, 62, 7, 4, 2, 2, 62, 64,
	3, 2, 2, 2, 63, 55, 3, 2, 2, 2, 63, 59, 3, 2, 2, 2, 64, 13, 3, 2, 2, 2,
	65, 66, 5, 16, 9, 2, 66, 67, 7, 7, 2, 2, 67, 68, 5, 18, 10, 2, 68, 15,
	3, 2, 2, 2, 69, 70, 9, 2, 2, 2, 70, 17, 3, 2, 2, 2, 71, 72, 9, 3, 2, 2,
	72, 19, 3, 2, 2, 2, 6, 25, 36, 53, 63,
}
var literalNames = []string{
	"", "'['", "']'", "'[<unknown>]'", "':'", "'='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "IP", "DATE_TIME", "NUM", "FLOAT", "DURATION",
	"LEVEL", "PATHLOC", "STRING", "IDENT", "WS", "NL",
}

var ruleNames = []string{
	"logs", "log", "date_time", "level", "location", "message", "kv_pair",
	"key", "value",
}

type UNIFORM_LOGParser struct {
	*antlr.BaseParser
}

// NewUNIFORM_LOGParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *UNIFORM_LOGParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewUNIFORM_LOGParser(input antlr.TokenStream) *UNIFORM_LOGParser {
	this := new(UNIFORM_LOGParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "UNIFORM_LOG.g4"

	return this
}

// UNIFORM_LOGParser tokens.
const (
	UNIFORM_LOGParserEOF       = antlr.TokenEOF
	UNIFORM_LOGParserT__0      = 1
	UNIFORM_LOGParserT__1      = 2
	UNIFORM_LOGParserT__2      = 3
	UNIFORM_LOGParserT__3      = 4
	UNIFORM_LOGParserT__4      = 5
	UNIFORM_LOGParserIP        = 6
	UNIFORM_LOGParserDATE_TIME = 7
	UNIFORM_LOGParserNUM       = 8
	UNIFORM_LOGParserFLOAT     = 9
	UNIFORM_LOGParserDURATION  = 10
	UNIFORM_LOGParserLEVEL     = 11
	UNIFORM_LOGParserPATHLOC   = 12
	UNIFORM_LOGParserSTRING    = 13
	UNIFORM_LOGParserIDENT     = 14
	UNIFORM_LOGParserWS        = 15
	UNIFORM_LOGParserNL        = 16
)

// UNIFORM_LOGParser rules.
const (
	UNIFORM_LOGParserRULE_logs      = 0
	UNIFORM_LOGParserRULE_log       = 1
	UNIFORM_LOGParserRULE_date_time = 2
	UNIFORM_LOGParserRULE_level     = 3
	UNIFORM_LOGParserRULE_location  = 4
	UNIFORM_LOGParserRULE_message   = 5
	UNIFORM_LOGParserRULE_kv_pair   = 6
	UNIFORM_LOGParserRULE_key       = 7
	UNIFORM_LOGParserRULE_value     = 8
)

// ILogsContext is an interface to support dynamic dispatch.
type ILogsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogsContext differentiates from other interfaces.
	IsLogsContext()
}

type LogsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogsContext() *LogsContext {
	var p = new(LogsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UNIFORM_LOGParserRULE_logs
	return p
}

func (*LogsContext) IsLogsContext() {}

func NewLogsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogsContext {
	var p = new(LogsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UNIFORM_LOGParserRULE_logs

	return p
}

func (s *LogsContext) GetParser() antlr.Parser { return s.parser }

func (s *LogsContext) AllLog() []ILogContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILogContext)(nil)).Elem())
	var tst = make([]ILogContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILogContext)
		}
	}

	return tst
}

func (s *LogsContext) Log(i int) ILogContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILogContext)
}

func (s *LogsContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(UNIFORM_LOGParserNL)
}

func (s *LogsContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(UNIFORM_LOGParserNL, i)
}

func (s *LogsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.EnterLogs(s)
	}
}

func (s *LogsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.ExitLogs(s)
	}
}

func (s *LogsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UNIFORM_LOGVisitor:
		return t.VisitLogs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *UNIFORM_LOGParser) Logs() (localctx ILogsContext) {
	localctx = NewLogsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, UNIFORM_LOGParserRULE_logs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Log()
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(19)
				p.Match(UNIFORM_LOGParserNL)
			}
			{
				p.SetState(20)
				p.Log()
			}

		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(26)
		p.Match(UNIFORM_LOGParserNL)
	}

	return localctx
}

// ILogContext is an interface to support dynamic dispatch.
type ILogContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogContext differentiates from other interfaces.
	IsLogContext()
}

type LogContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogContext() *LogContext {
	var p = new(LogContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UNIFORM_LOGParserRULE_log
	return p
}

func (*LogContext) IsLogContext() {}

func NewLogContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogContext {
	var p = new(LogContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UNIFORM_LOGParserRULE_log

	return p
}

func (s *LogContext) GetParser() antlr.Parser { return s.parser }

func (s *LogContext) Date_time() IDate_timeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_timeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_timeContext)
}

func (s *LogContext) Level() ILevelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILevelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILevelContext)
}

func (s *LogContext) Location() ILocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocationContext)
}

func (s *LogContext) AllMessage() []IMessageContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageContext)(nil)).Elem())
	var tst = make([]IMessageContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageContext)
		}
	}

	return tst
}

func (s *LogContext) Message(i int) IMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *LogContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.EnterLog(s)
	}
}

func (s *LogContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.ExitLog(s)
	}
}

func (s *LogContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UNIFORM_LOGVisitor:
		return t.VisitLog(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *UNIFORM_LOGParser) Log() (localctx ILogContext) {
	localctx = NewLogContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, UNIFORM_LOGParserRULE_log)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Date_time()
	}
	{
		p.SetState(29)
		p.Level()
	}
	{
		p.SetState(30)
		p.Location()
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == UNIFORM_LOGParserT__0 {
		{
			p.SetState(31)
			p.Message()
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDate_timeContext is an interface to support dynamic dispatch.
type IDate_timeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDate_timeContext differentiates from other interfaces.
	IsDate_timeContext()
}

type Date_timeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_timeContext() *Date_timeContext {
	var p = new(Date_timeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UNIFORM_LOGParserRULE_date_time
	return p
}

func (*Date_timeContext) IsDate_timeContext() {}

func NewDate_timeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_timeContext {
	var p = new(Date_timeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UNIFORM_LOGParserRULE_date_time

	return p
}

func (s *Date_timeContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_timeContext) DATE_TIME() antlr.TerminalNode {
	return s.GetToken(UNIFORM_LOGParserDATE_TIME, 0)
}

func (s *Date_timeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_timeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Date_timeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.EnterDate_time(s)
	}
}

func (s *Date_timeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.ExitDate_time(s)
	}
}

func (s *Date_timeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UNIFORM_LOGVisitor:
		return t.VisitDate_time(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *UNIFORM_LOGParser) Date_time() (localctx IDate_timeContext) {
	localctx = NewDate_timeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, UNIFORM_LOGParserRULE_date_time)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Match(UNIFORM_LOGParserT__0)
	}
	{
		p.SetState(38)
		p.Match(UNIFORM_LOGParserDATE_TIME)
	}
	{
		p.SetState(39)
		p.Match(UNIFORM_LOGParserT__1)
	}

	return localctx
}

// ILevelContext is an interface to support dynamic dispatch.
type ILevelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLevelContext differentiates from other interfaces.
	IsLevelContext()
}

type LevelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLevelContext() *LevelContext {
	var p = new(LevelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UNIFORM_LOGParserRULE_level
	return p
}

func (*LevelContext) IsLevelContext() {}

func NewLevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LevelContext {
	var p = new(LevelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UNIFORM_LOGParserRULE_level

	return p
}

func (s *LevelContext) GetParser() antlr.Parser { return s.parser }

func (s *LevelContext) LEVEL() antlr.TerminalNode {
	return s.GetToken(UNIFORM_LOGParserLEVEL, 0)
}

func (s *LevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LevelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.EnterLevel(s)
	}
}

func (s *LevelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.ExitLevel(s)
	}
}

func (s *LevelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UNIFORM_LOGVisitor:
		return t.VisitLevel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *UNIFORM_LOGParser) Level() (localctx ILevelContext) {
	localctx = NewLevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, UNIFORM_LOGParserRULE_level)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(UNIFORM_LOGParserT__0)
	}
	{
		p.SetState(42)
		p.Match(UNIFORM_LOGParserLEVEL)
	}
	{
		p.SetState(43)
		p.Match(UNIFORM_LOGParserT__1)
	}

	return localctx
}

// ILocationContext is an interface to support dynamic dispatch.
type ILocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocationContext differentiates from other interfaces.
	IsLocationContext()
}

type LocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocationContext() *LocationContext {
	var p = new(LocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UNIFORM_LOGParserRULE_location
	return p
}

func (*LocationContext) IsLocationContext() {}

func NewLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocationContext {
	var p = new(LocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UNIFORM_LOGParserRULE_location

	return p
}

func (s *LocationContext) GetParser() antlr.Parser { return s.parser }

func (s *LocationContext) PATHLOC() antlr.TerminalNode {
	return s.GetToken(UNIFORM_LOGParserPATHLOC, 0)
}

func (s *LocationContext) NUM() antlr.TerminalNode {
	return s.GetToken(UNIFORM_LOGParserNUM, 0)
}

func (s *LocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.EnterLocation(s)
	}
}

func (s *LocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.ExitLocation(s)
	}
}

func (s *LocationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UNIFORM_LOGVisitor:
		return t.VisitLocation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *UNIFORM_LOGParser) Location() (localctx ILocationContext) {
	localctx = NewLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, UNIFORM_LOGParserRULE_location)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case UNIFORM_LOGParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(UNIFORM_LOGParserT__2)
		}

	case UNIFORM_LOGParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)
			p.Match(UNIFORM_LOGParserT__0)
		}
		{
			p.SetState(47)
			p.Match(UNIFORM_LOGParserPATHLOC)
		}
		{
			p.SetState(48)
			p.Match(UNIFORM_LOGParserT__3)
		}
		{
			p.SetState(49)
			p.Match(UNIFORM_LOGParserNUM)
		}
		{
			p.SetState(50)
			p.Match(UNIFORM_LOGParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMessageContext is an interface to support dynamic dispatch.
type IMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageContext differentiates from other interfaces.
	IsMessageContext()
}

type MessageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageContext() *MessageContext {
	var p = new(MessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UNIFORM_LOGParserRULE_message
	return p
}

func (*MessageContext) IsMessageContext() {}

func NewMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContext {
	var p = new(MessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UNIFORM_LOGParserRULE_message

	return p
}

func (s *MessageContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *MessageContext) Kv_pair() IKv_pairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKv_pairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKv_pairContext)
}

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.ExitMessage(s)
	}
}

func (s *MessageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UNIFORM_LOGVisitor:
		return t.VisitMessage(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *UNIFORM_LOGParser) Message() (localctx IMessageContext) {
	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, UNIFORM_LOGParserRULE_message)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Match(UNIFORM_LOGParserT__0)
		}
		{
			p.SetState(54)
			p.Value()
		}
		{
			p.SetState(55)
			p.Match(UNIFORM_LOGParserT__1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Match(UNIFORM_LOGParserT__0)
		}
		{
			p.SetState(58)
			p.Kv_pair()
		}
		{
			p.SetState(59)
			p.Match(UNIFORM_LOGParserT__1)
		}

	}

	return localctx
}

// IKv_pairContext is an interface to support dynamic dispatch.
type IKv_pairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKv_pairContext differentiates from other interfaces.
	IsKv_pairContext()
}

type Kv_pairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKv_pairContext() *Kv_pairContext {
	var p = new(Kv_pairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UNIFORM_LOGParserRULE_kv_pair
	return p
}

func (*Kv_pairContext) IsKv_pairContext() {}

func NewKv_pairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kv_pairContext {
	var p = new(Kv_pairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UNIFORM_LOGParserRULE_kv_pair

	return p
}

func (s *Kv_pairContext) GetParser() antlr.Parser { return s.parser }

func (s *Kv_pairContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Kv_pairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Kv_pairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kv_pairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kv_pairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.EnterKv_pair(s)
	}
}

func (s *Kv_pairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.ExitKv_pair(s)
	}
}

func (s *Kv_pairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UNIFORM_LOGVisitor:
		return t.VisitKv_pair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *UNIFORM_LOGParser) Kv_pair() (localctx IKv_pairContext) {
	localctx = NewKv_pairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, UNIFORM_LOGParserRULE_kv_pair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Key()
	}
	{
		p.SetState(64)
		p.Match(UNIFORM_LOGParserT__4)
	}
	{
		p.SetState(65)
		p.Value()
	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UNIFORM_LOGParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UNIFORM_LOGParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) IDENT() antlr.TerminalNode {
	return s.GetToken(UNIFORM_LOGParserIDENT, 0)
}

func (s *KeyContext) STRING() antlr.TerminalNode {
	return s.GetToken(UNIFORM_LOGParserSTRING, 0)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.ExitKey(s)
	}
}

func (s *KeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UNIFORM_LOGVisitor:
		return t.VisitKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *UNIFORM_LOGParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, UNIFORM_LOGParserRULE_key)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		_la = p.GetTokenStream().LA(1)

		if !(_la == UNIFORM_LOGParserSTRING || _la == UNIFORM_LOGParserIDENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UNIFORM_LOGParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UNIFORM_LOGParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) IDENT() antlr.TerminalNode {
	return s.GetToken(UNIFORM_LOGParserIDENT, 0)
}

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(UNIFORM_LOGParserSTRING, 0)
}

func (s *ValueContext) NUM() antlr.TerminalNode {
	return s.GetToken(UNIFORM_LOGParserNUM, 0)
}

func (s *ValueContext) DURATION() antlr.TerminalNode {
	return s.GetToken(UNIFORM_LOGParserDURATION, 0)
}

func (s *ValueContext) IP() antlr.TerminalNode {
	return s.GetToken(UNIFORM_LOGParserIP, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UNIFORM_LOGVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *UNIFORM_LOGParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, UNIFORM_LOGParserRULE_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<UNIFORM_LOGParserIP)|(1<<UNIFORM_LOGParserNUM)|(1<<UNIFORM_LOGParserDURATION)|(1<<UNIFORM_LOGParserSTRING)|(1<<UNIFORM_LOGParserIDENT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
