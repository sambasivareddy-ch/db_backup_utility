package logging

import (
	"errors"
	"log"
	"os"
)

type LoggerStruct struct {
	logger *log.Logger
	file   *os.File
}

var Logger *LoggerStruct

func InitializeLogger() error {
	// Create a log file
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return errors.New("log file creation failed")
	}

	// Now map logfile to logger
	// And specifies the format of the each entry in the log file
	// Here log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile tells each line will appears like
	// 2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	logger := log.New(logFile, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	Logger = &LoggerStruct{
		logger: logger,
		file:   logFile,
	}

	return nil
}

// Helper Functions
func (l *LoggerStruct) LogInfo(format string, v ...interface{}) {
	l.logger.Printf("INFO: "+format, v...)
}

func (l *LoggerStruct) LogError(format string, v ...interface{}) {
	l.logger.Printf("ERROR: "+format, v...)
}

func (l *LoggerStruct) LogCommand(command string) {
	l.logger.Printf("INFO: Running Command = %s", command)
}

func (l *LoggerStruct) CloseLogger() error {
	return l.file.Close()
}
