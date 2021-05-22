package db

import (
	"database/sql"
	mydb "github.com/chuanshan/file-server/db/mysql"
	"log"
)

func OnFileUploadFinished(filehash string, filename string, filesize int64, fileaddr string) bool {
	//INSERT INGORE语句，则会忽略导致错误的行，并将其余行插入到表中
	stmt, err := mydb.DBConn().Prepare("insert ignore into " +
		"tbl_file(file_sha1, file_name, file_size, file_addr, status) values(?,?,?,?,1)")
	if err != nil {
		log.Fatal("failed to prepare statement", err)
		return false
	}
	defer stmt.Close()

	exec, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		return false
	}

	if rf, err := exec.RowsAffected(); nil == err {
		if rf <= 0 {
			log.Fatal("File with hash has benn uploaded before", filehash)
		}
		return true
	}
	return false
}

type TableFile struct {
	FileSha1 string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

//GetFileMetaDB:从mysql获取文件元信息
func GetFileMetaDB(filehash string) (*TableFile, error) {
	stmt, err := mydb.DBConn().Prepare("select file_sha1, file_name, file_size, file_addr " +
		"from tbl_file where file_sha1 = ? and status = 1 limit 1;")
	if err != nil {
		log.Fatalf("get file from db error:%v", err)
		return nil, err
	}
	defer stmt.Close()

	tfile := TableFile{}
	err = stmt.QueryRow(filehash).Scan(&tfile.FileSha1, &tfile.FileName, &tfile.FileSize, &tfile.FileAddr)
	if err != nil {
		log.Fatalf("get file from db error:%v", err)
		return &tfile, err
	}
	return &tfile, nil
}
