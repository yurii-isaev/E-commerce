package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"

	"e-commerce/internal/controllers"
	"e-commerce/internal/storage/configs"
	"e-commerce/internal/storage/repositories"
	"e-commerce/internal/usecases/rest/services"

	"e-commerce"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	svr, handlers := initDB()
	if err := svr.Run(viper.GetString("port"), handlers.InitHTTPEndpoints()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDB() (*server.Server, *controllers.Router) {
	db, err := configs.NewPostgresDB(configs.ConfigDB{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repositories.NewRepository(db)
	service := services.NewService(repo)
	srv := new(server.Server)
	handlers := controllers.NewHandler(service)

	return srv, handlers
}
