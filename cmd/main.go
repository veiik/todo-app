package main

import (
	"todo"
	"todo/pkg/handler"
)

func main() {
	handl := handler.Handler{}
	srv := new(todo.Server)
	if err := srv.Run("8080", handl.InitRoutes()); err != nil {
		panic(err)
	}
}
