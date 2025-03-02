package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/ipfs/go-log"
)

var _ log.StandardLogger = &Logger{}

type Logger struct {
	level log.LogLevel
	zap   *zap.SugaredLogger
}

func New(lvl log.LogLevel) *Logger {
	cfg := zap.Config{
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Level:            zap.NewAtomicLevelAt(zapcore.Level(lvl)),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	sugar := logger.Sugar()
	return &Logger{level: lvl, zap: sugar}
}

func joinMsg(args ...interface{}) (message string) {
	for i, m := range args {
		if i > 0 {
			message += " "
		}
		message += fmt.Sprintf("%v", m)
	}
	return
}

func (l *Logger) Debug(args ...interface{}) {
	l.zap.Debug(joinMsg(args...))
}

func (l *Logger) Debugf(f string, args ...interface{}) {
	l.zap.Debugf(f, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.zap.Error(joinMsg(args...))
}

func (l *Logger) Errorf(f string, args ...interface{}) {
	l.zap.Errorf(f, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.zap.Fatal(joinMsg(args...))
}

func (l *Logger) Fatalf(f string, args ...interface{}) {
	l.zap.Fatalf(f, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.zap.Info(joinMsg(args...))
}

func (l *Logger) Infof(f string, args ...interface{}) {
	l.zap.Infof(f, args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.Fatal(args...)
}

func (l *Logger) Panicf(f string, args ...interface{}) {
	l.Fatalf(f, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.zap.Warn(joinMsg(args...))
}

func (l *Logger) Warnf(f string, args ...interface{}) {
	l.zap.Warnf(f, args...)
}

func (l *Logger) Warning(args ...interface{}) {
	l.Warn(args...)
}

func (l *Logger) Warningf(f string, args ...interface{}) {
	l.Warnf(f, args...)
}
