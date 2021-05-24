package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql",
		"root:123456@tcp(115.159.34.56:33306)/young?charset=utf8")
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// DBConn:返回数据库连接对象
func DBConn() *sql.DB {
	return db
}

func ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		//将行数据保存到record字典
		err := rows.Scan(scanArgs...)
		checkErr(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		records = append(records, record)
	}
	return records
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
