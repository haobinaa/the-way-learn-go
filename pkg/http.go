package pkg

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

type HelloHandleStruct struct {
	content string
}

// ServeHTTP 实现 http.Handler 接口
func (h *HelloHandleStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, h.content)
}

func startServer() {
	http.HandleFunc("/", index)
	http.Handle("/hello", &HelloHandleStruct{content: "Hello Hanlder"})
	http.ListenAndServe(":8080", nil)
	fmt.Println("sever started:8080")
}
