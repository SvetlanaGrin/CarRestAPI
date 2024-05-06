package main

import (
	"carRestAPI/cmd/initStart"
	"carRestAPI/internal/handler"
	"carRestAPI/internal/repository"
	"carRestAPI/internal/server"
	"carRestAPI/internal/service"

	"os"
)

//	@title			Order details API
//	@version		1.0
//	@description	API server for receiving order information

//	@host		localhost:8000
//	@BasePath	/

func main() {
	initStart.InitLogrus()
	db := initStart.InitDB()
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		return
	}
}
