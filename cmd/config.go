package cmd

import (
	"fmt"

	"github.com/sambasivareddy-ch/db_backup_utility/context"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/utils"

	"github.com/spf13/cobra"
)

/*
	configCommand represents the config command
	Usage: db_backup_utility config
	Description: This command is used to configure the database connection details.
	You can set the host, port, database name, user, and password for your database.
*/
var configCommand = &cobra.Command{
	Use:   "config",
	Short: "Configure your Database Backup Utility",
	Long:  "Configure your Database Backup Utility with following details Host, Port, Database, User, Password etc.",
	Run: func(cmd *cobra.Command, args []string) {
		var dbCtx *context.DBSessionContext = &context.DBSessionContext{}

		// Get the Configuation details from the user
		utils.RunConfigInteractiveInputTerminal(dbCtx)

		// Save the configuration to a file called db_context.json
		dbCtx.SaveToFile()
		fmt.Println("now you can run backup & restore commands")
	},
}

func init() {
	rootCmd.AddCommand(configCommand)
}
