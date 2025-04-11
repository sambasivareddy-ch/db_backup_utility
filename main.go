package main

import (
	"fmt"

	"github.com/sambasivareddy-ch/db_backup_utility/cmd"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/logging"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {
	err := logging.InitializeLogger()
	if err != nil {
		fmt.Println("Error occurred intializing the logfile")
	}

	cmd.Execute()
}
