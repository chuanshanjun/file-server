package handler

import (
	"io"
	"io/ioutil"
	"net/http"
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

	}
}
