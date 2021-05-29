package db

import (
	"github.com/chuanshan/file-server/db/mysql"
	"log"
)

func OnUserFileUpLoadFinished(fileHash string, filesize int64, filename string, username string) bool {
	stmt, err := mysql.DBConn().Prepare(
		"insert into tbl_user_file(user_name, file_sha1, file_size, file_name) values(?,?,?,?)")
	if err != nil {
		log.Fatalf("insert into user file error:%v\n", err)
		return false
	}
	defer stmt.Close()

	exec, err := stmt.Exec(username, fileHash, filesize, filename)
	if err != nil {
		log.Fatalf("insert into user file error:%v\n", err)
		return false
	}

	if affected, err := exec.RowsAffected(); err == nil {
		if affected == 1 {
			return true
		}
		log.Fatalf("insert into user file affectRows:%d", affected)
		return false
	}
	return false
}
