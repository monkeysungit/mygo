package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "net/http/pprof"

	"github.com/golang/glog"
)

func main() {
	defer glog.Flush()
	glog.V(2).Info("Starting http server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	io.WriteString(w, "ok\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-REAL-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	fmt.Println("entering root handler " + ip)
	glog.V(2).Infof("Client IP [%s], returnCode [%s]", ip, http.StatusOK)
	w.WriteHeader(200)
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	os.Setenv("VERSION", "1.0")
	VERSION := os.Getenv("VERSION")
	io.WriteString(w, fmt.Sprintf("系统环境变量：[version][%s]\n", VERSION))
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}
