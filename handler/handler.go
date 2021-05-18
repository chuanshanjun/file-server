package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// upload file method
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// GET 方法 获取页面内容
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {

			io.WriteString(w, "internal Server error!")
		}

		io.WriteString(w, string(data))
	}

	// POST 获取问价内容
	if r.Method == http.MethodPost {
		// index.html 使用form提交 name=file 所以此处key使用file
		file, header, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("upload file error:%s", err.Error())
		}
		defer file.Close()

		newFile, err := os.Create("./tmp/" + header.Filename)
		if err != nil {
			fmt.Printf("create file erroe:%s", err.Error())
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("save file erroe:%s", err.Error())
		}

		// http 302 暂时性转移
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload file success")
}
