package config

import (
	"github.com/zjz894251se/go-zero/rest"
	"github.com/zjz894251se/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Transform zrpc.RpcClientConf
}
