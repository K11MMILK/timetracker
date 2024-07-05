package main

import (
	timetracker "time-tracker"
	"time-tracker/pkg/handler"
	"time-tracker/pkg/repository"
	"time-tracker/pkg/service"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Time-Tracker
// @version 1.0
// @description API Server for Time-trackerList Application

// @host localhost:8000
// @BasePath

func main() {

	// Инициализация логгера
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(logrus.StandardLogger().Out)

	logger.Info("Starting the application")

	if err := godotenv.Load("configs/config.env"); err != nil {
		logger.WithError(err).Fatal("Error loading .env file")
	}
	logger.Info(".env file loaded")

	if err := initConfig(); err != nil {
		logger.WithError(err).Fatal("Error occurred while initializing config")
	}

	m, err := migrate.New(
		"file://migrations",
		"postgres://"+
			viper.GetString("DB_USERNAME")+":"+
			viper.GetString("DB_PASSWORD")+"@"+
			viper.GetString("DB_HOST")+":"+
			viper.GetString("DB_PORT")+"/"+
			viper.GetString("DB_DBNAME")+"?sslmode="+
			viper.GetString("DB_SSLMODE"),
	)

	if err != nil {
		logger.WithError(err).Fatal("Error occurred while creating migration instance")
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.WithError(err).Fatal("Error occurred while migrating database")
	}
	logger.Info("Database migration successful")

	// Инициализация подключения к базе данных
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		Username: viper.GetString("DB_USERNAME"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_DBNAME"),
		SSLMode:  viper.GetString("DB_SSLMODE"),
	})
	if err != nil {
		logger.WithError(err).Fatal("Error occurred while connecting to database")
	}
	logger.Info("Database connection established")

	// Инициализация репозиториев и сервисов
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	logger.Info("Repositories and services initialized")

	// Запуск сервера
	srv := new(timetracker.Server)
	port := viper.GetString("port")
	logger.Infof("Starting server on port %s", port)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		logger.WithError(err).Fatal("Error occurred while running HTTP server")
	}
	logger.Info("Server started successfully")

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		logger.WithError(err).Fatal("Error occurred while migrating database")
	}
	logger.Info("Database migration down")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	return viper.ReadInConfig()
}
