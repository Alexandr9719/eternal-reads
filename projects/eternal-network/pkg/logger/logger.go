package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

var (
	debugLogger *log.Logger = log.NewWithOptions(os.Stdout, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		Level:           log.DebugLevel,
	})
	infoLogger *log.Logger = log.NewWithOptions(os.Stdout, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		Level:           log.InfoLevel,
	})
	warnLogger *log.Logger = log.NewWithOptions(os.Stdout, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		Level:           log.WarnLevel,
	})
	errorLogger *log.Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		Level:           log.ErrorLevel,
	})
	fatalLogger *log.Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		Level:           log.FatalLevel,
	})
)

var (
	Debug  = debugLogger.Debug
	Debugf = debugLogger.Debugf
	Info   = infoLogger.Info
	Infof  = infoLogger.Infof
	Warn   = warnLogger.Warn
	Warnf  = warnLogger.Warnf
	Error  = errorLogger.Error
	Errorf = errorLogger.Errorf
	Fatal  = fatalLogger.Fatal
	Fatalf = fatalLogger.Fatalf
)
