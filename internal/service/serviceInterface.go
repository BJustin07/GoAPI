package service

import "GoAPIOnECHO/internal/model"

type Service interface {
	GetAllTodo() ([]model.Todo, error)
	GetTodoById(id uint) (model.Todo, error)
	CreateTodo(todo *model.Todo) error
	UpdateTodo(id uint, updatedTodo model.Todo) (model.Todo, error)
	DeleteTodo(id uint) (model.Todo, error)
}
