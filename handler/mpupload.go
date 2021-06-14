package handler

import (
	"fmt"
	"github.com/chuanshan/file-server/cache/redis"
	dblayer "github.com/chuanshan/file-server/db"
	"github.com/chuanshan/file-server/util"
	redis2 "github.com/garyburd/redigo/redis"
	"log"
	"math"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// MultipartUploadInfo : 分块初始化信息
type MultipartUploadInfo struct {
	FileSize   int64
	FileHash   string
	Chunk      int
	ChunkCount int
	// 每次上传的标志
	UploadId string
}

// InitialMultipartUploadHandler : 初始化分块上传
func InitialMultipartUploadHandler(w http.ResponseWriter, r *http.Request) {
	// 1.解析用户参数
	r.ParseForm()

	username := r.Form.Get("username")
	filehash := r.Form.Get("filehash")
	filesize, err := strconv.Atoi(r.Form.Get("filesize"))
	if err != nil {
		log.Fatalf("parse parmar err:%v\n", err)
		w.Write(util.NewRespMsg(-1, "param error", nil).JSONBytes())
		return
	}

	// 2.获得redis连接
	rConn := redis.RedisPool().Get()
	defer rConn.Close()

	// 3.生成分块上传的的初始化信息
	upInfo := MultipartUploadInfo{
		FileSize:   int64(filesize),
		FileHash:   filehash,
		UploadId:   username + fmt.Sprintf("%x", time.Now().UnixNano()),
		Chunk:      5 * 1024 * 1024, // 5MB
		ChunkCount: int(math.Ceil(float64(filesize) / (5 * 1024 * 1024))),
	}

	// 4.将初始化信息写到redis
	rConn.Do("HSET", "MP_"+upInfo.UploadId, "chunkcount", upInfo.ChunkCount)
	rConn.Do("HSET", "MP_"+upInfo.UploadId, "filehash", upInfo.FileHash)
	rConn.Do("HSET", "MP_"+upInfo.UploadId, "filesize", upInfo.FileSize)

	chunkCountR, _ := rConn.Do("HGET", "MP_"+upInfo.UploadId, "chunkcount", upInfo.ChunkCount)
	filehashR, _ := rConn.Do("HGET", "MP_"+upInfo.UploadId, "filehash", upInfo.FileHash)
	fileziser, _ := rConn.Do("HGET", "MP_"+upInfo.UploadId, "filesize", upInfo.FileSize)

	fmt.Printf("chunkCount:%d filehashR:%s fileziser:%d", chunkCountR, filehashR, fileziser)

	// 5.将响应初始化信息返回客户端
	w.Write(util.NewRespMsg(0, "OK", upInfo).JSONBytes())
}

// UploadPartHandler 上传分块文件
func UploadPartHandler(w http.ResponseWriter, r *http.Request) {
	// 1 解析文件
	r.ParseForm()

	uploadId := r.Form.Get("uploadid")
	chunkIndex := r.Form.Get("index")

	// 2 获得redis连接
	rConn := redis.RedisPool().Get()
	defer rConn.Close()

	// 3 获得文件句柄，用于存储分块内容
	p := filepath.Join("data", uploadId, chunkIndex)
	os.MkdirAll(path.Dir(p), 0644)
	fd, err := os.Create(p)
	if err != nil {
		w.Write(util.NewRespMsg(-1, "Upload part failed", nil).JSONBytes())
		return
	}
	defer fd.Close()

	// 循环写 每次写1M
	buf := make([]byte, 1024*1024)
	for {
		n, err := r.Body.Read(buf)
		fd.Write(buf[:n])
		if err != nil {
			break
		}
	}

	// 4 将分块信息存到redis
	rConn.Do("HSET", "MP_"+uploadId, "chkidx_"+chunkIndex, 1)

	// 4 给响应
	w.Write(util.NewRespMsg(0, "ok", nil).JSONBytes())
}

// CompleteUploadHandler:通知上传合并
func CompleteUploadHandler(w http.ResponseWriter, r *http.Request) {
	// 1 解析请求参数
	r.ParseForm()

	username := r.Form.Get("username")
	uploadid := r.Form.Get("uploadid")
	filename := r.Form.Get("filename")
	filehash := r.Form.Get("filehash")
	filesize := r.Form.Get("filesize")

	// 2 获得redis连接池中的一个连接
	rConn := redis.RedisPool().Get()
	defer rConn.Close()

	// 3 通过uploadid查询redis并判断是否所有分块上传完成
	data, err := redis2.Values(rConn.Do("HGET", "MP_"+uploadid))
	if err != nil {
		w.Write(util.NewRespMsg(-1, "complete upload failed", nil).JSONBytes())
	}

	totalCounts := 0
	chunkCounts := 0
	for i := 0; i < len(data); i += 2 {
		k := string(data[i].([]byte))
		v := string(data[i+1].([]byte))

		if k == "chunkcount" {
			totalCounts, _ = strconv.Atoi(v)
		} else if strings.HasPrefix(k, "chkidx_") && v == "1" {
			chunkCounts++
		}
	}

	if totalCounts != chunkCounts {
		w.Write(util.NewRespMsg(-2, "invalid request", nil).JSONBytes())
	}

	// 4 TODO: 合并分块
	// 5 更新唯一文件表及用户文件表
	fsize, _ := strconv.Atoi(filesize)
	dblayer.OnFileUploadFinished(filehash, filename, int64(fsize), "")
	dblayer.OnUserFileUpLoadFinished(filehash, int64(fsize), filename, username)

	w.Write(util.NewRespMsg(0, "ok", nil).JSONBytes())
}
