package log

import (
	"io"
	"sync"
)

type Config struct {
	Enabled *bool `json:"enabled"`
	SubLoggerConfig
	LoggerFileConfig *loggerFileConfig `json:"fileSettings,omitempty"`
	AdvancedSettings advancedSettings  `json:"advancedSettings"`
	SubLoggers       []SubLoggerConfig `json:"subloggers,omitempty"`
}
type advancedSettings struct {
	ShowLogSystemName *bool   `json:"showLogSystemName"`
	Spacer            string  `json:"spacer"`
	TimeStampFormat   string  `json:"timeStampFormat"`
	Headers           headers `json:"headers"`
}
type headers struct {
	Info  string `json:"info"`
	Warn  string `json:"warn"`
	Debug string `json:"debug"`
	Error string `json:"error"`
}
type SubLoggerConfig struct {
	Name   string `json:"name,omitempty"`
	Level  string `json:"level"`
	Output string `json:"output"`
}
type loggerFileConfig struct {
	FileName string `json:"filename,omitempty"`
	Rotate   *bool  `json:"rotate,omitempty"`
	MaxSize  int64  `json:"maxsize,omitempty"`
}
type Logger struct {
	ShowLogSystemName                                bool
	Timestamp                                        string
	InfoHeader, ErrorHeader, DebugHeader, WarnHeader string
	Spacer                                           string
}
type Levels struct {
	Info, Debug, Warn, Error bool
}
type subLogger struct {
	name string
	Levels
	output io.Writer
}
type Event struct {
	data   []byte
	output io.Writer
}
type multiWriter struct {
	writers []io.Writer
	mu      sync.RWMutex
}
