package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	//获取所有列表
	op := c.Input().Get("op")
	beego.Error(op)
	switch op {
	case "add":
		{
			beego.Info("\n_______________________add____________")
			name := c.Input().Get("name")
			if len(name) == 0 {
				break
			}
			err := models.AddCategory(name) //加到数据库
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/category", 301)
			return
		}
	case "del":
		{
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}

			beego.Info("\n----------------del------------")
			err := models.DellCategory(id)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/category", 301)
			return
		}
	}

	beego.Info("\n________________________________________switch done__\n")
	c.Data["IsCategory"] = true
	c.TplName = "category.html"

	var err error
	c.Data["Categories"], err = models.GetAllCategory()

	if err != nil {
		beego.Error(err)
	}
}
