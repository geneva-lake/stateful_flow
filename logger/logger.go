package logger

import (
	"fmt"
	"log"
	"strings"
)

type LogLevel string

const (
	Info  LogLevel = "info"
	Error LogLevel = "error"
	Panic LogLevel = "panic"
)

type LogString struct {
	strings.Builder
}

func (l *LogString) LevelBegin(level LogLevel, function string) *LogString {
	l.WriteString(fmt.Sprintf("%s: function: %s", level, function))
	return l
}

func (l *LogString) ErrorCheck(err error) *LogString {
	if err != nil {
		l.WriteString(fmt.Sprintf(", error: %s", err))
	}
	return l
}

func (l *LogString) PanicCheck(panic interface{}) *LogString {
	if panic != nil {
		l.WriteString(fmt.Sprintf(", panic: %s", panic))
	}
	return l
}

func (l *LogString) AdditionalWrite(additional map[string]interface{}) *LogString {
	for k, v := range additional {
		l.WriteString(fmt.Sprintf(", %s: %v", k, v))
	}
	return l
}

func (l *LogString) String() string {
	return l.Builder.String()
}

func Log(level LogLevel, function string, err error, panic interface{}, additional map[string]interface{}) {
	logString := new(LogString)
	s := logString.LevelBegin(level, function).PanicCheck(panic).ErrorCheck(err).String()
	log.Println(s)
}
