package cmd

import (
	"fmt"

	"github.com/sambasivareddy-ch/db_backup_utility/context"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/utils"

	"github.com/spf13/cobra"
)

// Root Command
var configCommand = &cobra.Command{
	Use:   "config",
	Short: "Configure your Database Backup Utility",
	Long:  "Configure your Database Backup Utility with following details Host, Port, Database, User, Password etc.",
	Run: func(cmd *cobra.Command, args []string) {
		var dbCtx *context.DBSessionContext = &context.DBSessionContext{}
		utils.RunConfigInteractiveInputTerminal(dbCtx)
		dbCtx.SaveToFile()
		fmt.Println("now you can run backup & restore commands")
	},
}

func init() {
	rootCmd.AddCommand(configCommand)
}
