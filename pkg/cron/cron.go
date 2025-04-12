package cron

import (
	"fmt"
	"os"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sambasivareddy-ch/db_backup_utility/context"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/backup"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/logging"
)

/*
	ScheduleBackup function to schedule the backup and restore operations
	This function takes a cron expression and an operation type (backup/restore)
	and schedules the corresponding operation using the cron library.
	It uses the cron expression to determine when to run the operation.
	It uses the context package to load the session context and perform the operation.
	Note: The cron expression should be in the format of a standard cron expression.
	For example, "0 0 * * *" means every day at midnight.
*/
func ScheduleBackup(cronExpr, operation string) {
	c := cron.New()

	context.LoadSession()

	_, err := c.AddFunc(cronExpr, func() {
		fmt.Println("Scheduled Backup Started")
		switch operation {
		case "backup":
			switch context.GlobalSessionCtx.DBType {
			case "postgres":
				backup.BackupPostgreSQL(*context.GlobalSessionCtx)
			case "sql":
				backup.BackupMySQL(*context.GlobalSessionCtx)
			default:
				fmt.Println("Unsupported Database Type")
			}
		case "restore":
			switch context.GlobalSessionCtx.DBType {
			case "postgres":
				backup.RestorePostgreSQL(*context.GlobalSessionCtx)
			case "sql":
				backup.RestoreMySQL(*context.GlobalSessionCtx)
			default:
				fmt.Println("Unsupported Database Type")
			}
		default:
			fmt.Println("Unsupported Operation")
		}
	})

	if err != nil {
		logging.Logger.LogError("Error occurred while scheduling the backup: ", err)
		os.Exit(1)
	}

	logging.Logger.LogInfo("Scheduled Backup with cron expression: ", cronExpr)
	logging.Logger.LogInfo("Backup operation starts at: %v", time.Now().Format("2006-01-02 15:04:05"))

	c.Start()
	defer c.Stop()

	// Keep the program running to allow cron jobs to execute
	select {}
}

/*
	IsValidCron function to validate the cron expression
	This function takes a cron expression as input and checks if it is valid.
	It uses the cron library to parse the expression and returns true if it is valid,
	and false otherwise.
*/
func IsValidCron(cronExpr string) bool {
	_, err := cron.ParseStandard(cronExpr)
	return err == nil
}
