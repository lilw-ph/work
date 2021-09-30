/*
* @author: lilw
* @date: 2021-09-30
* @describe:
*	1.接收客户端 request，并将 request 中带的 header 写入 response header
*	2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
*	3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
*	4.当访问 localhost/healthz 时，应返回200
 */

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	listenAddress = flag.String("web.listen-address", ":8080", "server port.")
	serverURI     = flag.String("web.server-uri", "/work/001", "server uri")
)

func main() {
	flag.Parse()
	http.HandleFunc(*serverURI, httpHandler)
	http.HandleFunc("/localhost/healthz", healthHandler)
	http.ListenAndServe(*listenAddress, nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		for _, h := range v {
			w.Header().Add(k, h)
		}
	}
	// get VERSION
	version := os.Getenv("VERSION")
	w.Header().Add("Version", version)
	code := 200
	w.WriteHeader(code)
	// log IP HTTP CODE
	log.Println("IP", r.RemoteAddr, "code:", code)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
