package main

import (
	"fmt"
	"net/http"
	"os"
)

// Simple counter server.
type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

// 参数服务器。
func ArgServer(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, os.Args)
}

func main() {
	ctr := Counter{}
	http.Handle("/counter", &ctr)
	http.Handle("/arg", http.HandlerFunc(ArgServer))
	http.ListenAndServe(":8080", nil)
}
