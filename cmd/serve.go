package main

import (
	"strmap/rpcserver"
	"strmap/config"
	"github.com/spf13/cobra"

	sigutil "strmap/signal"
	"strmap/tracing"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start string mapping service",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		ctx := sigutil.RegisterDoneSignal()
		if cfg.Debug.Trace {
			go tracing.StartGRPCTraceHTTPServer(ctx, cfg.Debug.GRPCTraceAddress)
		}
		rpcserver.StartRPCServer(ctx, cfg.Listen)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
