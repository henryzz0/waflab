package routers

import (
	"github.com/astaxie/beego"

	"github.com/waflab/waflab/controllers"
)

func init() {
	initAPI()
}

func initAPI() {
	ns :=
		beego.NewNamespace("/api",
			beego.NSInclude(
				&controllers.ApiController{},
			),
		)
	beego.AddNamespace(ns)

	beego.Router("/api/list-rulesets", &controllers.ApiController{}, "GET:ListRulesets")
}
