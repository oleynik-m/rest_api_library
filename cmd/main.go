package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"rest_api"
	"rest_api/pkg/handler"
	repo "rest_api/pkg/repository"
	"rest_api/pkg/service"

)






func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repo.SetupPostgreConnection(repo.Config{
		viper.GetString("dbHost"),
		viper.GetString("dbUser"),
		viper.GetString("dbPass"),
		viper.GetString("dbName"),
				viper.GetString("dbPort"),
	}) // инициализация БД

	if err != nil {
		log.Fatalf("error initializing database: %s", err.Error())
	}

	repos := repo.NewRepo(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	r := handlers.InitRoutes()






	srv := new(rest_api.Server)
	if err := srv.Run(viper.GetString("port"),r); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}




	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	log.Println("RestApiServer Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}

	if err := repo.CloseDatabaseConnection(db); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}



}

func initConfig () error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
