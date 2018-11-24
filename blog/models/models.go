package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" //驱动，要初始化注册
	"github.com/unknwon/com"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	_db_name        = "data/blog.db"
	_sqlite3_driver = "sqlite3"
) //常量

type category struct {
	Id              int64 `orm:"pk;auto"`
	Title           string
	Created         time.Time `orm:"index"` //有可能按照这个排序，所以加一个索引，，tag，反射时时候用；
	Views           int64     `orm:"index"` //orm时一个标志，只有rom才会读，
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type topic struct {
	Id              int64 `orm:"pk;auto"`
	Tid             int64
	Title           string
	Content         string    `orm:"size(5000)"`
	Atttachment     string    //附件
	Created         time.Time `orm:"index"` //有可能按照这个排序，所以加一个索引，，tag，反射时时候用；
	Views           int64     `orm:"index"` //orm时一个标志，只有rom才会读，
	Updated         time.Time `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"` //可能通过回复时间排序，建索引
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDb() {
	if !com.IsExist(_db_name) {
		os.MkdirAll(path.Dir(_db_name), os.ModePerm) //创建路径
		os.Create(_db_name)                          //创建db
	}
	//orm注册模型
	orm.RegisterModel(new(category), new(topic))
	orm.RegisterDriver(_sqlite3_driver, orm.DRSqlite) //默认已经注册，注册驱动
	//注册默认数据库
	orm.RegisterDataBase("default", _sqlite3_driver, _db_name, 10)

}

//插入category到数据库
func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &category{Title: name, Created: time.Now(), TopicTime: time.Now(), Views: 1} //创建一个对象 category

	//判定name是否被用了
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	} //zhaodao

	_, err2 := o.Insert(cate)
	if err2 != nil {
		return err2
	}

	return nil
}

//获取所有category
func GetAllCategory() ([]*category, error) {
	o := orm.NewOrm()

	cates := make([]*category, 0)

	qs := o.QueryTable("category")

	_, err := qs.All(&cates)
	return cates, err
}

//del
func DellCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err

	}
	o := orm.NewOrm()

	cate := &category{Id: cid}
	_, err2 := o.Delete(cate)
	if err2 != nil {
		return err2
	}
	return nil
}
