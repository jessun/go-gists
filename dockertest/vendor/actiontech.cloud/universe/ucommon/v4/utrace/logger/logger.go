package logger

import (
	"actiontech.cloud/universe/ucommon/v4/utrace/logger/otellogrus"

	log "github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

const (
	defaultOutFileName = "./logs/tracing.log"
	defaultLogMaxSizeM = 500
	defaultLogLevel    = "info"
)

var (
	traceLogger   *log.Logger
	isInitialized bool
)

func Logger() *log.Logger {
	if !isInitialized {
		panic("please call InitLogger() initialization before use")
	}
	return traceLogger
}

func InitLogger(opts ...LoggerOption) {
	isInitialized = true

	cfg := LoggerConfig{}

	for _, op := range opts {
		cfg = op.apply(cfg)
	}

	if cfg.OutputFileName == "" {
		cfg.OutputFileName = defaultOutFileName
	}

	if cfg.MaxSize == 0 {
		cfg.MaxSize = defaultLogMaxSizeM
	}

	if cfg.Level == "" {
		cfg.Level = defaultLogLevel
	}

	traceLogger = log.New()

	traceLogger.SetFormatter(&log.JSONFormatter{
		PrettyPrint: cfg.PrettyPrint,
	})
	traceLogger.SetOutput(&lumberjack.Logger{
		Filename:   cfg.OutputFileName,
		MaxSize:    cfg.MaxSize, // megabytes
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge, //days
		Compress:   true,       // disabled by default
	})
	logLevel, err := log.ParseLevel(cfg.Level)
	if err != nil {
		panic(err)
	}
	traceLogger.SetLevel(logLevel)
	traceLogger.SetReportCaller(cfg.ReportCaller)

	//set log events to span
	traceLogger.AddHook(otellogrus.NewHook(otellogrus.WithLevels(
		log.AllLevels...,
	)))
}

type LoggerConfig struct {
	PrettyPrint    bool
	OutputFileName string
	MaxBackups     int
	MaxSize        int
	MaxAge         int
	ReportCaller   bool
	Level          string
}

type LoggerOption interface {
	apply(LoggerConfig) LoggerConfig
}

type loggerOptionFunc func(LoggerConfig) LoggerConfig

func (fn loggerOptionFunc) apply(cfg LoggerConfig) LoggerConfig {
	return fn(cfg)
}

func WithFileName(fileNameString string) LoggerOption {
	return loggerOptionFunc(func(cfg LoggerConfig) LoggerConfig {
		if fileNameString != "" {
			cfg.OutputFileName = fileNameString
		}
		return cfg
	})
}

func WithLoggerLevel(level string) LoggerOption {
	return loggerOptionFunc(func(cfg LoggerConfig) LoggerConfig {
		if level != "" {
			cfg.Level = level
		}
		return cfg
	})
}

func WithMaxBackups(maxBackups int) LoggerOption {
	return loggerOptionFunc(func(cfg LoggerConfig) LoggerConfig {
		cfg.MaxBackups = maxBackups
		return cfg
	})
}

func WithMaxSize(maxSizeM int) LoggerOption {
	return loggerOptionFunc(func(cfg LoggerConfig) LoggerConfig {
		cfg.MaxSize = maxSizeM
		return cfg
	})
}

func WithMaxAge(maxAgeDays int) LoggerOption {
	return loggerOptionFunc(func(cfg LoggerConfig) LoggerConfig {
		cfg.MaxAge = maxAgeDays
		return cfg
	})
}

func WithReportCaller(reportCaller bool) LoggerOption {
	return loggerOptionFunc(func(cfg LoggerConfig) LoggerConfig {
		cfg.ReportCaller = reportCaller
		return cfg
	})
}

func WithPrettyPrint(prettyPrint bool) LoggerOption {
	return loggerOptionFunc(func(cfg LoggerConfig) LoggerConfig {
		cfg.PrettyPrint = prettyPrint
		return cfg
	})
}
