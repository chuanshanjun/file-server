package handler

import (
	"encoding/json"
	"fmt"
	"github.com/chuanshan/file-server/meta"
	"github.com/chuanshan/file-server/util"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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

		filemeta := meta.FileMeta{
			FileName: header.Filename,
			Location: "./tmp/" + header.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile, err := os.Create(filemeta.Location)
		if err != nil {
			fmt.Printf("create file erroe:%s", err.Error())
		}
		defer newFile.Close()

		// 巧妙的使用io.Copy返回的字节数给到对象中
		filemeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("save file erroe:%s", err.Error())
		}

		// 上面有io.Copy() 所以将下次文件读的offset设置为0
		newFile.Seek(0, 0)
		filemeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(filemeta)
		// sha1操作中有io.Copy() 所以将下次文件读的offset设置为0
		newFile.Seek(0, 0)
		if err != nil {
			fmt.Printf("read file to []byte error:%s", err.Error())
		}

		// http 302 暂时性转移
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload file success")
}

func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	// 将参数解析完后放入r.Form
	r.ParseForm()
	// map[string][]string map key 为string 值为 string数组 所以 取[0] 拿第一个值
	// value =  r.Form["filehash"][0]
	// 个人觉得r.FormValue() 更好用
	value := r.FormValue("filehash")
	fileMeta := meta.GetFileMeta(value)
	marshal, err := json.Marshal(fileMeta)
	if err != nil {
		fmt.Printf("json.Marshal error:%s\n", err.Error())
	}

	w.Write(marshal)
}
