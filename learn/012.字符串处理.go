package main

import (
	//"errors"
	"fmt"
	//"io"
	"encoding/xml"
	"strconv"
	"strings"
)

func main() {
	var s = "hello world"

	fmt.Println(strings.Contains(s, "hello"))

	fmt.Println(strings.Index(s, "o"))

	ss := "1#2#345"
	splitedStr := strings.Split(ss, "#")
	fmt.Println(splitedStr)

	fmt.Println(strings.Join(splitedStr, "#"))

	fmt.Println(strings.HasPrefix(s, "he"), strings.HasSuffix(s, "ld"))

	//字符串转化
	fmt.Println(strconv.Itoa(2))
	fmt.Println(strconv.Atoi("293"))
	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseFloat("3.14", 64))
	fmt.Println(strconv.FormatInt(123, 2)) //进制转换

	{
		fmt.Println("____________________________xml问题")
		//xml， 结构体数据序列化
		type person struct {
			Name string `xml:"nameattr,attr"` //成为属性
			Age  int    `xml:"xx,attr""`
		}

		p := person{Name: "lfjkl", Age: 12}

		var data []byte
		var err error

		//不能用：=，会不是定义的data
		if data, err = xml.MarshalIndent(p, "---", " "); err != nil {
			fmt.Println(err, string(data))
			return
		} else {
			fmt.Println(string(data))
		}

		p2 := new(person)
		if err = xml.Unmarshal(data, p2); err != nil {
			fmt.Println("err= ", err)
			return
		}

		fmt.Println("p2= ", p2)
	}
}
