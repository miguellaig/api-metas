package routes

import (
	"projeto-metas/handlers"
	"projeto-metas/middleware"

	"github.com/labstack/echo/v4"
)

func RegistrarRequests(e *echo.Echo, userHandler *handlers.UserHandler, metaHandler *handlers.MetaHandler) {
	e.POST("/registro", userHandler.RegistroUserHandler)
	e.POST("/login", userHandler.LoginUserHandler)

	r := e.Group("/meta")
	r.Use(middleware.AuthMiddleware)
	r.POST("", metaHandler.RegistrarMeta)
	r.GET("", metaHandler.ListarMetas)
	r.GET("/:id", metaHandler.ListarMetaPorID)
	r.PUT("/:id", metaHandler.UpdateMeta)
	r.DELETE("/:id", metaHandler.DeleteMeta)
}
