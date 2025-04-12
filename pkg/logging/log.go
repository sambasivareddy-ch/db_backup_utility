package logging

import (
	"errors"
	"log"
	"os"
)

// LoggerStruct is a structure that holds the logger and the log file.
// It is used to create a logger that writes to a file.
type LoggerStruct struct {
	logger *log.Logger
	file   *os.File
}

// Logger is a global variable that holds the logger instance.
// It is used to log messages to a file.
var Logger *LoggerStruct

/*
InitializeLogger function is used to create a logger that writes to a file.
It creates a log file named app.log in the current directory.
It sets the log format to include the date, time, and file name.
It returns an error if the log file creation fails.
*/
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

// LogInfo function is used to log an info message.
func (l *LoggerStruct) LogInfo(format string, v ...interface{}) {
	l.logger.Printf("INFO: "+format, v...)
}

// LogError function is used to log an error message.
func (l *LoggerStruct) LogError(format string, v ...interface{}) {
	l.logger.Printf("ERROR: "+format, v...)
}

// LogCommand function is used to log a command message.
func (l *LoggerStruct) LogCommand(command string) {
	l.logger.Printf("INFO: Running Command = %s", command)
}

// CloseLogger function is used to close the log file.
// It is called when the application is shutting down.
func (l *LoggerStruct) CloseLogger() error {
	return l.file.Close()
}
