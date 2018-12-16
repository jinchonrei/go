package main

import (
	"fmt"
	"io"
	"log"
	"myTemplate"
	"net/http"
	"net/url"
	"os/exec"
)

type Option struct {
	Value, Id, Text string
	Selected        bool
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("content-Type", "text/plain")
	(*w).WriteHeader(http.StatusOK)
}

func onCmd(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Println(r.URL)
	fmt.Println(r.RequestURI)
	fmt.Println(r.RemoteAddr)
	io.WriteString(w, "这是从后台发送的数据")

	u, e := url.Parse(r.RequestURI)
	if e != nil {
		log.Fatal(e)
	}
	q := u.Query()
	fmt.Println(q["cmd"][0])
}

func onSetPath(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	u, e := url.Parse(r.RequestURI)
	if e != nil {
		log.Fatal(e)
	}
	q := u.Query()
	fmt.Println(q["path"][0], q["value"][0])
	cmd := exec.Command("Setx", q["path"][0], q["value"][0])
	//fmt.Println(cmd)
	cmd.CombinedOutput()
}

type Persion struct {
	Name string
}

func main() {
	fmt.Println("start http server and listen 8787 port")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, "hello world 2\n")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, myTemplate.GetTemplate())
	})

	http.HandleFunc("/cmd", onCmd)

	http.HandleFunc("/setPath", onSetPath)

	http.ListenAndServe(":8787", nil)
}
