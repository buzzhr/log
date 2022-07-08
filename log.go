package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	WarningLogger = log.New(os.Stderr, "WARNING: ", log.LstdFlags)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.LstdFlags)
}

func meta(skip int) string {
	var ok bool
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "???"
		line = 0
	}
	short := file
	p := strings.Split(file, "/")
	if len(p) > 1 {
		short = p[len(p)-1]
	}
	return fmt.Sprintf(" %s:%d", short, line)
}
func Info(v ...interface{}) {
	v = append(v, meta(2))
	InfoLogger.Println(v...)
}
func Infof(format string, v ...interface{}) {
	InfoLogger.Println(fmt.Sprintf(format, v...), meta(2))
}
func Warn(v ...interface{}) {
	v = append(v, meta(2))
	WarningLogger.Println(v...)
}
func Warnf(format string, v ...interface{}) {
	WarningLogger.Printf(fmt.Sprintf(format, v...), meta(2))
}
func Error(v ...interface{}) {
	v = append(v, meta(2))
	ErrorLogger.Print(v...)
}
func PrintError(err error) {
	if err != nil {
		ErrorLogger.Println(err, meta(2))
	}
}
func PrintErrorMsg(msg string, err error) {
	if err != nil {
		ErrorLogger.Println(msg, " ", err.Error(), meta(2))
	}
}
func Errorf(format string, v ...interface{}) {
	ErrorLogger.Printf(fmt.Sprintf(format, v...), meta(2))
}
func Fatal(v ...interface{}) {
	v = append(v, meta(2))
	ErrorLogger.Fatalln(v...)
}
func Fatalf(format string, v ...interface{}) {
	ErrorLogger.Fatalf(fmt.Sprintf(format, v...), meta(2))
}
