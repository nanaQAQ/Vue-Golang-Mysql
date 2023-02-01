package initDB

import (
	"git.gorio.top/xw/src/utils"
	"github.com/jmoiron/sqlx"
	"log"
)

//var path = "root:root@tcp(127.0.0.1:3306)/mytest"

func getPath() string {
	var c utils.YmlConfig
	config := c.GetYml()
	dbConfig := config.User + ":" + config.Pwd + "@tcp(" + config.Ip + ":" + config.DbPort + ")/" + config.DbName
	return dbConfig
}

//InitDB
/*
 * @author 	nana
 * @description Initialize database
 * @input	string	path
 * @return	*sqlx.DB
 * @return	bool
 */
func InitDB() (*sqlx.DB, bool) {
	database, err := sqlx.Connect("mysql", getPath())
	if err != nil {
		log.Println("Connected to database failed: ", err)
		return nil, false
	}
	log.Println("Connected to database succeeded!")
	return database, true
}
