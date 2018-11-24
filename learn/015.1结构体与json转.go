package main

//http://blog.csdn.net/kic18/article/details/72820205
import (
	"encoding/json"
	"fmt"
)

type appInfo struct {
	Appid string `json:"appId"`
}

type response struct {
	RespCode string  `json:"respCode"`
	RespMsg  string  `json:"respMsg"`
	AppInfo  appInfo `json:"app"` //这样后，key就是小写的app
}

func (t *response) String() string {
	b, _ := json.Marshal(t) //转为json
	return string(b)
}

type JsonResult struct {
	Resp response `json:"resp"`
}

func main() {
	jsonstr := `{"resp": {"respCode": "000000","respMsg": "成功","app": {"appId": "d12abd3da59d47e6bf13893ec43730b8"}}}`

	fmt.Println("str = %s", jsonstr)
	var JsonRes JsonResult

	json.Unmarshal([]byte(jsonstr), &JsonRes)
	fmt.Println("after parse", JsonRes)

	fmt.Println(JsonRes.Resp.String())

	fmt.Println("____________________________")
	main2()
}

type ColorGroup struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Colors []string `json:"colors"`
}

func main2() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
