package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Root Command
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
