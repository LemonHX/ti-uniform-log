package test

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Log struct {
	file *os.File
}

func (log *Log) CloseLog() {
	log.file.Close()
}

func (log *Log) GenerateLogs(estimated_max_file_size int) error {
	if log.file == nil {
		if file, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE, 0755); err != nil {
			return err
		} else {
			log.file = file
		}
	}
while:
	info, err := log.file.Stat()
	if err != nil {
		return err
	} else {
		if info.Size() < int64(estimated_max_file_size) {
			str := GenerateLog(rand.Intn(30))
			if _, err = log.file.WriteString(str + "\n"); err != nil {
				return err
			} else {
				goto while
			}
		} else {
			str := GenerateLog(rand.Intn(30))
			if _, err = log.file.WriteString(str); err != nil {
				return err
			} else {
				return nil
			}
		}
	}
}

func GenerateLog(message_size int) string {
	date := generateDateTime()
	level := generateLogLevel()
	pos := generatePosition()
	custom_message := generateCustomMessage()
	var messageBuilder strings.Builder
	for i := 0; i < message_size; i += 1 {
		messageBuilder.WriteString(generateKVPair())
	}
	messages := messageBuilder.String()
	var resBuilder strings.Builder
	resBuilder.WriteString(date)
	resBuilder.WriteString(" ")
	resBuilder.WriteString(level)
	resBuilder.WriteString(" ")
	resBuilder.WriteString(pos)
	resBuilder.WriteString(" ")
	resBuilder.WriteString(custom_message)
	if len(messages) > 0 {
		resBuilder.WriteString(" ")
		resBuilder.WriteString(messages)
	}
	return resBuilder.String()
}

func generateDateTime() string {
	return "[2021/07/19 11:45:14.191 +08:00]"
}

func generateLogLevel() string {
	level := []string{"[INFO]", "[DEBUG]", "[WARN]", "[ERROR]"}
	return level[rand.Intn(4)]
}

func generatePosition() string {
	if rand.Intn(2) == 0 {
		file := []string{"a.rs", "b.go", "c.cpp", "d.py"}
		pos := rand.Intn(114514)
		var positionBuilder strings.Builder
		positionBuilder.WriteString("[")
		positionBuilder.WriteString(file[rand.Intn(4)])
		positionBuilder.WriteString(":")
		positionBuilder.WriteString(strconv.Itoa(pos))
		positionBuilder.WriteString("]")
		return positionBuilder.String()
	} else {
		return "[<unknown>]"
	}
}

func generateCustomMessage() string {
	message := []string{"[\"Slow query\"]", "[Fast_Query]", "[\"114514\"]", "[\"hhhaaaaaaa\"]"}
	return message[rand.Intn(4)]
}

func generateKVPair() string {
	var builder strings.Builder
	builder.WriteString("[")
	builder.WriteString(generateKey())
	builder.WriteString("=")
	builder.WriteString(generateValue())
	builder.WriteString("]")
	return builder.String()
}

func generateKey() string {
	keys := []string{"sql", "duration", "client", "txn"}
	return keys[rand.Intn(4)]
}

func generateValue() string {
	values := []string{"\"SELECT * FROM TABLE \\n WHERE ID=\\\"abc\\\"\"", "1.14514s", "114.51.4.19:19810", "1145141919810"}
	return values[rand.Intn(4)]
}
