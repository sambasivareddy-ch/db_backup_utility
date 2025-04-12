package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

/*
	rootCmd represents the root command
	Usage: db_backup_utility
	Description: This command is used to manage database backups and restores.
*/
var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Supports Database Backup Utility",
	Long:  "Supports Database Backup Utility for various databases like PG, SQL",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello DB World!!")
		fmt.Println("Run config command first to configure your CLI")
	},
}

func Execute() {
	rootCmd.Execute()
}
