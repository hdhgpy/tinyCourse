package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type HTTPProxy struct {
	proxy *httputil.ReverseProxy
}

func NewHttpProxy(target string) (*HTTPProxy, error) {
	u, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	return &HTTPProxy{httputil.NewSingleHostReverseProxy(u)}, nil
}

func (h *HTTPProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.proxy.ServeHTTP(w, r)
	n, err := fmt.Fprintln(w, "Hello world!!!")
	if err != nil {
		return
	}
	fmt.Println(n)
}

func main() {
	proxy, err := NewHttpProxy("http//:127.0.0.1:8888")
	if err != nil {

	}
	http.Handle("/", proxy)

	http.ListenAndServe(":8081", nil)
}
