package logger

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type AggregatedLogger struct {
  infoLogger *log.Logger
  warnLogger *log.Logger
  errorLogger *log.Logger
  LogFile *os.File
}

func (l *AggregatedLogger) Info(v ...interface{}) {
  l.infoLogger.Println(v...)
}

func (l *AggregatedLogger) Warn(v ...interface{}) {
  l.warnLogger.Println(v...)
}

func (l *AggregatedLogger) Error(v ...interface{}) {
  l.errorLogger.Println(v...)
}

func NewAggregatedLogger() AggregatedLogger {
  logFilePath := "./app.log"
  logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
    log.Panicln("Error opening log file:", err)
  }

  flags := log.LstdFlags
  infoLogger := log.New(logFile, "INFO: ", flags)
  warnLogger := log.New(logFile, "WARN: ", flags)
  errorLogger := log.New(logFile, "ERROR: ", flags)

  return AggregatedLogger{
    infoLogger: infoLogger,
    warnLogger: warnLogger,
    errorLogger: errorLogger,
    LogFile: logFile,
  }
}

func EndpointLog(fn string, r *http.Request) string {
  return fmt.Sprintf("Endpoint hit: %v, method: %v, path: %v", fn, r.Method, r.URL.Path)
}
