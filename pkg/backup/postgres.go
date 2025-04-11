package backup

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/sambasivareddy-ch/db_backup_utility/context"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/executor"
)

// Utility function to Backup the PG using pg_dump
func BackupPostgreSQL(params context.DBSessionContext) {
	current_time := time.Now().Format(time.RFC3339)
	backup_file := fmt.Sprintf("%s/pg_backup_%s.dump", params.BackupDir, current_time)

	os.Setenv("PGPASSWORD", params.DBPassword) // Set Env for the Password

	backupCommand := exec.Command("pg_dump",
		"-h", params.DBHost,
		"-U", params.DBUsername,
		"-p", params.DBPort,
		"-d", params.DBName,
		"-F", "c", // SQL format
		"-f", backup_file, // Backup file
	)

	backupCommand.Env = append(os.Environ(), "PGPASSWORD="+params.DBPassword)

	executor.ExecuteCommand(backupCommand, "backup", params.DBType, backup_file)
}

// Utility function to Restore the PG using pg_restore
func RestorePostgreSQL(params context.DBSessionContext) {
	os.Setenv("PGPASSWORD", params.DBPassword) // Set Env for the Password

	restoreCommand := exec.Command("pg_restore",
		"-h", params.DBHost,
		"-U", params.DBUsername,
		"-p", params.DBPort,
		"-d", params.DBName,
		params.BackupDir+"/"+params.RestoreFile, // Backup file complete path
	)

	restoreCommand.Env = append(os.Environ(), "PGPASSWORD="+params.DBPassword)

	executor.ExecuteCommand(restoreCommand, "restore", params.DBType, params.RestoreFile)
}
