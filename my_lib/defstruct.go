package mlib

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"time"
)

type Student struct {
	Id    int    `orm:"pk(id)"`
	Name  string `json:"name"`
	Age   int
	Guake bool
	//Classes    []*string `orm:"reverse(many, 3)"`
	Price float32
}

type Sempty struct {
	Name string
}

func init() {
	orm.RegisterDataBase("default",
		"mysql",
		"root:feng123@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(Student), new(AuditRule), new(Rule))
	orm.RunSyncdb("default", false, true)
}

func Testx() {
	fmt.Printf("____________________done_______________________\n")
}

//测试获取audit_rule

type Rule struct {
	Id        int    `orm:"pk(auto)"`
	App_id    string `orm:"size(64);column(app_id)"`
	Rule_name string `orm:"column(rule_name)"` //名字不写下划线是获取不了的，RuleName这样是不行的
	Modified  string
	//CreateTime  time.Time
	//Audit_str	string  //为什么获取不了，，
}
type AuditRule struct {
	Id    int
	Error string `orm:"size(10)"`
	Errno int
	Rule  []*Rule `orm:"rel(m2m)"`
}
