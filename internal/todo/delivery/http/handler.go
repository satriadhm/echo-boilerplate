package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/satriadhm-echo-boilerplate/internal/todo"
	"github.com/satriadhm-echo-boilerplate/internal/todo/usecase"
)

type TodoHandler struct {
	usecase usecase.TodoUsecase
}

func NewTodoHandler(e *echo.Echo, uc usecase.TodoUsecase) {
	handler := &TodoHandler{usecase: uc}
	e.POST("/todo", handler.CreateTodo)
	e.GET("/todo/:id", handler.GetTodo)
	e.PUT("/todo/:id", handler.UpdateTodo)
	e.DELETE("/todo/:id", handler.DeleteTodo)
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
	todo := new(todo.Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}

	if err := h.usecase.Create(todo); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create todo"})
	}
	return c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) GetTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	todo, err := h.usecase.FindById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Todo not found"})
	}
	return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	todo := new(todo.Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}

	if err := h.usecase.Update(todo); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update todo"})
	}
	return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	if err := h.usecase.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete todo"})
	}
	return c.NoContent(http.StatusNoContent)
}
