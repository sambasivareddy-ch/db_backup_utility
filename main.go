package main

import (
	"fmt"

	"github.com/sambasivareddy-ch/db_backup_utility/cmd"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/logging"

	"github.com/joho/godotenv"
)

// Load environment variables from .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}
}

// Entry Point
func main() {
	// Initialize the logger
	// This will create a log file in the current directory
	err := logging.InitializeLogger()
	if err != nil {
		fmt.Println("Error occurred intializing the logfile")
	}

	// Initialize the command line interface
	cmd.Execute()
}
