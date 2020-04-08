package log

import (
	"runtime"

	"github.com/rs/zerolog"
)

type Caller struct {
	skip int
}

func (c Caller) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	_, file, line, ok := runtime.Caller(c.skip)
	if !ok {
		return
	}
	e.Dict("logging.googleapis.com/sourceLocation", zerolog.Dict().Str("file", file).Int("line", line))
}
