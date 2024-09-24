package testService

import (
	"GoAPIOnECHO/internal/model"
	"GoAPIOnECHO/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllTodo_Success(t *testing.T) {
	e := echo.New()
	mockRepo := new(repository.MockRepository)
	mockRepo.On("GetAllTodo").Return([]model.Todo{{ID: 1, Description: "Test Todo", Done: false}}, nil)

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := TestGetAllTodo(c, mockRepo)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `[{"ID":1,"Description":"Test Todo","Done":false}]`, rec.Body.String())

	mockRepo.AssertExpectations(t)
}

func TestGetTodoById_Success(t *testing.T) {
	e := echo.New()
	mockRepo := new(repository.MockRepository)
	mockRepo.On("GetTodoById", "1").Return(model.Todo{ID: 1, Description: "Test Todo", Done: false}, nil)

	req := httptest.NewRequest(http.MethodGet, "/todo/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("1")
	err := TestGetTodoById(c, mockRepo)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"ID":1,"Description":"Test Todo","Done":false}`, rec.Body.String())

	mockRepo.AssertExpectations(t)
}
