package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strmap/gateway"
	"strmap/config"
	sigutil "strmap/signal"
	"strmap/tracing"
)

var BackendEndpoint *string

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Start grpc gateway",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		ctx := sigutil.RegisterDoneSignal()
		if cfg.Debug.Trace {
			go tracing.StartGRPCTraceHTTPServer(ctx, cfg.Debug.GRPCTraceAddress)
		}
		gateway.StartGateway(ctx, cfg.Gateway.Endpoint, cfg.Listen)
	},
}

func init() {
	gatewayCmd.Flags().StringP("endpoint", "E", "", "Backend grpc endpoint")
	viper.BindPFlag("gateway.endpoint", gatewayCmd.Flags().Lookup("endpoint"))
	rootCmd.AddCommand(gatewayCmd)
}
