/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/Batyachelly/goBoard/internal/goboard"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations",
	Run: func(cmd *cobra.Command, args []string) {
		goboard.Migrate()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
