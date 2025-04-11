package cmd

import (
	"fmt"

	"github.com/sambasivareddy-ch/db_backup_utility/pkg/cron"
	"github.com/spf13/cobra"
)

var cronSchedule, operation, dbtype string

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedule a backup task",
	Long:  "Schedule a backup task for a specific database at a given time",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Scheduling a backup task...")
		if cronSchedule == "" || operation == "" || dbtype == "" {
			fmt.Println("Please provide a valid cron schedule, operation, and database type.")
			return
		}
		// Validate the cron schedule format
		if !cron.IsValidCron(cronSchedule) {
			fmt.Println("Invalid cron schedule format. Please use a valid cron expression.")
			return
		}
		// Call the backup function with the provided cron schedule and operation
		cron.ScheduleBackup(cronSchedule, operation, dbtype)
	},
}

func init() {
	// Define flags for the schedule command
	scheduleCmd.Flags().StringVarP(&cronSchedule, "cron", "c", "", "Cron schedule expression (e.g., '0 0 * * *')")
	scheduleCmd.Flags().StringVarP(&operation, "operation", "o", "", "Operation to perform (backup/restore)")
	scheduleCmd.Flags().StringVarP(&dbtype, "dbtype", "d", "", "Database type (e.g., postgres, mysql)")

	scheduleCmd.MarkFlagRequired("cron")
	scheduleCmd.MarkFlagRequired("operation")
	scheduleCmd.MarkFlagRequired("dbtype")

	// Append 'schedule' command to root
	rootCmd.AddCommand(scheduleCmd)
}
