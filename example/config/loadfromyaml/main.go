package main

import (
	"time"

	"github.com/zjz894251se/go-zero/core/conf"
	"github.com/zjz894251se/go-zero/core/logx"
)

type TimeHolder struct {
	Date time.Time `json:"date"`
}

func main() {
	th := &TimeHolder{}
	err := conf.LoadConfig("./date.yml", th)
	if err != nil {
		logx.Error(err)
	}
	logx.Infof("%+v", th)
}
