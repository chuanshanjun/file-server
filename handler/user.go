package handler

import (
	dblayer "github.com/chuanshan/file-server/db"
	"github.com/chuanshan/file-server/util"
	"io/ioutil"
	"log"
	"net/http"
)

var salut = "#*778"

func SingUP(w http.ResponseWriter, r *http.Request) {
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
