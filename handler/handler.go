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

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	filehash := r.Form.Get("filehash")
	fileMeta := meta.GetFileMeta(filehash)

	// 打开文件的时候,文件可能不存在,也可能打开失败
	file, err := os.Open(fileMeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("download file error%s\n", err.Error())
	}
	defer file.Close()

	// 通过拿到的文件句柄将文件读入内存(小文件)
	// 如果文件大的话，我们使用流的方式,每次读一小部分数据给客户端，然后再刷新缓存，继续读取到文件的末尾为止
	data, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("download file error%s\n", err.Error())
	}

	// 如果前段是浏览器则要加头让浏览器能识别
	// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Basics_of_HTTP/MIME_types
	// application/octet-stream 表明是某种二进制数据
	// 这是应用程序文件的默认值。意思是 未知的应用程序文件 ，浏览器一般不会自动执行或询问执行。浏览器会像对待 设置了HTTP头Content-Disposition 值为 attachment 的文件一样来对待这类文件。
	// 在常规的 HTTP 应答中，Content-Disposition 响应头指示回复的内容该以何种形式展示，是以内联的形式（即网页或者页面的一部分），还是以附件的形式下载并保存到本地。
	// attachment（意味着消息体应该被下载到本地；大多数浏览器会呈现一个“保存为”的对话框，将 filename 的值预填为下载后的文件名，假如它存在的话）
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename=\""+fileMeta.FileName+"\"")
	// 如果前段是app则直接写数据
	w.Write(data)
}
