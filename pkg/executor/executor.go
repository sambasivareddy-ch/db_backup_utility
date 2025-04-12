package executor

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/sambasivareddy-ch/db_backup_utility/pkg/logging"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/notify"

	"github.com/fatih/color"
)

/*
	Utility function to Execute the command
	This function takes the command to be executed, the operation type (backup/restore),
	and the database type (MySQL/PostgreSQL).
	It runs the command and logs the output.
	It also sends a notification to Discord with the success or failure message.
	@params: command - The command to be executed
	@params: operation - The type of operation (backup/restore)
	@params: dbType - The type of database (MySQL/PostgreSQL)
	@params: file - The file name of the backup/restore file
*/
func ExecuteCommand(command *exec.Cmd, operation, dbType, file string) {
	logging.Logger.LogCommand(command.String()) // Print the command about to run

	start_time := time.Now()

	color.Blue("🔧 Starting %s %s...", dbType, operation)
	if err := command.Run(); err != nil {
		err_msg := fmt.Sprintf("Failure:\n - Error occurred while doing %s in the %s: %v\n - At: %v",
			operation, dbType, err, time.Now())
		logging.Logger.LogError(err_msg)
		notify.SendNotificationOnDiscord(err_msg)
		panic(err)
	}

	end_time := time.Now()

	var log_message string
	switch operation {
	case "backup":
		log_message = fmt.Sprintf("Success:\n - %s backup successful.\n - Saved to: %s \n - Time taken: %v",
			dbType, file, end_time.Sub(start_time))
	case "restore":
		log_message = fmt.Sprintf("Success:\n - %s restore successful.\n - From backup: %s \n - Time taken: %v",
			dbType, file, end_time.Sub(start_time))
	}

	logging.Logger.LogInfo(log_message)
	notify.SendNotificationOnDiscord(log_message)
	color.Green("%s Successful.", operation)
}
