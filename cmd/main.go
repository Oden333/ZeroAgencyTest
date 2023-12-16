package main

import (
	"ZAtest/pkg/hanlder"
	"ZAtest/pkg/repository"
	"ZAtest/pkg/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize db (%s)", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := hanlder.NewHandler(services)

	// Инициализация Fiber-сервера
	srv := fiber.New()

	handlers.InitRoutes(srv)

	//Graceful shutdown
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGTERM)

	// Запуск сервера в горутине
	go func() {
		if err := srv.Listen(":8080"); err != nil {
			// Обработка ошибки при запуске сервера
			os.Exit(1)
		}
	}()

	// Ожидание сигнала о выключении
	<-shutdownChannel

	// Остановка сервера
	if err := srv.Shutdown(); err != nil {
		// Обработка ошибки при graceful shutdown
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
