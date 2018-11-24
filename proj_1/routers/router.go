package routers

import (
	"github.com/astaxie/beego"
	"proj_1/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
