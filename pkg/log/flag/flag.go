package flag

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/taisph/go-common/pkg/log"
)

func AddFlags(app *kingpin.Application, cfg *log.Config) {
	cfg.Format = &log.CheckedFormat{}
	app.Flag("log.format", "Output format of log messages. One of: [console, json]").Default("console").SetValue(cfg.Format)

	cfg.Level = &log.CheckedLevel{}
	app.Flag("log.level", "Only log messages with the given severity or above. One of: [trace, debug, info, warn, error, fatal, panic]").Default("info").SetValue(cfg.Level)
}
