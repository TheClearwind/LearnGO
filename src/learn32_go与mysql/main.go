package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var dsn = "root:123456@tcp(49.232.149.165:3306)/learn_go" //user:pwd@tcp(url)/db
func iniDB() (err error) {
	db, err = sql.Open("mysql", dsn) //不校验用户名密码是否正确 只校验dsn格式
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	fmt.Println("连接数据库成功")
	return nil
}

type user struct {
	id   int
	name string
	age  int
}

func query(id int) {
	// 查询单个
	queryStr := "select id,name,age from user where id=? and name =?"
	var u user
	err := db.QueryRow(queryStr, id, "李二妮").Scan(&u.id, &u.name, &u.age) // 需要调用scan 才能取到 同时scan 会将连接归还到连接池
	if err != nil {
		fmt.Printf("scan failed %v\n", err)
		return
	}
	fmt.Printf("%#v\n", u)
}
func querys() {
	// 多行查询
	queryStr := "select id,name,age from user"
	rows, err := db.Query(queryStr)
	if err != nil {
		fmt.Printf("query failed %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed %v\n", err)
			return
		}
		fmt.Printf("%#v\n", u)
	}
}

// sql 预处理
func test() {
	//方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。
	sqlstr := "insert into user(name,age) values(?,?)"
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed %v\n", err)
		return
	}
	defer stmt.Close()
	ret, err := stmt.Exec("琉紫", 20)
	if err != nil {
		fmt.Printf("exec failed %v\n", err)
		return
	}
	fmt.Println(ret.RowsAffected())
}
func main() {
	err := iniDB()
	if err != nil {
		fmt.Printf("连接失败:%v\n", err)
		return
	}

	//db.exec //执行某条sql 语句
	//query(2)

	test()
}
