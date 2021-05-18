package main

import (
	"fmt"
	"net/http"

	"github.com/chuanshan/file-server/handler"
)

// main method
// 文件上传 POST /file/upload
// 文件查询 GET  /file/query
// 文件下载 GET  /file/download
// 文件删除 POST /file/delete
// 文件修改(重命名) POST /file/update
func main() {
	// 设定路由规则
	http.HandleFunc("/file/upload", handler.UploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("failed to start server, err:%s", err.Error())
	}
}
