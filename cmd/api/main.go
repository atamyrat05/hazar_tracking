package main

import (
	handler "hazar_tracking/internal/delivery/http"
	"hazar_tracking/internal/repository"
	"hazar_tracking/internal/service"
	"hazar_tracking/pkg/database"
	"hazar_tracking/pkg/server"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// knhb pvzi lawk iwqq
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	//gin.SetMode(gin.ReleaseMode)

	if err := InitConfig(); err != nil {
		logrus.Fatalf("Can not read config files,  %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {

		logrus.Fatalf("Failed loading env variables,  %s", err.Error())
	}
	db, err := database.ConnectDB(database.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Sslmode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("Can not connect to DB %s ", err.Error())
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	srv := new(server.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		logrus.Fatalf("Server not running...!, %s", err)
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
