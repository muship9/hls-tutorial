package pkg

import (
	"log"
	"net/http"
)

type Router interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Print("GET")
	default:
		// 指定メソッド以外はアクションを実行しない
		w.WriteHeader(405)
	}
}
