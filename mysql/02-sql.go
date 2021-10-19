package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Stu struct {
	Id int `db:"id"`
	StuName string `db:"stuname"`
	Age     int    `db:"age"`
}

func main() {
	db, err := sql.Open("mysql", "root:123456@(localhost:3306)/shop")
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}

	defer func() {
		_ = db.Close()
	}()

	//	执行数据操作语句
	/*
		sqlString := "create table if not exists stu3 (id int primary key auto_increment,stuname varchar(20) not null,age tinyint default null)"
		result, _ := db.Exec(sqlString)

		affected, err := result.RowsAffected() // 获取受影响的记录数
		if err != nil {
			fmt.Println("更新数据", err)
			return
		}

		fmt.Println(affected)
	*/

	//	执行预处理
	/*
		stus := [3]Stu{}
		stus[0] = Stu{"张三", 15}
		stus[1] = Stu{"李四", 46}
		stus[2] = Stu{"王五", 30}

		stmt, err := db.Prepare("insert into stu3(stuname,age) values (?,?)")
		if err != nil {
			fmt.Println("45行", err)
			return
		}
		for _, v := range stus {
			result, err := stmt.Exec(v.StuName, v.Age)
			if err != nil {
				fmt.Println("50行", err)
				return
			}
			affected, err := result.RowsAffected()
			if err != nil {
				fmt.Println("54行", err)
				return
			}
			fmt.Println(affected)
		}
	*/

	// 单行查询
	/*
		var (
			id int
			name string
			age int
		)
		row := db.QueryRow("select * from stu3")
		_ = row.Scan(&id, &name, &age)
		fmt.Println(id, name, age)
	*/

	// 多行查询
	stus := make([]Stu,1)
	rows, _ := db.Query("select * from stu3")
	for {
		var stu Stu
		_ = rows.Scan(&stu.Id, &stu.StuName, &stu.Age)
		stus = append(stus, stu)
		if flag := rows.Next();!flag {
			break
		}
	}
	fmt.Println(stus)
}
