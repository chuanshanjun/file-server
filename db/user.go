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

func UserSignIn(username string, encpassword string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"select user_pwd from tbl_user where user_name = ? limit 1")
	if err != nil {
		log.Fatalf("get user error:%v", err)
		return false
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		log.Fatalf("get user error:%v", err)
		return false
	}
	if rows == nil {
		log.Fatalf("get user error:%v", err)
		return false
	}

	pRows := mydb.ParseRows(rows)
	// len() 接受map数组, 其实就是看下数组长度
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encpassword {
		return true
	}
	return false
}

func UpdateToken(username string, token string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"replace into tbl_user_token (`user_name`,`user_token`) values (?,?)")
	if err != nil {
		log.Fatalf("update tbl_user error:%v", err)
		return false
	}
	defer stmt.Close()

	exec, err := stmt.Exec(username, token)
	if err != nil {
		log.Fatalf("update tbl_user error:%v", err)
		return false
	}

	if affected, err := exec.RowsAffected(); nil == err && affected > 0 {
		return true
	}

	return false
}
