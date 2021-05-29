package handler

import (
	"fmt"
	dblayer "github.com/chuanshan/file-server/db"
	"github.com/chuanshan/file-server/util"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var salut = "*#890"

func SingUPHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		file, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			log.Fatalf("read signup.html failed:%v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(file)
		return
	}

	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	if len(username) < 3 || len(password) < 5 {
		w.Write([]byte("Invalid parameter"))
	}

	enc_passwd := util.Sha1([]byte(password + salut))
	suc := dblayer.UserSignUP(username, enc_passwd)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signin.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	// 1 校验密码
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	encPassword := util.Sha1([]byte(password + salut))
	checkPassword := dblayer.UserSignIn(username, encPassword)
	if !checkPassword {
		w.Write([]byte("FAILED PASSWORD!"))
		return
	}

	// 2 生成token
	token := GenToken(username)
	updateToken := dblayer.UpdateToken(username, token)
	if !updateToken {
		w.Write([]byte("FAILED"))
		return
	}

	// 3 重定向首页
	//w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
	resp := util.RespMsg{
		Code: 200,
		Msg:  "success",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			//Location: "http://" + r.Host + "/static/view/home.html",
			Location: "http://" + r.Host + "/home",
			Username: username,
			Token:    token,
		},
	}
	w.Write(resp.JSONBytes())
}

func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	// time.now().Unix() = 1621864600, %x 就是将 1621864600 用16进制表示 60abb098 正好8位
	timestamp := fmt.Sprintf("%x", time.Now().Unix())
	prefix := util.MD5([]byte(username + timestamp + "_tokensalt"))
	return prefix + timestamp[:8]
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/home.html")
		if err != nil {
			w.Write([]byte("FAILED"))
		}
		w.Write(data)
	}
}

func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// 1 解析请求参数
	r.ParseForm()

	username := r.Form.Get("username")
	token := r.Form.Get("token")
	// 2 验证token有效性
	if ValidToken(username, token) {
		w.Write([]byte("token is invalid"))
		return
	}
	// 3 获取用户信息
	user, err := dblayer.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	// 4 组装参数返回
	resp := util.RespMsg{
		Code: 200,
		Msg:  "success",
		Data: user,
	}
	w.Write(resp.JSONBytes())
}

// ValidToken : 验证token
func ValidToken(username string, token string) bool {
	// 1 验证token长度
	if len(token) != 40 {
		return false
	}

	// 2 判断token实效性
	suffix := token[len(token)-8:]
	parseInt, err := strconv.ParseInt(suffix, 16, 64)
	if err != nil {
		log.Fatalf("Valid token failed err:%v", err)
		return false
	}

	diff := time.Now().Unix() - parseInt
	overtime := diff / 3600
	if overtime > 1 {
		log.Println("token invalid")
		return false
	}

	// 3 从数据库表tbl_user_token查询username对应的token信息
	userToken, err := dblayer.GetUserToken(username)
	if err != nil {
		log.Fatalf("get user token from DB error:%v", err)
		return false
	}

	// 4 对比两个token是否一致
	if !(token == userToken.UserToken) {
		log.Println("frontEnd token not equals backEnd token")
		return false
	}
	return true
}
