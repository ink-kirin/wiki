package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println(r.Form)
	_, _ = io.WriteString(w, "hello")
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Panic(err)
	}
}
