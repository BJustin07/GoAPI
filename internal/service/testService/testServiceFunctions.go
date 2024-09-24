package testService

import (
	"GoAPIOnECHO/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TodoRepository interface {
	GetAllTodo() ([]model.Todo, error)
	GetTodoById(id string) (model.Todo, error)
}

func TestGetAllTodo(c echo.Context, repo TodoRepository) error {
	todos, err := repo.GetAllTodo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch todos"})
	}
	if len(todos) == 0 {
		return c.JSON(http.StatusNoContent, map[string]string{"message": "Database is empty"})
	}
	return c.JSON(http.StatusOK, todos)
}

func TestGetTodoById(c echo.Context, repo TodoRepository) error {
	var todo model.Todo
	id := c.Param("id")
	todo, err := repo.GetTodoById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "ID couldn't be found"})
	}
	return c.JSON(http.StatusOK, todo)
}
