package main

import (
	"fmt"
	"time"

	"github.com/zjz894251se/go-zero/core/logx"
)

func main() {
	logx.MustSetup(logx.LogConf{
		Mode: "console",
	})
	logx.CollectSysLog()

	line := "asdkg"
	logx.Info(line)
	fmt.Print(line)
	time.Sleep(time.Second)
}
