package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id     int
	NameAx string `orm:"size(100);column(Name)"` //1.把name_ax改为Name
}

//2. 自定义表名（系统自动调用）,否则是表明user
func (u *User) TableName() string {
	return "test_user"
}

//3.转化为json
func (t *User) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}

//4。init mysql
func init() {
	orm.RegisterDataBase("default", "mysql",
		"root:feng123@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("test")

	//5.直接执行sql语句
	//sql := fmt.Sprintf("insert into test_user(Id,Name)"+ " values(1000, 'rjx')")
	sql := "insert into test_user(Id,Name)" + " values(1000, 'rjx')"
	fmt.Println(sql)
	_, err := o.Raw(sql).Exec()
	if err != nil {
		fmt.Println("插入数据至：t_studentInfo出错")
	}

	user := User{NameAx: "slene"}
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	user1 := User{NameAx: "tom"}
	id, err3 := o.Insert(&user1)
	fmt.Printf("ID: %d, ERR: %v\n", id, err3)

	u := User{Id: user.Id}
	err1 := o.Read(&u)
	fmt.Printf("result: %s, ERR: %v\n", u.String(), err1)
	fmt.Printf("%d \n", user.Id)

	//num, err2 := o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err2)

	InsertMutil()
	ReadMutil()
}

//6.插入多个。。。。
func InsertMutil() {
	o := orm.NewOrm()
	user := []User{{NameAx: "aa"}, {NameAx: "bb"}}
	o.InsertMulti(2, &user)
}

//7.查询多个。。。。
func ReadMutil() {
	o := orm.NewOrm()
	// 获取 QuerySeter 对象，user 为表名
	//qs := o.QueryTable("user")
	// 也可以直接使用对象作为表名
	//user := new(User)
	//qs = o.QueryTable(user) // 返回 QuerySete
	//qs.Filter("id__in", 1, 20)
	//qs.Filter("id__in", 18, 20).Exclude("profile__lt", 1000)
	var users []*User
	num, err := o.QueryTable("test_user").Filter("id__in", 1, 2, 1000).All(&users)
	//num, err := o.QueryTable("test_user").Filter("id",1).All(&users)
	fmt.Printf("Returned Rows Num: %d, %s", num, err)
	fmt.Println(users)

	b, _ := json.Marshal(users)
	fmt.Println("\n结果" + string(b))

	var uArr []User
	json.Unmarshal([]byte(b), &uArr)
	for _, v := range uArr {
		fmt.Println(v.NameAx + string(v.Id))
	}
}
