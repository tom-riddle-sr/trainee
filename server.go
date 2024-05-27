package main

import (
	mongoDB "trainee/fibertrainee3/database/mongo"
	mysql "trainee/fibertrainee3/database/mysql"
	"trainee/fibertrainee3/database/redis"
	"trainee/fibertrainee3/handlers"
	"trainee/fibertrainee3/repository"
	"trainee/fibertrainee3/router"
	"trainee/fibertrainee3/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysql.New()
	redis.New()
	mongoDB.New()

	redis := repository.NewRedisRepo()
	sql := repository.NewSqlRepo()
	mongo := repository.NewMongoRepo()

	repo := repository.NewRepo(redis, sql, mongo)

	servicesAccount := services.NewAccount(repo)
	servicesAuth := services.NewAuth(repo)
	servicesMongo := services.NewServicesMongo(repo)
	servicesRedis := services.NewServicesRedis(repo)
	services := services.NewServices(servicesAuth, servicesAccount, servicesMongo, servicesRedis)

	handlersAuth := handlers.NewAuth(services)
	handlersAccount := handlers.NewAccount(services)
	handlersRedis := handlers.NewHandlersRedis(services)
	handlersMongo := handlers.NewHandlersMongo(services)
	handlers := handlers.NewHandlers(handlersAuth, handlersAccount, handlersMongo, handlersRedis)

	router.Router(handlers)
}
