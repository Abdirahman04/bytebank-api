package logger

import (
	"log"
	"os"
)

type AggregatedLogger struct {
  infoLogger *log.Logger
  warnLogger *log.Logger
  errorLogger *log.Logger
}

func (l *AggregatedLogger) info(v ...interface{}) {
  l.infoLogger.Println(v...)
}

func (l *AggregatedLogger) warn(v ...interface{}) {
  l.infoLogger.Println(v...)
}

func (l *AggregatedLogger) error(v ...interface{}) {
  l.infoLogger.Println(v...)
}

func NewAggregatedLogger() AggregatedLogger {
  logFilePath := "../app.log"
  logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
    log.Panicln("Error opening log file:", err)
  }
  defer logFile.Close()

  flags := log.LstdFlags | log.Lshortfile
  infoLogger := log.New(logFile, "INFO: ", flags)
  warnLogger := log.New(logFile, "WARN: ", flags)
  errorLogger := log.New(logFile, "ERROR: ", flags)

  return AggregatedLogger{
    infoLogger: infoLogger,
    warnLogger: warnLogger,
    errorLogger: errorLogger,
  }
}
