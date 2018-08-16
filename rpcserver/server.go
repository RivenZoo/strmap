package rpcserver

import (
	context "golang.org/x/net/context"
	"strmap/proto/apidef"
	sigutil "strmap/signal"
	"google.golang.org/grpc"
	"net"
)

var _ apidef.StringMapServer = &strmapServer{}

type strmapServer struct {
}

func (s *strmapServer) Map(ctx context.Context, req *apidef.StrMapReq) (*apidef.StrMapResp, error) {
	return nil, nil
}

func StartRPCServer(listen string) error {
	ctx := sigutil.RegisterDoneSignal()

	server := grpc.NewServer()
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
