package log

import (
	"github.com/fanke15/g2a-admin/pkg/basic"
	"os"

	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	DefaultPath   = "./logfiles/"
	DefaultSuffix = ".log"
	DefaultSep    = ""
)

var (
	// zap操作对象
	zapLog *zap.Logger
)

// 项目初始化
func New(kvs ...interface{}) {
	var (
		// 文件存储设置 输出hook
		hook = lumberjack.Logger{
			Filename:   basic.AnySliceToStr(DefaultSep, DefaultPath, time.Now().Format(basic.DefaultDateFormat), DefaultSuffix),
			MaxSize:    1,
			MaxBackups: 5,
			MaxAge:     30,
			Compress:   false,
		}
		// 编码器配置,zap日志初始化设置
		logEncodeConfig = zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "source",
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.LowercaseLevelEncoder, // 小写编码器
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(basic.DefaultTimeFormat))
			},
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder, // 半路径编码器
		}
		// 日志输出级别设置
		tree = []zapcore.Core{
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(logEncodeConfig),
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 输出到控制台os.Stdout
				zapcore.InfoLevel,
			)}
	)
	tree = append(tree, zapcore.NewCore( // 输出到文件hook
		zapcore.NewJSONEncoder(logEncodeConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)),
		zapcore.ErrorLevel,
	))
	core := zapcore.NewTee(tree...)
	zapLog = zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(customKV(kvs...)...)) // 构造日志
}

func Info(errMsg string, others ...interface{}) {
	filed := customKV(others...)
	zapLog.Info(errMsg, filed...)
}

func Error(errMsg string, others ...interface{}) {
	filed := customKV(others...)
	zapLog.Error(errMsg, filed...)
}

//---------------------------内部私有方法---------------------------//

// 自定义默认k-v值
func customKV(kvs ...interface{}) []zap.Field {
	var (
		num = len(kvs) / 2
		fs  []zap.Field
	)
	if num > 0 {
		for i := 1; i <= num; i++ {
			f := zap.Any(basic.AnyToStr(kvs[2*i-2]), kvs[2*i-1])
			fs = append(fs, f)
		}
	}
	return fs
}
