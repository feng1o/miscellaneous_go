package main

import "github.com/astaxie/beego"

type MyController struct {
	beego.Controller
}

func (this *MyController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	beego.Router("/", &MyController{})
	beego.Run()
}
