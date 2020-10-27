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
	beego.Router("/api/list-rulefiles", &controllers.ApiController{}, "GET:ListRulefiles")
	beego.Router("/api/list-rules", &controllers.ApiController{}, "GET:ListRules")

	beego.Router("/api/get-testcases", &controllers.ApiController{}, "GET:GetTestcases")
	beego.Router("/api/get-testcase", &controllers.ApiController{}, "GET:GetTestcase")
	beego.Router("/api/update-testcase", &controllers.ApiController{}, "POST:UpdateTestcase")
	beego.Router("/api/add-testcase", &controllers.ApiController{}, "POST:AddTestcase")
	beego.Router("/api/delete-testcase", &controllers.ApiController{}, "POST:DeleteTestcase")
}
