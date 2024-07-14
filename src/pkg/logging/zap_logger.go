package logging

import (
	"github.com/MrRezoo/CarApp/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapSingletonLogger *zap.SugaredLogger

var logLevelMapping = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

type zapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func newZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (logger *zapLogger) getLogLevel() zapcore.Level {
	level, exists := logLevelMapping[logger.cfg.Logger.Level]
	if !exists {
		return zapcore.DebugLevel
	}
	return level
}

func (logger *zapLogger) Init() {
	once.Do(func() {
		writeSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logger.cfg.Logger.FilePath,
			MaxSize:    1,
			MaxAge:     5,
			LocalTime:  true,
			MaxBackups: 10,
			Compress:   true,
		})
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			writeSyncer,
			logger.getLogLevel(),
		)
		zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
		zapLogger = zapLogger.With("AppName", "CarApp", "LoggerName", "ZapLogger")
		zapSingletonLogger = zapLogger
	})
	logger.logger = zapSingletonLogger

}

func (logger *zapLogger) Info(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(category, subCategory, extra)
	logger.logger.Infow(message, params...)
}

func (logger *zapLogger) InfoF(template string, args ...interface{}) {
	logger.logger.Infof(template, args)
}

func (logger *zapLogger) Debug(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(category, subCategory, extra)
	logger.logger.Debugw(message, params...)
}

func (logger *zapLogger) DebugF(template string, args ...interface{}) {
	logger.logger.Debugf(template, args)
}

func (logger *zapLogger) Warn(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(category, subCategory, extra)
	logger.logger.Warnw(message, params...)
}

func (logger *zapLogger) WarnF(template string, args ...interface{}) {
	logger.logger.Warnf(template, args)
}

func (logger *zapLogger) Error(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(category, subCategory, extra)
	logger.logger.Errorw(message, params...)
}

func (logger *zapLogger) ErrorF(template string, args ...interface{}) {
	logger.logger.Errorf(template, args)
}

func (logger *zapLogger) Fatal(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(category, subCategory, extra)
	logger.logger.Fatalw(message, params...)
}

func (logger *zapLogger) FatalF(template string, args ...interface{}) {
	logger.logger.Fatalf(template, args)
}

func prepareLogKeys(category Category, subCategory SubCategory, extra map[ExtraKey]interface{}) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 2)
	}
	extra["Category"] = category
	extra["SubCategory"] = subCategory
	params := logParamsToZapParams(extra)
	return params
}
