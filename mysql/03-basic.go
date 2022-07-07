package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	Name  string  `db:"name"`
	Age   int     `db:"age"`
	Money float64 `db:"rmb"`
}

var Db *sqlx.DB

func cud() {
	//res, err := Db.Exec("insert into person(name, age, rmb, gender, birthday) values(?,?,?,?,?);", "张没蛋", 20, 1343, "女", "2020-3-29")
	//_, err := Db.Exec("delete from person where name not like ?;", "%蛋")
	res, err := Db.Exec("update person set name = ? where id = ?;", "张二", 1)
	if err != nil {
		fmt.Println("语句执行失败", err)
		return
	}
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
}

func query() {
	var ps []Person

	// 执行查询，得到Person对象的集合，丢入预定义的ps地址
	err := Db.Select(&ps, "select name, age, rmb from person where name like ?", "%蛋")
	if err != nil {
		fmt.Println("查询失败", err)
		return
	}
	fmt.Println(ps)
}

func main() {
	db, err := sqlx.Open("mysql", "root:MyPassword@123@tcp(127.0.0.1:3306)/shop")
	if err != nil {
		fmt.Println("数据库连接错误")
		return
	}
	defer func() {
		_ = db.Close()
	}()
	Db = db
	//cud()
	query()
}
