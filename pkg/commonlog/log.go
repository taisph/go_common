package commonlog

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New(cfg *LogCliConfig) zerolog.Logger {
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "time"
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.SetGlobalLevel(cfg.Level.l)

	var l zerolog.Logger
	switch cfg.Format.String() {
	case "console":
		l = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.StampNano})
	case "json":
		l = zerolog.New(os.Stdout)
	}

	return l.Hook(Caller{3}).With().Timestamp().Logger().Level(cfg.Level.l)
}
