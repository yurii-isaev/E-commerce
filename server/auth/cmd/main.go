package main

import (
	"e-commerce"
	"e-commerce/internal/services"
	"e-commerce/internal/usecases/rest/handlers"
	"e-commerce/repositories"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repo := repositories.NewRepository()
	service := services.NewService(repo)
	srv := new(server.Server)
	handlers := handler.NewHandler(service)

	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
