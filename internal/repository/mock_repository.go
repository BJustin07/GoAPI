package repository

import (
	"GoAPIOnECHO/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAllTodo() ([]model.Todo, error) {
	args := m.Called()
	return args.Get(0).([]model.Todo), args.Error(1)
}

func (m *MockRepository) GetTodoById(id string) (model.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(model.Todo), args.Error(1)
}
