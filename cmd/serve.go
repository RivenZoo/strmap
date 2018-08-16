package main

import (
	"strmap/rpcserver"
	"strmap/config"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start string mapping service",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		rpcserver.StartRPCServer(cfg.Listen)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
