package logging

import (
	"github.com/MrRezoo/CarApp/config"
	"github.com/rs/zerolog"
	"os"
	"sync"
)

var once sync.Once
var zeroSingletonLogger *zerolog.Logger

type zeroLogger struct {
	cfg    *config.Config
	logger *zerolog.Logger
}

var zeroLogLevelMapping = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

func newZeroLogger(cfg *config.Config) *zeroLogger {
	logger := &zeroLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *zeroLogger) getLogLevel() zerolog.Level {
	level, exists := zeroLogLevelMapping[l.cfg.Logger.Level]
	if !exists {
		return zerolog.DebugLevel
	}
	return level
}

func (l *zeroLogger) Init() {
	once.Do(func() {
		file, err := os.OpenFile(l.cfg.Logger.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic("could not open log file")
		}

		var logger = zerolog.New(file).
			With().
			Timestamp().
			Str("AppName", "CarApp").
			Str("LoggerName", "ZeroLog").
			Logger()
		zerolog.SetGlobalLevel(l.getLogLevel())
		zeroSingletonLogger = &logger
	})
	l.logger = zeroSingletonLogger

}

func (l *zeroLogger) Info(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	l.logger.
		Info().
		Str("Category", string(category)).
		Str("SubCategory", string(subCategory)).
		Fields(logParamsToZeroParams(extra)).
		Msg(message)
}

func (l *zeroLogger) InfoF(template string, args ...interface{}) {
	l.logger.Info().Msgf(template, args...)
}

func (l *zeroLogger) Debug(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	l.logger.
		Debug().
		Str("Category", string(category)).
		Str("SubCategory", string(subCategory)).
		Fields(logParamsToZeroParams(extra)).
		Msg(message)
}

func (l *zeroLogger) DebugF(template string, args ...interface{}) {
	l.logger.Debug().Msgf(template, args...)
}

func (l *zeroLogger) Warn(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	l.logger.
		Warn().
		Str("Category", string(category)).
		Str("SubCategory", string(subCategory)).
		Fields(logParamsToZeroParams(extra)).
		Msg(message)
}

func (l *zeroLogger) WarnF(template string, args ...interface{}) {
	l.logger.Warn().Msgf(template, args...)
}

func (l *zeroLogger) Error(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	l.logger.
		Error().
		Str("Category", string(category)).
		Str("SubCategory", string(subCategory)).
		Fields(logParamsToZeroParams(extra)).
		Msg(message)
}

func (l *zeroLogger) ErrorF(template string, args ...interface{}) {
	l.logger.Error().Msgf(template, args...)
}

func (l *zeroLogger) Fatal(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	l.logger.
		Fatal().
		Str("Category", string(category)).
		Str("SubCategory", string(subCategory)).
		Fields(logParamsToZeroParams(extra)).
		Msg(message)
}

func (l *zeroLogger) FatalF(template string, args ...interface{}) {
	l.logger.Fatal().Msgf(template, args...)
}
