package main

import (
	"net/http"

	"github.com/klevtcov/controller"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe(":3000", mux)
}
