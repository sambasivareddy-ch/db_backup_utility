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

// ScheduleBackup function to schedule the backup using cron
func ScheduleBackup(cronExpr, operation, dbtype string) {
	c := cron.New(cron.WithSeconds())

	context.LoadSession()

	_, err := c.AddFunc(cronExpr, func() {
		fmt.Println("Scheduled Backup Started")
		switch operation {
		case "backup":
			switch dbtype {
			case "postgres":
				backup.BackupPostgreSQL(*context.GlobalSessionCtx)
			case "sql":
				backup.BackupMySQL(*context.GlobalSessionCtx)
			default:
				fmt.Println("Unsupported Database Type")
			}
		case "restore":
			switch dbtype {
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

func IsValidCron(cronExpr string) bool {
	_, err := cron.ParseStandard(cronExpr)
	if err != nil {
		return false
	}
	return true
}
