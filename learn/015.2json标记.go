package main

import (
	"encoding/json"
	"fmt"
)

/*
在定义 struct 字段的时候，可以在字段后面添加 tag，
来控制 encode/decode 的过程：是否要 decode/encode 某个字段，JSON 中的字段名称是什么。
omitempty
-
FieldName
 */


 func DynamicDecode() {
	 data := []byte(`{"Name":"cizixs","IsAdmin":true,"Followers":36}`)

	 var f interface{}
	 json.Unmarshal(data, &f)

	 name := f.(map[string]interface{})["Name"].(string)  // type assertion
	 fmt.Println(" name = %s", name)

	 //如果不是复杂的，未嵌套结构，可以这样
	 var sf map[string]interface{}
	 json.Unmarshal(data, &sf) // 省去了上面 f 的 type assertion 步骤
	 name2 := sf["Name"].(string)
	 fmt.Println(" name = %s", name2)
 }

func main()  {
	DynamicDecode()
}