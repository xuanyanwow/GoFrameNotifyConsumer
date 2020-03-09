package router

import (
	"gf-app/app/api/hello"
	"gf-app/app/api/notify"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", hello.Hello)
	})

	s.Group("/notify/", func(group *ghttp.RouterGroup) {
		group.ALL("QueryResidualLength", notify.QueryResidualLength)
		group.ALL("OpenQueue", notify.OpenQueue)
		group.ALL("RefuseQueue", notify.RefuseQueue)
		group.ALL("PushQueue", notify.PushQueue)
	})

}
