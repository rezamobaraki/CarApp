package logging

import "github.com/MrRezoo/CarApp/config"

type Logger interface {
	Init()

	Info(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	InfoF(template string, args ...interface{})

	Debug(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	DebugF(template string, args ...interface{})

	Warn(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	WarnF(template string, args ...interface{})

	Error(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	ErrorF(template string, args ...interface{})

	Fatal(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	FatalF(template string, args ...interface{})
}

func NewLogger(cfg *config.Config) Logger {
	if cfg.Logger.Type == "zap" {
		return newZapLogger(cfg)
	} else if cfg.Logger.Type == "zero" {
		return newZeroLogger(cfg)
	}
	panic("Logger not supported")
}

// file <- filebeat -> elastic search -> kibana
