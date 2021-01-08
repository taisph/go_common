package commonlog

import (
	"runtime"
	"strings"

	"github.com/rs/zerolog"
)

var (
	BasePath string
)

type Caller struct {
	skip int
}

func (c Caller) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	_, file, line, ok := runtime.Caller(c.skip)
	if !ok {
		return
	}
	e.Dict("logging.googleapis.com/sourceLocation", zerolog.Dict().Str("file", strings.TrimPrefix(file, BasePath)).Int("line", line))
}
