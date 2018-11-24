package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db := opendb("root:feng123@tcp(127.0.0.1:3306)/test?charset=utf8")
	//db := opendb("test:123456@tcp(192.168.127.128:3306)/test?port=3306&charset=utf8")
	id := insert(db)
	query(db)
	update(db, id)
	del(db, 2)

}

//打开数据库连接
func opendb(dbstr string) *sql.DB {
	log.Println("xxxxxxxxxxxxxxxxxxxxx-")
	//dsn: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&paramN=valueN]
	db, err := sql.Open("mysql", dbstr)
	log.Print("log", err)
	prerr(err)
	return db
}

//插入数据
func insert(db *sql.DB) int64 {

	stmt, err := db.Prepare("INSERT INTO student SET id=?, name=?,age=?,created=?")
	prerr(err)

	res, err := stmt.Exec(0, "abloz1", 28, "2013-8-20")
	prerr(err)

	id, err := res.LastInsertId()
	prerr(err)
	log.Println(id)

	fmt.Println(id)
	return id

}

//更新数据
func update(db *sql.DB, id int64) {
	stmt, err := db.Prepare("update student set name=? where id=?")
	prerr(err)

	res, err := stmt.Exec("abloz2", id)
	prerr(err)

	affect, err := res.RowsAffected()
	prerr(err)

	fmt.Println(affect)
}

//查询数据
func query(db *sql.DB) {

	rows, err := db.Query("SELECT * FROM student")
	prerr(err)

	for rows.Next() {
		var id int
		var name string
		var department string
		var created string
		err = rows.Scan(&id, &name, &department, &created)
		prerr(err)
		fmt.Printf("[%d] %s %s %s \n", id, name, department, created)
	}
}

//删除数据
func del(db *sql.DB, id int64) {
	stmt, err := db.Prepare("delete from student where id=?")
	prerr(err)

	res, err := stmt.Exec(id)
	prerr(err)

	affect, err := res.RowsAffected()
	prerr(err)

	fmt.Printf("delete = %d", affect)
}
func prerr(err error) {
	if err != nil {
		panic(err)
	}
}
