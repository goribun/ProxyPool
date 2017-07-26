package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/goribun/ProxyPool/storage"
	"github.com/goribun/ProxyPool/util"
)

//程序版本
const VERSION = "/v1"

// 提供api服务
func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc(VERSION+"/ip", ProxyHandler)
	log.Println("Starting server", util.NewConfig().Host)
	http.ListenAndServe(util.NewConfig().Host, mux)
}

//代理处理器
func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		b, err := json.Marshal(storage.ProxyGet())
		if err != nil {
			return
		}
		w.Write(b)
	}
}
