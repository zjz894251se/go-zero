package config

import "github.com/zjz894251se/go-zero/core/logx"

type Config struct {
	logx.LogConf
	ListenOn string
	FileDir  string
	FilePath string
}
