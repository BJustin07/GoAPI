package controller

import (
	"GoAPIOnECHO/internal/middleware"
	"GoAPIOnECHO/internal/service"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/todo", service.GetAllTodo)
	e.GET("/todo/:id", service.GetTodoById, middleware.JWTMiddleware())
	e.GET("/todo/token", service.Login, middleware.AuthUser())
	e.POST("/todo", service.CreateTodo, middleware.JWTMiddleware())
	e.PUT("/todo/:id", service.UpdateTodo, middleware.JWTMiddleware())
	e.DELETE("/todo/:id", service.DeleteTodo, middleware.JWTMiddleware())
}
