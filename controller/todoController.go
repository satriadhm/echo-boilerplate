package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/Satria2133/echo-CRUD/config"
	"github.com/Satria2133/echo-CRUD/model"
	"github.com/labstack/echo/v4"
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

	sql := "INSERT INTO todo(id, Name, isDone) VALUES(?, ?, ?)"
	stmt, err := config.Db.Prepare(sql)

	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	ctx := context.Background()
	result, err2 := stmt.ExecContext(ctx, todo.Id, todo.Name, todo.IsDone)
	defer stmt.Close()

	if err2 != nil {
		fmt.Print(err2.Error())
		return err2
	}
	fmt.Println(result.LastInsertId())

	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	defer stmt.Close()
	return c.JSON(http.StatusCreated, todo)
}

// GetTodoList ...
func GetTodoList(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, model.Todos[id])
}

// UpdateTodoList ...

func UpdateTodoList(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}
	if _, ok := model.Todos[todo.Id]; !ok {
		return echo.NewHTTPError(http.StatusNotFound, model.ErrtodoNotFound.Error())
	}
	model.Todos[todo.Id] = todo
	return c.JSON(http.StatusOK, todo)
}

// DeleteTodoList ...

func DeleteTodoList(c echo.Context) error {
	Lock.Lock()
	defer Lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	if _, ok := model.Todos[id]; !ok {
		return echo.NewHTTPError(http.StatusNotFound, model.ErrtodoNotFound.Error())
	}
	delete(model.Todos, id)
	return c.NoContent(http.StatusNoContent)
}
