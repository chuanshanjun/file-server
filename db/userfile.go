package db

import (
	"github.com/chuanshan/file-server/db/mysql"
	"log"
)

type UserFile struct {
	UserName   string
	FileHash   string
	FileSize   int64
	FileName   string
	UploadAt   string
	LastUpdate string
}

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

// 查询用户文件信息
func QueryUserFileMetas(username string, limit int) ([]UserFile, error) {
	stmt, err := mysql.DBConn().Prepare(
		"select file_sha1, file_size, file_name, upload_at, last_update from tbl_user_file where user_name = ? limit ?")
	if err != nil {
		log.Fatalf("query user file metas error:%v\n", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(username, limit)
	if err != nil {
		log.Fatalf("query user file metas error:%v\n", err)
		return nil, err
	}

	var userFiles []UserFile
	for rows.Next() {
		userFile := UserFile{}
		err := rows.Scan(&userFile.FileHash, &userFile.FileSize,
			&userFile.FileName, &userFile.UploadAt, &userFile.LastUpdate)
		if err != nil {
			log.Fatalf("query user file metas error:%v\n", err)
			return nil, err
		}
		userFiles = append(userFiles, userFile)
	}
	return userFiles, nil
}
