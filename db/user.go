package db

import (
	mydb "github.com/chuanshan/file-server/db/mysql"
	"log"
)

func UserSignUP(username string, password string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user(`user_name`, `user_pwd`) values(?,?)")
	if err != nil {
		log.Fatal("insert into user error", err)
		return false
	}
	defer stmt.Close()

	exec, err := stmt.Exec(username, password)
	if err != nil {
		log.Fatal("insert into user error", err)
		return false
	}

	if rowsAffected, err := exec.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}
	return false
}
