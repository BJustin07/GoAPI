package service

import (
	"GoAPIOnECHO/internal/model"
	"GoAPIOnECHO/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

var validate = validator.New()

func GetAllTodo(c echo.Context) error {
	var todos []model.Todo
	result := repository.DB.Find(&todos)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch todos"})
	}
	if len(todos) == 0 {
		return c.JSON(http.StatusNoContent, map[string]string{"message": "Database is empty"})
	}
	return c.JSON(http.StatusOK, todos)
}

func GetTodoById(c echo.Context) error {
	var todo model.Todo
	id := c.Param("id")
	result := repository.DB.First(&todo, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "ID couldn't be found"})
	}
	return c.JSON(http.StatusOK, todo)
}

func CreateTodo(c echo.Context) error {
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid/missing data"})
	}
	if err := validate.Struct(todo); err != nil {
		errorMap := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errorMap[err.Field()] = err.Field() + " is required"
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
	}
	result := repository.DB.Create(&todo)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create todo"})
	}
	return c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c echo.Context) error {
	var todo model.Todo
	id := c.Param("id")
	result := repository.DB.First(&todo, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "ID couldn't be found"})
	}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid data"})
	}
	if err := validate.Struct(todo); err != nil {
		errorMap := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errorMap[err.Field()] = err.Field() + " is required"
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
	}
	result = repository.DB.Save(&todo)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update todo"})
	}

	return c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c echo.Context) error {
	var todo model.Todo
	id := c.Param("id")
	result := repository.DB.First(&todo, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "ID couldn't be found"})
	}
	repository.DB.Delete(&todo, id)
	return c.JSON(http.StatusOK, "Successfully deleted an accomplished task!")
}

func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
