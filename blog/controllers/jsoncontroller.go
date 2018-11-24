package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type JsonController struct {
	beego.Controller
}

type Student struct {
	Name     string `json:"name"`
	Age      int
	Guake    bool
	Classes  []string
	Price    float32
	AddTjson string `json:"add"`
}

func (c *JsonController) Get() {
	//var obj string = "{"a":1, "b":2}"
	//jstr := json.Marshal(obj)
	st := &Student{
		"Xiao Ming",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99,
		"add_test---",
	}
	b, _ := json.Marshal(st)
	c.Ctx.WriteString(string(b))
	//c.Data["json"] = st;
	return
}

func (c *JsonController) Post() {
	var jsonBlob = []byte(`[
    	{"Name": "Platypus", "Age": 1},
    	{"Name": "Quoll",    "Age": 2}
		]`)
	c.Ctx.WriteString(string(jsonBlob))
}
