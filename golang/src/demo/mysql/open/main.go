package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/beego1")
	db.Ping() // 立刻建立链接
	defer func() {
		if db != nil {
			db.Close()
		}
	}()
	if err != nil {
		fmt.Println("链接数据库失败", err)
		return
	}

	// smt, err := db.Prepare("insert into user (name, password_hash, mobile) values(?, ?, ?)")
	// defer func() {
	// 	if smt != nil {
	// 		smt.Close()
	// 	}
	// }()
	// if err != nil {
	// 	fmt.Println("预处理失败", err)
	// 	return
	// }

	// e, err := smt.Exec("simba4", "123456", "5")
	// if err != nil {
	// 	fmt.Println("执行SQL失败", err)
	// 	return
	// }
	// // 受影响的行数
	// count, err := e.RowsAffected()
	// if err != nil {
	// 	fmt.Println("获取受影响行数失败", err)
	// 	return
	// }

	// if count > 0 {
	// 	fmt.Println("新增成功")
	// } else {
	// 	fmt.Println("新增失败")
	// }
	// // 获取新增主键的值
	// id, err := e.LastInsertId()
	// if err != nil {
	// 	fmt.Println("获取ID失败", err)
	// } else {
	// 	fmt.Println(id)
	// }

	p, err := db.Prepare("select name, password_hash, mobile from user limit 5")
	if err != nil {
		fmt.Println("预处理失败")
		return
	}
	defer func() {
		if p != nil {
			p.Close()
		}
	}()
	rows, err := p.Query()
	if err != nil {
		fmt.Println("获取结果失败")
		return
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	for rows.Next() {
		var name string
		var passwordHash string
		var mobile string
		rows.Scan(&name, &passwordHash, &mobile)
		fmt.Println(name, passwordHash, mobile)
	}
}
