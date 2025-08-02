package main

import (
	"log"
	_ "user-auth-service/docs"
	"user-auth-service/internal/handlers"
	"user-auth-service/internal/repository"
	api "user-auth-service/internal/server"
	"user-auth-service/internal/service"
)

func main() {

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handlers.NewHandler(service)

	srv := new(api.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured  while running http server: %s", err.Error())
	}
}
