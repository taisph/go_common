package log

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type CheckedFormat struct {
	s string
}

func (f *CheckedFormat) String() string {
	return f.s
}

func (f *CheckedFormat) Set(s string) error {
	switch s {
	case "console", "json":
		f.s = s
	default:
		return errors.Errorf("unrecognized log format %q", s)
	}
	return nil
}

type CheckedLevel struct {
	l zerolog.Level
}

func (l *CheckedLevel) String() string {
	return l.String()
}

func (l *CheckedLevel) Set(s string) error {
	var err error
	l.l, err = zerolog.ParseLevel(s)
	return err
}
