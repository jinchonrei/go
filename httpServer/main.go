package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
)

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
		t, _ := template.ParseFiles("index.html")
		p := Persion{Name: "dipro"}
		t.Execute(w, p)
	})

	http.HandleFunc("/cmd", onCmd)

	http.ListenAndServe(":8787", nil)
}
