package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Class struct {
	Id          int    `db:"id"`
	Cname       string `db:"cname"`
	Description string `db:"description"`
}

var Db *sqlx.DB

func sqlSelect () {
	var p []Class
	err := Db.Select(&p, "select * from class")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)
}

func sqlInsert() {
	result, err := Db.Exec("insert into class(cname,description) values(?,?)",  "ruby", "过时了")
	if err != nil {
		fmt.Println(err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)
}

func main() {
	open, err := sqlx.Open("mysql", "root:MyPassword@123@tcp(127.0.0.1:3306)/shop")
	if err != nil {
		fmt.Println(err)
		return
	}
	Db = open
	defer func() {
		_ = Db.Close()
	}()
	sqlInsert()
	sqlSelect()
}
