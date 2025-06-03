package logx

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// Init initializes the global logger with either "dev" or "prod" mode
func Init(mode string) {
	once.Do(func() {
		var err error
		var cfg zap.Config

		switch mode {
		case "prod":
			cfg = zap.NewProductionConfig()
			cfg.EncoderConfig.TimeKey = "ts"
			cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		default:
			cfg = zap.NewDevelopmentConfig()

			// Customize output format
			cfg.EncoderConfig.TimeKey = "time"
			cfg.EncoderConfig.LevelKey = "level"
			cfg.EncoderConfig.NameKey = "logger"
			cfg.EncoderConfig.CallerKey = "caller"
			cfg.EncoderConfig.MessageKey = "msg"
			cfg.EncoderConfig.StacktraceKey = "" // Disable stack trace

			// Color and format
			cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05"))
			}
		}

		// Enable file:line caller
		logger, err = cfg.Build(zap.AddCaller(), zap.AddCallerSkip(1))
		if err != nil {
			panic("failed to initialize logger: " + err.Error())
		}
	})
}

// L returns the base zap logger
func L() *zap.Logger {
	if logger == nil {
		Init("dev")
	}
	return logger
}

// S returns the sugared logger
func S() *zap.SugaredLogger {
	return L().Sugar()
}

// Shortcut log functions (structured)
func Debug(msg string, fields ...zap.Field) { L().Debug(msg, fields...) }
func Info(msg string, fields ...zap.Field)  { L().Info(msg, fields...) }
func Warn(msg string, fields ...zap.Field)  { L().Warn(msg, fields...) }
func Error(msg string, fields ...zap.Field) { L().Error(msg, fields...) }
func Fatal(msg string, fields ...zap.Field) { L().Fatal(msg, fields...) }

// Sugared log functions (for printf-style)
func Debugf(format string, args ...interface{}) { S().Debugf(format, args...) }
func Infof(format string, args ...interface{})  { S().Infof(format, args...) }
func Warnf(format string, args ...interface{})  { S().Warnf(format, args...) }
func Errorf(format string, args ...interface{}) { S().Errorf(format, args...) }
func Fatalf(format string, args ...interface{}) { S().Fatalf(format, args...) }
