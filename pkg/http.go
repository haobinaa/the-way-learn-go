package pkg

import (
	"fmt"
	"net/http"
	"runtime/debug"
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

func PanicRecover(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(string(debug.Stack()))
			}
		}()

		handler.ServeHTTP(w, r)
	})
}

type Middleware func(http.Handler) http.Handler

func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}

type MyMux struct {
	*http.ServeMux
	middlewares []Middleware
}

func NewMyMux() *MyMux {
	return &MyMux{
		ServeMux: http.NewServeMux(),
	}
}

func (m *MyMux) Use(middlewares ...Middleware) {
	m.middlewares = append(m.middlewares, middlewares...)
}

func (m *MyMux) Handle(pattern string, handler http.Handler) {
	handler = applyMiddlewares(handler, m.middlewares...)
	m.ServeMux.Handle(pattern, handler)
}

func (m *MyMux) HandleFunc(pattern string, handler http.HandlerFunc) {
	newHandler := applyMiddlewares(handler, m.middlewares...)
	m.ServeMux.Handle(pattern, newHandler)
}

func startServer() {
	// handleFunc 函数签名
	http.HandleFunc("/", index)
	// handle 实现 http.Handler 接口
	http.Handle("/hello", &HelloHandleStruct{content: "Hello Hanlder"})
	http.ListenAndServe(":8080", nil)
}
