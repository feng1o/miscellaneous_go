package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"my_lib"
	"net/http"
	"strings"
)

//
//type Student struct {
//	Name    string 	`json:"name"`
//	Age        int
//	Guake    bool
//	Classes    []string
//	Price    float32
//}

func httpGet() {
	//resp, err := http.Get("http://localhost:8080/json")
	resp, err := http.Get("http://10.120.135.245:443/cdb2/fun_logic/cgi-bin/public_api_20/query_audit_rule.cgi")
	//resp, err := http.Get("http://localhost:8080/category")
	//resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	stb := &mlib.Sempty{}
	//stb := &mlib.Student{}
	json.Unmarshal([]byte(body), stb)
	fmt.Println(stb.Name)
	mlib.Testx()

}

func httpPost() []mlib.Student {
	resp, err := http.Post("http://localhost:8080/json", "application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	//stb := &mlib.Student{}
	var stb []mlib.Student
	json.Unmarshal([]byte(body), &stb)
	//fmt.Printf("%v", stb)
	fmt.Printf("%s  %d\n", stb[0].Name, len(stb))
	for _, v := range stb {
		fmt.Printf("name = %s  age = %d \n", v.Name, v.Age)
	}
	return stb
}

func main() {
	//httpGet()
	starr := httpPost()
	fmt.Printf("长度是 %d \n", len(starr))
	//InsertDb(starr)
	httpGetRule()
}

//把student结构插入到表中
func InsertDb(stuArr []mlib.Student) {
	o := orm.NewOrm()
	o.Using("test")

	id, _ := o.InsertMulti(len(stuArr), stuArr)
	fmt.Println(id)

}

func httpGetRule() {
	resp, err := http.Get("http://10.120.135.245:443/cdb2/fun_logic/cgi-bin/public_api_20/query_audit_rule.cgi")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	stb := &mlib.AuditRule{}
	json.Unmarshal([]byte(body), stb)
	fmt.Println(stb.Errno)
	fmt.Println("err == " + stb.Error)
	fmt.Printf("rule size = %d ", len(stb.Rule))
	for _, v := range stb.Rule {
		fmt.Printf("+++++++++ %d %s ", v.Id, v.App_id)

		//mlib.Testx()

		//测试插入mysql
		o := orm.NewOrm()
		id, _ := o.Insert(stb)
		//o.Delete(stb.Rule)
		o.InsertMulti(len(stb.Rule), stb.Rule)
		fmt.Println(id)
	}
}
