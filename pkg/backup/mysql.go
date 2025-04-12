package backup

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/sambasivareddy-ch/db_backup_utility/context"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/executor"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/logging"
)

/*
	Utility function to Backup the MySQL using mysqldump
	This function creates a backup of the MySQL database using the mysqldump command.
	It takes the database connection parameters from the context and creates a backup file in the specified directory.
	The backup file is named with the current timestamp to ensure uniqueness.
	The backup file is created in the format mysql_backup_<timestamp>.sql.
*/
func BackupMySQL(params context.DBSessionContext) {
	current_time := time.Now().Format(time.RFC3339)
	backup_file := fmt.Sprintf("%s/mysql_backup_%s.sql", params.BackupDir, current_time)

	backupCommand := exec.Command("mysqldump",
		"-h", params.DBHost,
		"-u", params.DBUsername,
		"-P", params.DBPort,
		"-p"+params.DBPassword,
		params.DBName,
	)

	// Create the output file
	outFile, err := os.Create(backup_file)
	if err != nil {
		logging.Logger.LogError("Failed to create backup file: %v", err)
		panic(err)
	}
	defer outFile.Close()

	backupCommand.Stdout = outFile // Pipe the file to write dumped data

	executor.ExecuteCommand(backupCommand, "backup", params.DBType, backup_file)
}

/*
	Utility function to Restore the MySQL using mysql
	This function restores the MySQL database from a backup file.
	It takes the database connection parameters from the context and reads the backup file.
*/
func RestoreMySQL(params context.DBSessionContext) {
	backupFileContent, err := os.ReadFile(params.BackupDir + "/" + params.RestoreFile)
	if err != nil {
		logging.Logger.LogError("Failed to read backup file: %v", err)
		panic(err)
	}

	restoreCommand := exec.Command("mysql",
		"-h", params.DBHost,
		"-u", params.DBUsername,
		"-P", params.DBPort,
		"-p"+params.DBPassword,
		params.DBName,
	)

	restoreCommand.Stdin = bytes.NewReader(backupFileContent) // File from where SQL reads

	executor.ExecuteCommand(restoreCommand, "restore", params.DBType, params.RestoreFile)
}
