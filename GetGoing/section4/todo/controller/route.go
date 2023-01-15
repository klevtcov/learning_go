package controller

import (
	"encoding/json"
	"net/http"

	"github.com/klevtcov/views"
)

func Register() *http.ServeMux {
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
	return mux
}
