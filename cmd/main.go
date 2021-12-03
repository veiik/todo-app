package main

import (
	"github.com/spf13/viper"
	"log"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Printf("config file error: %#v", err)
	}
	repos := repository.NewRepository()
	serv := service.NewService(repos)
	handl := handler.NewHandler(serv)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handl.InitRoutes()); err != nil {
		panic(err)
	}
}

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
