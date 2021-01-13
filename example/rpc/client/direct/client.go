package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/zjz894251se/go-zero/core/discov"
	"github.com/zjz894251se/go-zero/example/rpc/remote/unary"
	"github.com/zjz894251se/go-zero/zrpc"
)

const timeFormat = "15:04:05"

func main() {
	flag.Parse()

	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"localhost:2379"},
			Key:   "zrpc",
		},
	})

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			conn := client.Conn()
			greet := unary.NewGreeterClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			resp, err := greet.Greet(ctx, &unary.Request{
				Name: "kevin",
			})
			if err != nil {
				fmt.Printf("%s X %s\n", time.Now().Format(timeFormat), err.Error())
			} else {
				fmt.Printf("%s => %s\n", time.Now().Format(timeFormat), resp.Greet)
			}
			cancel()
		}
	}
}
