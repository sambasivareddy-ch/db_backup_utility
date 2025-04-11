// Backup Command
package cmd

import (
	"fmt"

	"github.com/sambasivareddy-ch/db_backup_utility/context"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/backup"

	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Command to backup",
	Long:  "A Command which helps in backing up various DBs like PG, SQL",
	Run: func(cmd *cobra.Command, args []string) {
		// Load DB Information
		context.LoadSession()

		switch context.GlobalSessionCtx.DBType {
		case "postgres":
			backup.BackupPostgreSQL(*context.GlobalSessionCtx)
		case "sql":
			backup.BackupMySQL(*context.GlobalSessionCtx)
		default:
			fmt.Println("Unsupported Database Type")
		}
	},
}

func init() {
	// Append 'backup' command to root
	rootCmd.AddCommand(backupCmd)
}
