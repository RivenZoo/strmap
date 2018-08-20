package gateway

import (
	"net/http"
	"strmap/proto/apidef"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/balancer/roundrobin"
	"context"
	"time"
)

func StartGateway(ctx context.Context, endpoint, listen string) error {
	marshaler := &runtime.JSONPb{
		OrigName:     true,
		EmitDefaults: true,
	}
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, marshaler))
	dialOpt := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name)}
	err := apidef.RegisterStringMapHandlerFromEndpoint(ctx, mux, endpoint, dialOpt)
	if err != nil {
		return err
	}

	server := http.Server{
		Addr:    listen,
		Handler: mux,
	}
	go func() {
		select {
		case <-ctx.Done():
			func() {
				tmCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()
				server.Shutdown(tmCtx)
			}()
		}
	}()
	return server.ListenAndServe()
}
