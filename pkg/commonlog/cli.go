package commonlog

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/rs/zerolog"
)

type LogCliConfig struct {
	Format CheckedFormat `name:"format" help:"Output format of log messages. One of: [console, json, json-sd]" default:"console" env:"LOG_FORMAT"`
	Level  CheckedLevel  `name:"level" help:"Only log messages with the given severity or above. One of: [trace, debug, info, warn, error, fatal, panic]" default:"info" env:"LOG_LEVEL"`
}

type LogCli struct {
	LogCliConfig `embed prefix:"log."`
}

type CheckedFormat struct {
	kong.MapperValue

	s string
}

func (f *CheckedFormat) Decode(ctx *kong.DecodeContext) error {
	var format string
	err := ctx.Scan.PopValueInto("format", &format)
	if err != nil {
		return err
	}

	return f.Set(format)
}

func (f *CheckedFormat) String() string {
	return f.s
}

func (f *CheckedFormat) Set(s string) error {
	switch s {
	case "console", "json", "json-sd":
		f.s = s
	default:
		return fmt.Errorf("unrecognized log format: %q", s)
	}

	return nil
}

type CheckedLevel struct {
	kong.MapperValue

	l zerolog.Level
}

func (l *CheckedLevel) Decode(ctx *kong.DecodeContext) error {
	var level string
	err := ctx.Scan.PopValueInto("level", &level)
	if err != nil {
		return err
	}

	return l.Set(level)
}

func (l *CheckedLevel) String() string {
	return l.String()
}

func (l *CheckedLevel) Set(s string) error {
	var err error
	l.l, err = zerolog.ParseLevel(s)

	return err
}
