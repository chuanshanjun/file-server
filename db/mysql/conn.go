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
