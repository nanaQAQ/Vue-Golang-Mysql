package sqlFunctions

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

//Exec
/*
 * @author 	nana
 * @description  DO SQL except query
 * @input	*sqlx.DB	database
 * @input	string	SQL
 * @return		bool
 */
func Exec(database *sqlx.DB, SQL string) bool {
	res, err := database.Exec(SQL)
	if err != nil {
		log.Println("Exec SQL error: ", err)
		return false
	}
	insID, _ := res.LastInsertId()
	log.Println(insID)
	return true
}

//Query
/*
 * @author 	nana
 * @description
 * @input	*sqlx.DB	database
 * @input	string	SQL
 * @return	[]map[string]string	${ret_name}
 * @return	bool	${ret_name}
 */
func Query(database *sqlx.DB, SQL string) ([]map[string]string, bool) {
	rows, err := database.Query(SQL) //执行SQL语句，比如select * from users
	if err != nil {
		panic(err)
	}
	cols, _ := rows.Columns() //获取列的信息 	//列的数量
	//values是每个列的值，这里获取到byte里
	values := make([][]byte, len(cols))
	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans := make([]interface{}, len(cols))
	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}
	res := make([]map[string]string, 0)
	for rows.Next() {
		_ = rows.Scan(scans...)
		row := make(map[string]string)
		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			row[key] = string(v)
		}
		res = append(res, row)
	}
	fmt.Println(res)
	return res, true
}

//func Query(database *sqlx.DB, SQL string) ([]map[string]string, bool) {
//	rows, err := database.Query(SQL)
//	if err != nil {
//		log.Println("Exec SQL query error:", err)
//		return nil, bool
//	}
//	columns, _ := rows.Columns()
//	columnsType, _ := rows.ColumnTypes()
//	columnsCount := len(columns)
//	log.Println(columns)
//	values := make([]interface{}, columnsCount)
//	for rows.Next() {
//		rows.Scan(values...)
//		fmt.Println(values)
//	}
//	database.Select()
//}

//https://blog.csdn.net/weixin_34343000/article/details/89025070
func GetJSON(db *sqlx.DB, sqlString string) (string, error) {
	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}
