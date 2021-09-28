package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// 单独写回调函数
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthzHandler)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// root 函数
func rootHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		log.Printf("client_addr:%s, response_code:%d, method:%s,path:%s", r.RemoteAddr, 404, r.Method, r.URL.Path)
		w.WriteHeader(404)
		w.Write([]byte("The Page Not Found"))
		return
	}
	// 请求头添加到响应头
	for header_k, header_v := range r.Header {
		// fmt.Println(header_k, header_v, header_v[0])
		// 添加响应头信息
		w.Header().Set(header_k, header_v[0])
	}

	// go version添加到响应头
	w.Header().Set("Go-Version", os.Getenv("VERSION"))

	// 添加响应状态码
	w.WriteHeader(200)

	// 写入数据到响应实体
	w.Write([]byte("hello,world"))

	// 请求日志打印在标准输出
	log.Printf("client_addr:%s, response_code:%d, method:%s,path:%s", r.RemoteAddr, 200, r.Method, r.URL.Path)

}

// 健康检查
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	log.Printf("client_addr:%s, response_code:%d, method:%s,path:%s", r.RemoteAddr, 200, r.Method, r.URL.Path)
	fmt.Fprintln(w, "ok")
}
