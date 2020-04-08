package echo

import (
	"io"

	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

var (
	gomZeroMap = map[log.Lvl]zerolog.Level{
		log.DEBUG: zerolog.DebugLevel,
		log.INFO:  zerolog.InfoLevel,
		log.WARN:  zerolog.WarnLevel,
		log.ERROR: zerolog.ErrorLevel,
		log.OFF:   zerolog.NoLevel,
	}
	zeroGomMap = map[zerolog.Level]log.Lvl{
		zerolog.DebugLevel: log.DEBUG,
		zerolog.InfoLevel:  log.INFO,
		zerolog.WarnLevel:  log.WARN,
		zerolog.ErrorLevel: log.ERROR,
		zerolog.NoLevel:    log.OFF,
	}
)

type Logger struct {
	zl zerolog.Logger
}

func New(zerolog zerolog.Logger) *Logger {
	return &Logger{zl: zerolog}
}

func (l *Logger) Output() io.Writer {
	return l.zl
}

func (l *Logger) SetOutput(w io.Writer) {
	panic("implement me")
}

func (l *Logger) Prefix() string {
	panic("implement me")
}

func (l *Logger) SetPrefix(p string) {
	panic("implement me")
}

func (l *Logger) Level() log.Lvl {
	return zeroGomMap[l.zl.GetLevel()]
}

func (l *Logger) SetLevel(v log.Lvl) {
	l.zl = l.zl.Level(gomZeroMap[v])
}

func (l *Logger) SetHeader(h string) {
	panic("implement me")
}

func (l *Logger) Print(i ...interface{}) {
	panic("implement me")
}

func (l *Logger) Printf(format string, args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Printj(j log.JSON) {
	panic("implement me")
}

func (l *Logger) Debug(i ...interface{}) {
	panic("implement me")
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Debugj(j log.JSON) {
	panic("implement me")
}

func (l *Logger) Info(i ...interface{}) {
	panic("implement me")
}

func (l *Logger) Infof(format string, args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Infoj(j log.JSON) {
	panic("implement me")
}

func (l *Logger) Warn(i ...interface{}) {
	panic("implement me")
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Warnj(j log.JSON) {
	panic("implement me")
}

func (l *Logger) Error(i ...interface{}) {
	panic("implement me")
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Errorj(j log.JSON) {
	panic("implement me")
}

func (l *Logger) Fatal(i ...interface{}) {
	panic("implement me")
}

func (l *Logger) Fatalj(j log.JSON) {
	panic("implement me")
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	panic("implement me")
}

func (l *Logger) Panic(i ...interface{}) {
	panic("implement me")
}

func (l *Logger) Panicj(j log.JSON) {
	panic("implement me")
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	panic("implement me")
}
