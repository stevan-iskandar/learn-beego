// https://www.bookstack.cn/read/beego-2.0.7-en/web-router-ctrl_style-README.md
package main

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/stevan-iskandar/learn-beego/database"
	_ "github.com/stevan-iskandar/learn-beego/routers"
)

func main() {
	database.InitDB()

	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	web.Run()
}
