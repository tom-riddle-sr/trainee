package main

import (
	mysql "trainee/fibertrainee3/database/mysql"
	"trainee/fibertrainee3/handlers"
	"trainee/fibertrainee3/repository"
	"trainee/fibertrainee3/router"
	"trainee/fibertrainee3/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysql.New()
	repo := repository.NewRepo()

	servicesAccount := services.NewAccount(repo)
	servicesAuth := services.NewAuth(repo)
	services := services.NewServices(servicesAuth, servicesAccount)

	handlersAuth := handlers.NewAuth(services)
	handlersAccount := handlers.NewAccount(services)
	handlers := handlers.NewHandlers(handlersAuth, handlersAccount)

	router.Router(handlers)
}
