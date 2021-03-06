// Code generated from /Users/lemonhx/Desktop/Go/ti-uniform-log/parser/UNIFORM_LOG.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // UNIFORM_LOG

import "github.com/antlr/antlr4/runtime/Go/antlr"

// UNIFORM_LOGListener is a complete listener for a parse tree produced by UNIFORM_LOGParser.
type UNIFORM_LOGListener interface {
	antlr.ParseTreeListener

	// EnterLogs is called when entering the logs production.
	EnterLogs(c *LogsContext)

	// EnterLog is called when entering the log production.
	EnterLog(c *LogContext)

	// EnterDate_time is called when entering the date_time production.
	EnterDate_time(c *Date_timeContext)

	// EnterLevel is called when entering the level production.
	EnterLevel(c *LevelContext)

	// EnterLocation is called when entering the location production.
	EnterLocation(c *LocationContext)

	// EnterCustom_message is called when entering the custom_message production.
	EnterCustom_message(c *Custom_messageContext)

	// EnterMessage is called when entering the message production.
	EnterMessage(c *MessageContext)

	// EnterKnown_location is called when entering the known_location production.
	EnterKnown_location(c *Known_locationContext)

	// EnterKv_pair is called when entering the kv_pair production.
	EnterKv_pair(c *Kv_pairContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitLogs is called when exiting the logs production.
	ExitLogs(c *LogsContext)

	// ExitLog is called when exiting the log production.
	ExitLog(c *LogContext)

	// ExitDate_time is called when exiting the date_time production.
	ExitDate_time(c *Date_timeContext)

	// ExitLevel is called when exiting the level production.
	ExitLevel(c *LevelContext)

	// ExitLocation is called when exiting the location production.
	ExitLocation(c *LocationContext)

	// ExitCustom_message is called when exiting the custom_message production.
	ExitCustom_message(c *Custom_messageContext)

	// ExitMessage is called when exiting the message production.
	ExitMessage(c *MessageContext)

	// ExitKnown_location is called when exiting the known_location production.
	ExitKnown_location(c *Known_locationContext)

	// ExitKv_pair is called when exiting the kv_pair production.
	ExitKv_pair(c *Kv_pairContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
