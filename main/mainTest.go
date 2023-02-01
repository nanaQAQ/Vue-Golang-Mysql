package main

//
//import (
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jmoiron/sqlx"
//)
//
//type Users struct {
//	UserId   int    `db:"user_id"`
//	Username string `db:"username"`
//	Sex      string `db:"sex"`
//	Email    string `db:"email"`
//}
//
//var db *sqlx.DB
//var userName = "root"
//var userPasswd = "root"
//var sqlAddr = "127.0.0.1:3306"
//
////init
///*
// * @author 	nana
// * @description init函数sql连接
// */
//func init() {
//	dataSourceName := userName + ":" + userPasswd + "@tcp(" + sqlAddr + ")/mytest"
//
//	database, err := sqlx.Open("mysql", dataSourceName)
//	if err != nil {
//		fmt.Println("open mysql failed,", err)
//		return
//	}
//	db = database
//	fmt.Println("open mysql succeed")
//}
//
//func main() {
//	sql := "insert into user(username,sex, email)values (?,?,?)"
//	value := [3]string{"user01", "man", "user01@163.com"}
//
//	//执行SQL语句
//	r, err := db.Exec(sql, value[0], value[1], value[2])
//	if err != nil {
//		fmt.Println("exec failed,", err)
//		return
//	}
//	rows, _ := db.Query("select * from user")
//	tmp := Users{}
//	for rows.Next() {
//		rows.Scan(&tmp.UserId, &tmp.Username, &tmp.Sex, &tmp.Email)
//		fmt.Println(tmp)
//	}
//	rows.Columns()
//	//查询最后一天用户ID，判断是否插入成功
//	id, err := r.LastInsertId()
//	if err != nil {
//		fmt.Println("exec failed,", err)
//		return
//	}
//	fmt.Println("insert succ", id)
//}
