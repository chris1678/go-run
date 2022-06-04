package logger

import (
	"go.uber.org/zap/zapcore"
)

var levels map[string]zapcore.Level

func init() {
	levels = make(map[string]zapcore.Level, 0)
	levels["debug"] = zapcore.DebugLevel
	levels["info"] = zapcore.InfoLevel
	levels["warn"] = zapcore.WarnLevel
	levels["error"] = zapcore.ErrorLevel
	levels["dpanic"] = zapcore.DPanicLevel
	levels["panic"] = zapcore.PanicLevel
	levels["fatal"] = zapcore.FatalLevel

}
func GetLevel(levelStr string) zapcore.Level {
	l, e := levels[levelStr]
	if e {
		return l
	} else {
		return zapcore.DebugLevel
	}
}
