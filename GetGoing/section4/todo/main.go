package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/klevtcov/controller"
	"github.com/klevtcov/model"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
