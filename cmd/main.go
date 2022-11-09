package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/kirill0909/neurohacking-api"
	"github.com/kirill0909/neurohacking-api/pkg/handler"
	"github.com/kirill0909/neurohacking-api/pkg/repository"
	"github.com/kirill0909/neurohacking-api/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := initDB()
	if err != nil {
		logrus.Fatalf("faild to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDB() (*sqlx.DB, error) {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
