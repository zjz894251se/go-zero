package config

import (
	"github.com/zjz894251se/go-zero/core/stores/cache"
	"github.com/zjz894251se/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	Cache      cache.CacheConf
}
