package handler

import (
	"fmt"
	"github.com/chuanshan/file-server/cache/redis"
	"github.com/chuanshan/file-server/util"
	"log"
	"math"
	"net/http"
	"strconv"
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

	// 5.将响应初始化信息返回客户端
	w.Write(util.NewRespMsg(0, "OK", upInfo).JSONBytes())
}
