package meta

import (
	mydb "github.com/chuanshan/file-server/db"
	"log"
)

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

// init函数先于main函数自动执行
func init() {
	fileMetas = make(map[string]FileMeta)
}

func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

//UpdateFileMetaDb:新增/更新文件元信息到数据库
func UpdateFileMetaDb(fmeta FileMeta) bool {
	return mydb.OnFileUploadFinished(fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

//GetFileMetaDB从数据库中返回文件元信息
func GetFileMetaDB(fileSha1 string) (FileMeta, error) {
	tfile, err := mydb.GetFileMetaDB(fileSha1)
	if err != nil {
		log.Fatalf("get file from db error:%v", err)
		return FileMeta{}, err
	}

	fileMeta := FileMeta{
		FileSha1: tfile.FileSha1,
		FileSize: tfile.FileSize.Int64,
		FileName: tfile.FileName.String,
	}
	return fileMeta, nil
}

func DeleteFile(filesha1 string) {
	// 此处只是演练,多线程的时候肯定是线程不安全的
	delete(fileMetas, filesha1)
}
