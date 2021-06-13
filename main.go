package main

import (
	"fmt"
	"github.com/chuanshan/file-server/handler"
	"net/http"
)

// http://hu12340.vip/2020/10/GO-%E6%93%8D%E4%BD%9C%E6%96%87%E4%BB%B6/ go 操作文件
// https://gobyexample-cn.github.io/
// https://golang.hotexamples.com/zh/ golang例子网
// main method
// 文件上传 POST /file/upload
// 文件查询 GET  /file/query
// 文件下载 GET  /file/download
// 文件删除 POST /file/delete
// 文件修改(重命名) POST /file/update
func main() {
	// 设定路由规则
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.DeleteFile)
	http.HandleFunc("/user/signup", handler.SingUPHandler)
	http.HandleFunc("/user/signin", handler.SignInHandler)
	http.HandleFunc("/home", handler.HomeHandler)
	http.HandleFunc("/user/info", handler.HttpIntercept(handler.UserInfoHandler))
	http.HandleFunc("/file/query", handler.FileQueryHandler)
	http.HandleFunc("/file/fastupload", handler.TryFastUploadHandler)
	http.HandleFunc("/file/mpupload/init", handler.InitialMultipartUploadHandler)
	http.HandleFunc("/file/word", handler.WordHandler)
	http.HandleFunc("/onlyoffice/openfile", handler.OpenFile)
	http.HandleFunc("/onlyoffice/download/testword.docx", handler.DownloadOnlyoffice)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("failed to start server, err:%s", err.Error())
	}
}
