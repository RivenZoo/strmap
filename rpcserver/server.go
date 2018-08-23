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
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/metadata"
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

func (s *strmapServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	fmt.Printf("token: %s, err: %v\n", token, err)
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Printf("metadata: %v, %t, method: %s\n", md, ok, fullMethodName)
	fmt.Printf(":authority: %s\n", md.Get(":authority"))
	fmt.Printf(":method: %s\n", md.Get(":method"))
	return ctx, nil
}

func authorizeToken(ctx context.Context) (context.Context, error) {
	//authHeader := grpc_auth.AuthFromMD()
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Printf("metadata: %v, %t\n", md, ok)
	return ctx, nil
}

func StartRPCServer(ctx context.Context, listen string) error {
	server := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_recovery.UnaryServerInterceptor(),
		grpc_auth.UnaryServerInterceptor(authorizeToken),
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
