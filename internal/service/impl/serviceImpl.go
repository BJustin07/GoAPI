package impl

import (
	"GoAPIOnECHO/internal/model"
	"GoAPIOnECHO/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TodoService struct {
	Todo model.Todo
}

func (s *TodoService) GetAllTodo() ([]model.Todo, error) {
	var todos []model.Todo

	result := repository.DB.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

func (s *TodoService) GetTodoById(id uint) (model.Todo, error) {
	var todo model.Todo
	result := repository.DB.First(&todo, id)
	if result.Error != nil {
		return model.Todo{}, result.Error
	}
	return todo, nil
}

func (s *TodoService) CreateTodo(todo *model.Todo) error {
	result := repository.DB.Create(&todo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *TodoService) UpdateTodo(id uint, updatedTodo model.Todo) (model.Todo, error) {
	var todo model.Todo
	result := repository.DB.First(&todo, id)

	if result.Error != nil {
		return model.Todo{}, result.Error
	}
	todo.Description = updatedTodo.Description
	todo.Done = updatedTodo.Done
	saveResult := repository.DB.Save(&todo)
	if saveResult.Error != nil {
		return model.Todo{}, saveResult.Error
	}
	return todo, nil
}

func (s *TodoService) DeleteTodo(id uint) (model.Todo, error) {
	var todo model.Todo
	result := repository.DB.First(&todo, id)
	if result.Error != nil {
		return model.Todo{}, result.Error
	}
	repository.DB.Delete(&todo, id)
	return model.Todo{}, nil
}

func (s *TodoService) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
