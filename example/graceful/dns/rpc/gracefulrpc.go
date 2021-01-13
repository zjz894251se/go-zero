package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/zjz894251se/go-zero/core/conf"
	"github.com/zjz894251se/go-zero/example/graceful/dns/rpc/graceful"
	"github.com/zjz894251se/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/config.json", "the config file")

type GracefulServer struct{}

func NewGracefulServer() *GracefulServer {
	return &GracefulServer{}
}

func (gs *GracefulServer) Grace(ctx context.Context, req *graceful.Request) (*graceful.Response, error) {
	fmt.Println("=>", req)

	time.Sleep(time.Millisecond * 10)

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &graceful.Response{
		Host: hostname,
	}, nil
}

func main() {
	flag.Parse()

	var c zrpc.RpcServerConf
	conf.MustLoad(*configFile, &c)

	server := zrpc.MustNewServer(c, func(grpcServer *grpc.Server) {
		graceful.RegisterGraceServiceServer(grpcServer, NewGracefulServer())
	})
	defer server.Stop()
	server.Start()
}
