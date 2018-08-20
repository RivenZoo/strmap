package rpcserver

import (
	"golang.org/x/net/context"
	"strmap/proto/apidef"
	"google.golang.org/grpc"
	"net"
	"crypto/md5"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"fmt"
)

var _ apidef.StringMapServer = &strmapServer{}

type strmapServer struct {
}

func (s *strmapServer) Map(ctx context.Context, req *apidef.StrMapReq) (*apidef.StrMapResp, error) {
	return &apidef.StrMapResp{
		Ret:    0,
		Msg:    "ok",
		Result: fmt.Sprintf("%x", md5.Sum([]byte(req.Payload))),
	}, nil
}

func StartRPCServer(ctx context.Context, listen string) error {

	server := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_recovery.UnaryServerInterceptor(),
	)))
	apidef.RegisterStringMapServer(server, &strmapServer{})
	l, err := net.Listen("tcp", listen)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		server.GracefulStop()
	}()
	return server.Serve(l)
}
