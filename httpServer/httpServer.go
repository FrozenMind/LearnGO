package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handlePing)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method + ":" + r.Host + "" + r.RequestURI + " from " + r.RemoteAddr)
	io.WriteString(w, "im here")
}
