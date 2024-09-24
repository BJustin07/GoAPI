package service

import "github.com/labstack/echo/v4"

type Service interface {
	GetAllTodo(c echo.Context) error
	GetTodoById(c echo.Context) error
	CreateTodo(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
	Login(c echo.Context) error
}
