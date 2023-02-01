package server

import (
	"fmt"
	"git.gorio.top/xw/src/initDB"
	"git.gorio.top/xw/src/sqlFunctions"
	"git.gorio.top/xw/src/utils"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func InitHttpServer() {
	var c utils.YmlConfig
	config := c.GetYml()
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":"+config.HttpPort, nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var Database *sqlx.DB
	database, ok := initDB.InitDB()
	if !ok {
		return
	}
	Database = database
	//res, _ := sqlFunctions.Query(Database, "select * from user")
	res, _ := sqlFunctions.GetJSON(Database, "select * from user")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	fmt.Fprint(w, res)
}
