package stackdriver

import (
	"io"
	"time"

	"github.com/rs/zerolog"
)

func NewJsonStackDriverLog(w io.Writer) zerolog.Logger {
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "time"
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		return levelToServerity(l)
	}

	return zerolog.New(w).Hook(Caller{3}).With().Timestamp().Logger()
}

func levelToServerity(l zerolog.Level) string {
	switch l {
	case zerolog.TraceLevel:
		return "DEBUG"
	case zerolog.DebugLevel:
		return "DEBUG"
	case zerolog.InfoLevel:
		return "INFO"
	case zerolog.WarnLevel:
		return "WARNING"
	case zerolog.ErrorLevel:
		return "ERROR"
	case zerolog.FatalLevel:
		return "CRITICAL"
	case zerolog.PanicLevel:
		return "EMERGENCY"
	case zerolog.NoLevel:
		return "DEFAULT"
	}
	return "DEFAULT"
}
