package controllers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/Satria2133/echo-CRUD/model"
	"github.com/labstack/echo"
)

var Lock = sync.Mutex{}

// NewTodoList ...
func NewTodoList(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()
	todo := &model.Todo{
		Id: model.Seq,
	}
	if err := c.Bind(todo); err != nil {
		return err
	}
	model.Todos[todo.Id] = todo
	model.Seq++
	return c.JSON(http.StatusCreated, todo)
}

// GetTodoList ...
func GetTodoList(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, model.Todos[id])
}
