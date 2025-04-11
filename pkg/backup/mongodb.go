package backup

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/sambasivareddy-ch/db_backup_utility/context"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/executor"
)

// Utility function to Backup the MongoDB using mongodump
func BackupMongoDB(params context.DBSessionContext) {
	current_time := time.Now().Format(time.RFC3339)
	backup_folder := fmt.Sprintf("%s/mongodb_backup_%s", params.BackupDir, current_time)

	backupCommand := exec.Command("mongodump",
		"--uri="+fmt.Sprintf("mongodb://%s:%s", params.DBHost, params.DBPort),
		"--out="+backup_folder,
	)

	executor.ExecuteCommand(backupCommand, "backup", params.DBType, backup_folder)
}

// Utility function to Restore the MongoDB using mongorestore
func RestoreMongoDB(params context.DBSessionContext) {
	backup_folder := fmt.Sprintf("%s/%s", params.BackupDir, params.RestoreFile)

	restoreCommand := exec.Command("mongorestore",
		"--uri="+fmt.Sprintf("mongodb://%s:%s", params.DBHost, params.DBPort),
		backup_folder,
	)

	executor.ExecuteCommand(restoreCommand, "restore", params.DBType, params.RestoreFile)
}
