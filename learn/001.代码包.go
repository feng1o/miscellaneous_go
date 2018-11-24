package main

import (
	"fmt"
	"strings"
)
import str "strings" //str变短了
import . "strings"   //直接用方法
import _ "strings"   //初始，不能use

type (
	文本 string
)

const (
	a = 'a'
	b
	c = iota
	d
	e = iota
)

func main() {
	fmt.Printf("hello, world\n") //kang
	str.HasPrefix("anb", "a")
	strings.HasPrefix("anb", "anb")
	HasPrefix("anb", "anb")

	var a 文本
	a = "中文类型别名"
	fmt.Println(a)

	fmt.Printf("%d", d)
}
