package meta

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

func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

func DeleteFile(filesha1 string) {
	// 此处只是演练,多线程的时候肯定是线程不安全的
	delete(fileMetas, filesha1)
}
