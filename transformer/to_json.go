package transformer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
	"unified_log/parser"
)

func MessagesToJson(messages []antlr.Tree) string {
	var result_builder strings.Builder
	result_builder.WriteByte('{')
	for i := 0; i < len(messages); i++ {
		payload := messages[i].(*parser.MessageContext).GetChild(1)
		kv := payload.(*parser.Kv_pairContext)
		k := kv.GetChild(0).(*parser.KeyContext).GetText()
		if k[0] != '"' && k[len(k)-1] != '"' {
			var sb strings.Builder
			sb.WriteByte('"')
			sb.WriteString(k)
			sb.WriteByte('"')
			k = sb.String()
		}
		v := kv.GetChild(2).(*parser.ValueContext).GetText()
		if v[0] != '"' && v[len(v)-1] != '"' {
			var sb strings.Builder
			sb.WriteByte('"')
			sb.WriteString(v)
			sb.WriteByte('"')
			v = sb.String()
		}
		result_builder.WriteString(k)
		result_builder.WriteByte(':')
		result_builder.WriteString(v)
		if i != len(messages)-1 {
			result_builder.WriteString(", ")
		}
	}
	result_builder.WriteByte('}')
	//fmt.Println("Result: ",result_builder.String())
	return result_builder.String()
}

func ToJson(payload []antlr.Tree) string {
	var result_builder strings.Builder
	result_builder.WriteByte('{')
	date_time := payload[0].(*parser.Date_timeContext).GetChild(1).GetPayload().(*antlr.CommonToken).GetText()
	result_builder.WriteString("\"date_time\": ")
	result_builder.WriteByte('"')
	result_builder.WriteString(date_time)
	result_builder.WriteByte('"')
	result_builder.WriteString(", ")

	level := payload[1].(*parser.LevelContext).GetChild(1).(*antlr.TerminalNodeImpl).GetText()
	result_builder.WriteString("\"level\": ")
	result_builder.WriteByte('"')
	result_builder.WriteString(level)
	result_builder.WriteByte('"')
	result_builder.WriteString(", ")

	location_case := payload[2].(*parser.LocationContext).GetChild(1)
	var location string
	switch location_case.(type) {
	case *parser.Known_locationContext:
		location = location_case.(*parser.Known_locationContext).GetText()
	case *antlr.TerminalNodeImpl:
		location = location_case.(*antlr.TerminalNodeImpl).GetText()
	}
	result_builder.WriteString("\"location\": ")
	result_builder.WriteByte('"')
	result_builder.WriteString(location)
	result_builder.WriteByte('"')
	result_builder.WriteString(", ")

	custom_message := payload[3].(*parser.Custom_messageContext).GetChild(1).(*parser.KeyContext).GetText()
	result_builder.WriteString("\"custom_message\": ")
	if custom_message[0] != '"' && custom_message[len(custom_message)-1] != '"' {
		result_builder.WriteByte('"')
		result_builder.WriteString(custom_message)
		result_builder.WriteByte('"')
		result_builder.WriteString(", ")
	} else {
		result_builder.WriteString(custom_message)
		result_builder.WriteString(", ")
	}
	messages := MessagesToJson(payload[4:])
	result_builder.WriteString("\"messages\": ")
	result_builder.WriteString(messages)
	result_builder.WriteByte('}')
	//println("Result:",result_builder.String())
	return result_builder.String()
}
