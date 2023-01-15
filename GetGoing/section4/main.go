package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/getgoing/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request recived")
		fmt.Println(r.Method)
		w.Write([]byte("Hello world"))
	})
	mux.HandleFunc("/getgoing/go/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request recived")
		fmt.Println(r.Method)
		w.Write([]byte("Hello world go"))
	})
	http.ListenAndServe("localhost:3000", mux)
	fmt.Println("fin")
}
