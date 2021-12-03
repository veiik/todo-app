package main

import (
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	serv := service.NewService(repos)
	handl := handler.NewHandler(serv)
	srv := new(todo.Server)
	if err := srv.Run("8080", handl.InitRoutes()); err != nil {
		panic(err)
	}
}
