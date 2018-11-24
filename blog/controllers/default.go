package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	//c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "home.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.html"

	//c.TplName = "index.tpl"

	c.Data["trueCond"] = true
	c.Data["falseCond"] = false

	type u struct {
		Name string
		Age  int
		Sex  string
	}
	//模板渲染结构
	user := &u{
		Name: "name_jack",
		Age:  23,
		Sex:  "boye",
	}

	c.Data["user"] = user

	arr := []int{1, 2, 3, 4, 6, 7, 8, 9}
	c.Data["arr"] = arr
	c.Data["tplvar"] = "hi tmp var"

	c.Data["html"] = "<div>div hello </dvi>"

	c.Data["pipe"] = "<div>pipe, div hello </dvi>"

}
