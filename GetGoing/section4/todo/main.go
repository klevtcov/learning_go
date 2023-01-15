package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/klevtcov/views"
)

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}

			// w.WriteHeader()
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe(":3000", mux)

	fmt.Println("Fin")
}
