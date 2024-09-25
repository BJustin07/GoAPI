package controller

import (
	"GoAPIOnECHO/internal/middleware"
	"GoAPIOnECHO/internal/model"
	"GoAPIOnECHO/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	Service   service.TodoService
	Validator *validator.Validate
}

func (ctrl *Controller) GetAllTodo(c echo.Context) error {
	todos, err := ctrl.Service.GetAllTodo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todos)
}

func (ctrl *Controller) GetTodoById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}
	todo, err := ctrl.Service.GetTodoById(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := ctrl.Validator.Struct(todo); err != nil {
		errorMap := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errorMap[err.Field()] = err.Field() + " is required"
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
	}
	return c.JSON(http.StatusOK, todo)
}

func (ctrl *Controller) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func (ctrl *Controller) CreateTodo(c echo.Context) error {
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Something went wrong"})
	}
	if err := ctrl.Validator.Struct(todo); err != nil {
		errorMap := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errorMap[err.Field()] = err.Field() + " is required"
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
	}
	if err := ctrl.Service.CreateTodo(&todo); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Successfully created todo!")
}

func (ctrl *Controller) UpdateTodo(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}
	var updatedTodo model.Todo
	if err := c.Bind(&updatedTodo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request data"})
	}
	if err := ctrl.Validator.Struct(updatedTodo); err != nil {
		errorMap := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errorMap[err.Field()] = err.Field() + " is required"
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
	}
	todo, err := ctrl.Service.UpdateTodo(uint(id), updatedTodo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

func (ctrl *Controller) DeleteTodo(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}
	_, err = ctrl.Service.DeleteTodo(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "You accomplished another task, well done!")
}

func SetupRoutes(e *echo.Echo, ctrl *Controller) {
	e.GET("/todo", ctrl.GetAllTodo)
	e.GET("/todo/:id", ctrl.GetTodoById, middleware.JWTMiddleware())
	e.GET("/todo/token", ctrl.Login, middleware.AuthUser())
	e.POST("/todo", ctrl.CreateTodo, middleware.JWTMiddleware())
	e.PUT("/todo/:id", ctrl.UpdateTodo, middleware.JWTMiddleware())
	e.DELETE("/todo/:id", ctrl.DeleteTodo, middleware.JWTMiddleware())
}
