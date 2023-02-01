package main

import (
	"git.gorio.top/xw/src/server"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

//var Database *sqlx.DB
//
//func init() {
//	database, ok := initDB.InitDB()
//	if !ok {
//		return
//	}
//	Database = database
//}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go server.InitHttpServer()
	wg.Wait()
}
