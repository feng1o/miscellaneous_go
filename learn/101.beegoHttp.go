package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
)

func main() {

	b := httplib.Post("http://beego.me/")
	b.Param("username", "astaxie")
	b.Param("password", "123456")
	b.PostFile("uploadfile1", "httplib.pdf")
	b.PostFile("uploadfile2", "httplib.txt")
	str, err := b.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

	req := httplib.Get("http://beego.me/")

	str1, err1 := req.String()
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(str1)
}
