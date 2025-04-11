package cmd

import (
	"fmt"

	"github.com/sambasivareddy-ch/db_backup_utility/context"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/backup"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/utils"

	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Backup Restore Command",
	Long:  "Helps to restore the database based on provided backup file",
	Run: func(cmd *cobra.Command, args []string) {
		context.LoadSession()

		utils.RunRestoreInteractiveInputTerminal(context.GlobalSessionCtx)

		switch context.GlobalSessionCtx.DBType {
		case "postgres":
			backup.RestorePostgreSQL(*context.GlobalSessionCtx)
		case "sql":
			backup.RestoreMySQL(*context.GlobalSessionCtx)
		default:
			fmt.Println("Unsupported Database Type")
		}
	},
}

func init() {
	// Append Restore Command to root
	rootCmd.AddCommand(restoreCmd)
}
