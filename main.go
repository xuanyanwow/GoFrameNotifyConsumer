package main

import (
	_ "gf-app/boot"
	_ "gf-app/router"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

func main() {
	glog.SetDebug(false)
	g.Server().SetDumpRouterMap(false)

	g.Server().Run()
}
