package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	_ "net/http/pprof"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/monkeysungit/mygo/httpserver/metrics"
)

func main() {
	defer glog.Flush()
	glog.V(2).Info("Starting http server...")
	metrics.Register()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/metrics", promttp.Handler())
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	io.WriteString(w, "ok\n")
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-REAL-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	fmt.Println("entering root handler " + ip)
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	delay := randInt(10, 2000)
	time.Sleep(time.Millisecond * time.Duration(delay))
	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! clientip: %s", ip)
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
