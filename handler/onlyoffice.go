package handler

import (
	"io/ioutil"
	"log"
	"net/http"
)

func OpenFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	file, err := ioutil.ReadFile("./static/view/onlyoffice/openfile.html")
	if err != nil {
		log.Fatalf("read %s err:%v", "openfile.html", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(file)
}

func DownloadOnlyoffice(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./tmp/testword.docx")
	if err != nil {
		log.Fatalf("read file err:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(file)
}
