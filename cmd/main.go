package main

import (
	todo "github.com/brutemors/rest-api-to-do"
	"github.com/brutemors/rest-api-to-do/pkg/handler"
	"github.com/brutemors/rest-api-to-do/pkg/repository"
	"github.com/brutemors/rest-api-to-do/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init conf: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env varibles: %s", err.Error())
	}

	db, err := repository.NewPostgreDB(repository.Config{
		Host:      viper.GetString("db.host"),
		Port:      viper.GetString("db.port"),
		Username:  viper.GetString("db.username"),
		DBName:    viper.GetString("db.dbname"),
		SSLMode:   viper.GetString("db.sslmode"),
		Passoword: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
