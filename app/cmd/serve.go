/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/Batyachelly/goBoard/internal/goboard"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run goBoard server",
	Run: func(cmd *cobra.Command, args []string) {
		migrate, _ := cmd.Flags().GetBool("migrate")

		goboard.Serve(migrate)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().BoolP("migrate", "m", true, "Migrate on run")
}
