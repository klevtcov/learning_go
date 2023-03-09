package main

import (
	"log"

	"github.com/klevtcov/todo"
	"github.com/klevtcov/todo/pkg/handler"
	"github.com/klevtcov/todo/pkg/repository"
	"github.com/klevtcov/todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
