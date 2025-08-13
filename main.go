package main

import (
	"log"
	"projeto-metas/database"
	"projeto-metas/handlers"
	"projeto-metas/repository"
	"projeto-metas/routes"
	"projeto-metas/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	database.Conectar()

	v := validator.New()

	UserRepository := repository.NewUserRepository(database.DB)
	UserService := services.NewUserService(UserRepository)
	UserHandler := handlers.NewUserHandler(UserService)

	MetaRepository := repository.NewMetaRepository(database.DB)
	MetaService := services.NewMetaService(MetaRepository)
	MetaHandler := handlers.NewMetaHandler(MetaService, v)

	routes.RegistrarRequests(e, UserHandler, MetaHandler)

	log.Fatal(e.Start(":8080"))

}
