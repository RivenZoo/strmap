package main

import (
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start string mapping service",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
