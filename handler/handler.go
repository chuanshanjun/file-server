package handler

import (
	"encoding/json"
	"fmt"
	dblayer "github.com/chuanshan/file-server/db"
	"github.com/chuanshan/file-server/meta"
	"github.com/chuanshan/file-server/util"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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
		return
	}

	// POST 获取文件内容
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
		//meta.UpdateFileMeta(filemeta)
		_ = meta.UpdateFileMetaDb(filemeta)
		// sha1操作中有io.Copy() 所以将下次文件读的offset设置为0
		newFile.Seek(0, 0)
		if err != nil {
			fmt.Printf("read file to []byte error:%s", err.Error())
		}

		//更新用户文件记录表
		r.ParseForm()
		username := r.Form.Get("username")
		dblayer.OnUserFileUpLoadFinished(filemeta.FileSha1, filemeta.FileSize, filemeta.FileName, username)

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
	//fileMeta := meta.GetFileMeta(value)
	fileMeta, err := meta.GetFileMetaDB(value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	marshal, err := json.Marshal(fileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
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
		return
	}
	defer file.Close()

	// 通过拿到的文件句柄将文件读入内存(小文件)
	// 如果文件大的话，我们使用流的方式,每次读一小部分数据给客户端，然后再刷新缓存，继续读取到文件的末尾为止
	data, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
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

func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	r.ParseForm()

	op := r.Form.Get("op")
	if op != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	filehash := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	fileMeta := meta.GetFileMeta(filehash)
	fileMeta.FileName = newFileName
	meta.UpdateFileMeta(fileMeta)

	marshal, err := json.Marshal(fileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	filehash := r.Form.Get("filehash")
	fileMeta := meta.GetFileMeta(filehash)
	// 1 删除元信息
	meta.DeleteFile(filehash)

	// 2 删除文件 系统删除可能会失败,此处暂时不考虑
	os.Remove(fileMeta.Location)

	w.WriteHeader(http.StatusOK)
}

// 查询用户文件
func FileQueryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	limit, _ := strconv.Atoi(r.Form.Get("limit"))

	files, err := dblayer.QueryUserFileMetas(username, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	resp := util.RespMsg{
		Code: 200,
		Msg:  "success",
		Data: files,
	}
	w.Write(resp.JSONBytes())
}

// 尝试秒传
func TryFastUploadHandler(w http.ResponseWriter, r *http.Request) {
	// 1 解析参数
	r.ParseForm()

	filehash := r.Form.Get("filehash")
	username := r.Form.Get("username")
	filesize, _ := strconv.Atoi(r.Form.Get("filesize"))
	filename := r.Form.Get("filename")

	// 2 从文件表中获取文件
	file, err := dblayer.GetFileMetaDB(filehash)
	if err != nil {
		log.Fatalf("get file from db with filehash err:%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3 如果文件不存在则直接返回
	if file == nil {
		resp := util.RespMsg{
			Code: -1,
			Msg:  "妙传失败，请访问普通上传接口",
		}
		w.Write(resp.JSONBytes())
		return
	}

	// 4 如果文件存在则只写用户自己的表
	dblayer.OnUserFileUpLoadFinished(filehash, int64(filesize), filename, username)
}
