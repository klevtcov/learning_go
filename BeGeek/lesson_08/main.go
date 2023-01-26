package main

import (
	"./db"
	logger "./log"
	"fmt"
)

func main() {

	db.HelloDB()
	db.HelloDataBase()
	logger.HelloLog()
	fmt.Println("")

}
