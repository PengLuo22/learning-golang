package main

import (
	"log"
	"net/http"
)

// 当前目录下执行 go run main.go 启动 HTTP 服务
// curl 命令测试： curl http://127.0.0.1:50052/ping
func main() {
	http.HandleFunc("/ping", pong)
	log.Println("Starting http server ...")
	log.Fatal(http.ListenAndServe(":50052", nil))
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
