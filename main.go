package main

import (
	_ "learn-beego/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	web.Run()
}
