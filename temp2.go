package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandler{})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}

type myHandler struct {
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL: "+r.URL.String())
}
