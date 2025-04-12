package cmd

import (
	"fmt"

	"github.com/sambasivareddy-ch/db_backup_utility/pkg/cron"
	"github.com/spf13/cobra"
)

var cronSchedule, operation string

/*
scheduleCmd represents the schedule command
Usage: db_backup_utility schedule -c <cron_schedule> -o <operation>
Description: This command is used to schedule a backup task for a specific database at a given time.
You can specify the cron schedule and the operation (backup/restore) to perform.
Example: db_backup_utility schedule -c "0 0 * * *" -o "backup"
Note: The cron schedule should be in the format of a standard cron expression.
For example, "0 0 * * *" means every day at midnight.
You can use the following cron schedule format:
  - * * * *  -  minute (0-59)
    | | | | |
    | | | | +----- hour (0-23)
    | | | +------- day of month (1-31)
    | | +--------- month (1-12)
    | +----------- day of week (0-7) (Sunday is both 0 and 7)
    +------------- year (optional)
*/
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedule a backup task",
	Long:  "Schedule a backup task for a specific database at a given time",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Scheduling a backup task...")

		// Check if the cron schedule and operation are provided
		if cronSchedule == "" || operation == "" {
			fmt.Println("Please provide a valid cron schedule, operation, and database type.")
			return
		}
		// Validate the cron schedule format
		if !cron.IsValidCron(cronSchedule) {
			fmt.Println("Invalid cron schedule format. Please use a valid cron expression.")
			return
		}
		// Call the backup function with the provided cron schedule and operation
		cron.ScheduleBackup(cronSchedule, operation)
	},
}

func init() {
	// Define flags for the schedule command
	scheduleCmd.Flags().StringVarP(&cronSchedule, "cron", "c", "", "Cron schedule expression (e.g., '0 0 * * *')")
	scheduleCmd.Flags().StringVarP(&operation, "operation", "o", "", "Operation to perform (backup/restore)")

	// Mark the flags as required
	scheduleCmd.MarkFlagRequired("cron")
	scheduleCmd.MarkFlagRequired("operation")

	// Append 'schedule' command to root
	rootCmd.AddCommand(scheduleCmd)
}
