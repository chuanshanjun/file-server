package db

import (
	mydb "github.com/chuanshan/file-server/db/mysql"
	"log"
)

type User struct {
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}

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

// GetUserInfo 查询用户信息
func GetUserInfo(username string) (User, error) {
	user := User{}
	stmt, err := mydb.DBConn().Prepare(
		"select user_name, signup_at from tbl_user where user_name = ? limit 1")
	if err != nil {
		log.Fatalf("GetUserInfo tbl_user error:%v", err)
		return user, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	if err != nil {
		log.Fatalf("GetUserInfo tbl_user error:%v", err)
		return user, err
	}
	return user, nil
}

type UserToken struct {
	Username  string
	UserToken string
}

// 获取用户token
func GetUserToken(username string) (UserToken, error) {
	userToken := UserToken{}
	stmt, err := mydb.DBConn().Prepare(
		"select user_name, user_token from tbl_user_token where user_name = ? limit 1")
	if err != nil {
		log.Fatalf("get user token from tbl_user_token err:%v", err)
		return userToken, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&userToken.Username, &userToken.UserToken)
	if err != nil {
		log.Fatalf("get user token from tbl_user_token err:%v", err)
		return userToken, err
	}

	return userToken, nil
}
