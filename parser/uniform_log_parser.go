// Code generated from /Users/lemonhx/Desktop/Go/ti-uniform-log/parser/UNIFORM_LOG.g4 by ANTLR 4.9.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 80, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2, 3,
	2, 3, 2, 7, 2, 28, 10, 2, 12, 2, 14, 2, 31, 11, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 7, 3, 38, 10, 3, 12, 3, 14, 3, 41, 11, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5,
	6, 58, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 2, 2, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2, 4, 3, 2, 16,
	17, 6, 2, 9, 9, 11, 11, 13, 13, 16, 17, 2, 71, 2, 24, 3, 2, 2, 2, 4, 32,
	3, 2, 2, 2, 6, 42, 3, 2, 2, 2, 8, 46, 3, 2, 2, 2, 10, 57, 3, 2, 2, 2, 12,
	59, 3, 2, 2, 2, 14, 63, 3, 2, 2, 2, 16, 67, 3, 2, 2, 2, 18, 71, 3, 2, 2,
	2, 20, 75, 3, 2, 2, 2, 22, 77, 3, 2, 2, 2, 24, 29, 5, 4, 3, 2, 25, 26,
	7, 3, 2, 2, 26, 28, 5, 4, 3, 2, 27, 25, 3, 2, 2, 2, 28, 31, 3, 2, 2, 2,
	29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 3, 3, 2, 2, 2, 31, 29, 3, 2,
	2, 2, 32, 33, 5, 6, 4, 2, 33, 34, 5, 8, 5, 2, 34, 35, 5, 10, 6, 2, 35,
	39, 5, 12, 7, 2, 36, 38, 5, 14, 8, 2, 37, 36, 3, 2, 2, 2, 38, 41, 3, 2,
	2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 5, 3, 2, 2, 2, 41, 39,
	3, 2, 2, 2, 42, 43, 7, 4, 2, 2, 43, 44, 7, 10, 2, 2, 44, 45, 7, 5, 2, 2,
	45, 7, 3, 2, 2, 2, 46, 47, 7, 4, 2, 2, 47, 48, 7, 14, 2, 2, 48, 49, 7,
	5, 2, 2, 49, 9, 3, 2, 2, 2, 50, 51, 7, 4, 2, 2, 51, 52, 7, 6, 2, 2, 52,
	58, 7, 5, 2, 2, 53, 54, 7, 4, 2, 2, 54, 55, 5, 16, 9, 2, 55, 56, 7, 5,
	2, 2, 56, 58, 3, 2, 2, 2, 57, 50, 3, 2, 2, 2, 57, 53, 3, 2, 2, 2, 58, 11,
	3, 2, 2, 2, 59, 60, 7, 4, 2, 2, 60, 61, 5, 20, 11, 2, 61, 62, 7, 5, 2,
	2, 62, 13, 3, 2, 2, 2, 63, 64, 7, 4, 2, 2, 64, 65, 5, 18, 10, 2, 65, 66,
	7, 5, 2, 2, 66, 15, 3, 2, 2, 2, 67, 68, 7, 15, 2, 2, 68, 69, 7, 7, 2, 2,
	69, 70, 7, 11, 2, 2, 70, 17, 3, 2, 2, 2, 71, 72, 5, 20, 11, 2, 72, 73,
	7, 8, 2, 2, 73, 74, 5, 22, 12, 2, 74, 19, 3, 2, 2, 2, 75, 76, 9, 2, 2,
	2, 76, 21, 3, 2, 2, 2, 77, 78, 9, 3, 2, 2, 78, 23, 3, 2, 2, 2, 5, 29, 39,
	57,
}
var literalNames = []string{
	"", "'\n'", "'['", "']'", "'<unknown>'", "':'", "'='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "IP", "DATE_TIME", "NUM", "FLOAT", "DURATION",
	"LEVEL", "PATHLOC", "STRING", "IDENT", "WS",
}

var ruleNames = []string{
	"logs", "log", "date_time", "level", "location", "custom_message", "message",
	"known_location", "kv_pair", "key", "value",
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
	UNIFORM_LOGParserT__5      = 6
	UNIFORM_LOGParserIP        = 7
	UNIFORM_LOGParserDATE_TIME = 8
	UNIFORM_LOGParserNUM       = 9
	UNIFORM_LOGParserFLOAT     = 10
	UNIFORM_LOGParserDURATION  = 11
	UNIFORM_LOGParserLEVEL     = 12
	UNIFORM_LOGParserPATHLOC   = 13
	UNIFORM_LOGParserSTRING    = 14
	UNIFORM_LOGParserIDENT     = 15
	UNIFORM_LOGParserWS        = 16
)

// UNIFORM_LOGParser rules.
const (
	UNIFORM_LOGParserRULE_logs           = 0
	UNIFORM_LOGParserRULE_log            = 1
	UNIFORM_LOGParserRULE_date_time      = 2
	UNIFORM_LOGParserRULE_level          = 3
	UNIFORM_LOGParserRULE_location       = 4
	UNIFORM_LOGParserRULE_custom_message = 5
	UNIFORM_LOGParserRULE_message        = 6
	UNIFORM_LOGParserRULE_known_location = 7
	UNIFORM_LOGParserRULE_kv_pair        = 8
	UNIFORM_LOGParserRULE_key            = 9
	UNIFORM_LOGParserRULE_value          = 10
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
		p.SetState(22)
		p.Log()
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == UNIFORM_LOGParserT__0 {
		{
			p.SetState(23)
			p.Match(UNIFORM_LOGParserT__0)
		}
		{
			p.SetState(24)
			p.Log()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *LogContext) Custom_message() ICustom_messageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICustom_messageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICustom_messageContext)
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
		p.SetState(30)
		p.Date_time()
	}
	{
		p.SetState(31)
		p.Level()
	}
	{
		p.SetState(32)
		p.Location()
	}
	{
		p.SetState(33)
		p.Custom_message()
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == UNIFORM_LOGParserT__1 {
		{
			p.SetState(34)
			p.Message()
		}

		p.SetState(39)
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
		p.SetState(40)
		p.Match(UNIFORM_LOGParserT__1)
	}
	{
		p.SetState(41)
		p.Match(UNIFORM_LOGParserDATE_TIME)
	}
	{
		p.SetState(42)
		p.Match(UNIFORM_LOGParserT__2)
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
		p.SetState(44)
		p.Match(UNIFORM_LOGParserT__1)
	}
	{
		p.SetState(45)
		p.Match(UNIFORM_LOGParserLEVEL)
	}
	{
		p.SetState(46)
		p.Match(UNIFORM_LOGParserT__2)
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

func (s *LocationContext) Known_location() IKnown_locationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKnown_locationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKnown_locationContext)
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

	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Match(UNIFORM_LOGParserT__1)
		}
		{
			p.SetState(49)
			p.Match(UNIFORM_LOGParserT__3)
		}
		{
			p.SetState(50)
			p.Match(UNIFORM_LOGParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Match(UNIFORM_LOGParserT__1)
		}
		{
			p.SetState(52)
			p.Known_location()
		}
		{
			p.SetState(53)
			p.Match(UNIFORM_LOGParserT__2)
		}

	}

	return localctx
}

// ICustom_messageContext is an interface to support dynamic dispatch.
type ICustom_messageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCustom_messageContext differentiates from other interfaces.
	IsCustom_messageContext()
}

type Custom_messageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCustom_messageContext() *Custom_messageContext {
	var p = new(Custom_messageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UNIFORM_LOGParserRULE_custom_message
	return p
}

func (*Custom_messageContext) IsCustom_messageContext() {}

func NewCustom_messageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Custom_messageContext {
	var p = new(Custom_messageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UNIFORM_LOGParserRULE_custom_message

	return p
}

func (s *Custom_messageContext) GetParser() antlr.Parser { return s.parser }

func (s *Custom_messageContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Custom_messageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Custom_messageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Custom_messageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.EnterCustom_message(s)
	}
}

func (s *Custom_messageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.ExitCustom_message(s)
	}
}

func (s *Custom_messageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UNIFORM_LOGVisitor:
		return t.VisitCustom_message(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *UNIFORM_LOGParser) Custom_message() (localctx ICustom_messageContext) {
	localctx = NewCustom_messageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, UNIFORM_LOGParserRULE_custom_message)

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
		p.SetState(57)
		p.Match(UNIFORM_LOGParserT__1)
	}
	{
		p.SetState(58)
		p.Key()
	}
	{
		p.SetState(59)
		p.Match(UNIFORM_LOGParserT__2)
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
	p.EnterRule(localctx, 12, UNIFORM_LOGParserRULE_message)

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
		p.SetState(61)
		p.Match(UNIFORM_LOGParserT__1)
	}
	{
		p.SetState(62)
		p.Kv_pair()
	}
	{
		p.SetState(63)
		p.Match(UNIFORM_LOGParserT__2)
	}

	return localctx
}

// IKnown_locationContext is an interface to support dynamic dispatch.
type IKnown_locationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKnown_locationContext differentiates from other interfaces.
	IsKnown_locationContext()
}

type Known_locationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKnown_locationContext() *Known_locationContext {
	var p = new(Known_locationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UNIFORM_LOGParserRULE_known_location
	return p
}

func (*Known_locationContext) IsKnown_locationContext() {}

func NewKnown_locationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Known_locationContext {
	var p = new(Known_locationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UNIFORM_LOGParserRULE_known_location

	return p
}

func (s *Known_locationContext) GetParser() antlr.Parser { return s.parser }

func (s *Known_locationContext) PATHLOC() antlr.TerminalNode {
	return s.GetToken(UNIFORM_LOGParserPATHLOC, 0)
}

func (s *Known_locationContext) NUM() antlr.TerminalNode {
	return s.GetToken(UNIFORM_LOGParserNUM, 0)
}

func (s *Known_locationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Known_locationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Known_locationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.EnterKnown_location(s)
	}
}

func (s *Known_locationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UNIFORM_LOGListener); ok {
		listenerT.ExitKnown_location(s)
	}
}

func (s *Known_locationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UNIFORM_LOGVisitor:
		return t.VisitKnown_location(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *UNIFORM_LOGParser) Known_location() (localctx IKnown_locationContext) {
	localctx = NewKnown_locationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, UNIFORM_LOGParserRULE_known_location)

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
		p.SetState(65)
		p.Match(UNIFORM_LOGParserPATHLOC)
	}
	{
		p.SetState(66)
		p.Match(UNIFORM_LOGParserT__4)
	}
	{
		p.SetState(67)
		p.Match(UNIFORM_LOGParserNUM)
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
	p.EnterRule(localctx, 16, UNIFORM_LOGParserRULE_kv_pair)

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
		p.Key()
	}
	{
		p.SetState(70)
		p.Match(UNIFORM_LOGParserT__5)
	}
	{
		p.SetState(71)
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
	p.EnterRule(localctx, 18, UNIFORM_LOGParserRULE_key)
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
		p.SetState(73)
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
	p.EnterRule(localctx, 20, UNIFORM_LOGParserRULE_value)
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
		p.SetState(75)
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
