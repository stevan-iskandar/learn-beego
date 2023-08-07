// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"learn-beego/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	ns := web.NewNamespace("/v1",
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
