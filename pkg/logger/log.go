package logger

import (
	"log"
	"os"
	"sync"
)

type Logger struct {
	errLog  *log.Logger
	infoLog *log.Logger
}

var lg *Logger
var once sync.Once

func (l *Logger) Error(format string, v ...any) {
	l.errLog.Printf(format, v...)
}

func (l *Logger) Info(format string, v ...any) {
	l.infoLog.Printf(format, v...)
}

func GetLogger() *Logger {
	once.Do(func() {
		if lg == nil {
			lg = setupLogs()
		}
	})
	return lg
}

func setupLogs() *Logger {
	var (
		infoLog  *log.Logger
		errorLog *log.Logger
	)

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	lg := &Logger{errLog: errorLog, infoLog: infoLog}

	return lg
}
