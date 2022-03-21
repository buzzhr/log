package defaultlog

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	WarningLogger = log.New(os.Stderr, "WARNING: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)
}

func Info(v ...interface{}) {
	InfoLogger.Print(v...)
}
func Infof(format string, v ...interface{}) {
	InfoLogger.Printf(format, v...)
}
func Warn(v ...interface{}) {
	WarningLogger.Print(v...)
}
func Warnf(format string, v ...interface{}) {
	WarningLogger.Printf(format, v...)
}
func Error(v ...interface{}) {
	ErrorLogger.Print(v...)
}
func Errorf(format string, v ...interface{}) {
	ErrorLogger.Printf(format, v...)
}
func Fatal(v ...interface{}) {
	ErrorLogger.Fatal(v...)
}
func Fatalf(format string, v ...interface{}) {
	ErrorLogger.Fatalf(format, v...)
}
