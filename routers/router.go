package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stevan-iskandar/learn-beego/controllers"
)

func init() {
	web.Get("/hello", func(ctx *context.Context) {
		ctx.WriteString("hello, world")
	})
	ns := web.NewNamespace("/api",
		web.NSNamespace("/book",
			web.NSInclude(
				&controllers.BookController{},
			),
		),
		web.NSNamespace("/object",
			web.NSInclude(
				&controllers.ObjectController{},
			),
		),
		web.NSNamespace("/user",
			web.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	web.AddNamespace(ns)
}
