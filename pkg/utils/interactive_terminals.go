package utils

import (
	"os"

	"github.com/sambasivareddy-ch/db_backup_utility/context"

	"github.com/AlecAivazis/survey/v2"
)

/*
	RunConfigInteractiveInputTerminal
	This function is used to get the configuration details from the user.
	It takes the database connection parameters from the user, backup dir and saves them to the context.
*/
func RunConfigInteractiveInputTerminal(config *context.DBSessionContext) {
	survey.AskOne(&survey.Select{
		Message: "Select database type:",
		Options: []string{"postgres", "sql"},
		Default: "postgres",
	}, &config.DBType)

	survey.AskOne(&survey.Input{
		Message: "Enter database host:",
		Default: "localhost",
	}, &config.DBHost)

	defaultPort := GetDefaultPort(config.DBType)
	survey.AskOne(&survey.Input{
		Message: "Enter database port:",
		Default: defaultPort,
	}, &config.DBPort)

	survey.AskOne(&survey.Input{
		Message: "Enter username:",
	}, &config.DBUsername, survey.WithValidator(survey.Required))

	survey.AskOne(&survey.Password{
		Message: "Enter password:",
	}, &config.DBPassword)

	survey.AskOne(&survey.Input{
		Message: "Enter database name:",
	}, &config.DBName, survey.WithValidator(survey.Required))

	survey.AskOne(&survey.Input{
		Message: "Enter backup directory path:",
	}, &config.BackupDir, survey.WithValidator(IsDirExists))
}

/*
	RunRestoreInteractiveInputTerminal
	This function is used to get the restore file from the user.
	It takes the backup directory path from the user and lists all the files in the directory.
	It gives the user a chance to select one of them.
*/
func RunRestoreInteractiveInputTerminal(restoreParams *context.DBSessionContext) {
	// Now get all backup file in provided Backup folder and
	// Give the user chance to select one of them.
	var backup_files = make([]string, 0)
	files, _ := os.ReadDir(restoreParams.BackupDir)
	for _, file := range files {
		backup_files = append(backup_files, file.Name())
	}

	if len(files) > 0 {
		survey.AskOne(&survey.Select{
			Message: "Select Backup File",
			Options: backup_files,
			Default: backup_files[0],
		}, &restoreParams.RestoreFile, survey.WithValidator(survey.Required))
	}
}

/*
	GetDefaultPort
	This function returns the default port for the given database type.
	It takes the database type as input and returns the default port as a string.
*/
func GetDefaultPort(dbType string) string {
	if dbType == "postgres" {
		return "5432"
	} else if dbType == "sql" {
		return "3306"
	}
	return "27017"
}
