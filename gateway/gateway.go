package gateway

import (
	"net/http"
	"strmap/proto/apidef"
	sigutil "strmap/signal"
	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/balancer/roundrobin"
	"context"
	"time"
)

func StartGateway(endpoint, listen string) error {
	ctx := sigutil.RegisterDoneSignal()

	mux := runtime.NewServeMux()
	dialOpt := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name)}
	err := apidef.RegisterStringMapHandlerFromEndpoint(ctx, mux, endpoint, dialOpt)
	if err != nil {
		return err
	}

	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger())
	engine.Any("/", gin.WrapH(mux))
	server := http.Server{
		Addr:    listen,
		Handler: engine,
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
