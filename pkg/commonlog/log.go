package commonlog

import (
	"os"
	"time"

	"github.com/rs/zerolog"

	"github.com/taisph/go_common/pkg/commonlog/stackdriver"
)

func New(cfg *LogCliConfig) zerolog.Logger {
	zerolog.SetGlobalLevel(cfg.Level.l)

	var l zerolog.Logger
	switch cfg.Format.String() {
	case "console":
		l = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Stamp}).With().Caller().Timestamp().Logger()
	case "json":
		l = zerolog.New(os.Stdout).With().Caller().Timestamp().Logger()
	case "json-sd":
		l = stackdriver.NewJsonStackDriverLog(os.Stdout)
	}

	return l.Level(cfg.Level.l)
}
