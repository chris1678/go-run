package logger

import (
	"github.com/chris1678/go-run/config"
	"github.com/chris1678/go-run/utils"
	"github.com/chris1678/go-run/utils/writer"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
)

var LogHelper *zap.SugaredLogger

// Setup 设置logger
func Initialize() {
	c := config.LoggerConfig
	SetupLogger(
		WithPath(c.Path),
		WithLevel(c.Level),
		WithStdout(c.Stdout),
		WithCap(c.Cap),
	)
}

// SetupLogger 日志 cap 单位为kb
func SetupLogger(opts ...Option) {
	op := setDefault()
	for _, o := range opts {
		o(&op)
	}
	if !utils.PathExist(op.path) {
		err := utils.PathCreate(op.path)
		if err != nil {
			log.Fatalf("create dir error: %s", err.Error())
		}
	}
	var err error
	var output io.Writer
	switch op.stdout {
	case "file":
		output, err = writer.NewFileWriter(
			writer.WithPath(op.path),
			writer.WithCap(op.cap<<10),
		)
		if err != nil {
			log.Fatalf("logger setup error: %s", err.Error())
		}
	default:
		output = os.Stdout
	}
	level := GetLevel(op.level)
	if err != nil {
		log.Fatalf("get logger level error, %s", err.Error())
	}
	encoder := getEncoder()
	writeSyncer := zapcore.AddSync(output)
	core := zapcore.NewCore(encoder, writeSyncer, level)
	l := zap.New(core, zap.AddCaller())
	LogHelper = l.Sugar()

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 修改时间编码器

	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder 打印更符合人们观察的方式
	return zapcore.NewConsoleEncoder(encoderConfig)
}
