package handler

import (
	"fmt"
	dblayer "github.com/chuanshan/file-server/db"
	"github.com/chuanshan/file-server/util"
	"io/ioutil"
	"log"
	"net/http"
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
