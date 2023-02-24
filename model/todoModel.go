package model

import (
	"errors"
	"sync"
)

var (
	todos = map[int]*todo{}
	seq   = 1
	lock  = sync.Mutex{}

	ErrtodoNotFound = errors.New("todo not found")
)

type (
	todo struct {
		Name   string `json:"name"`
		IsDone bool   `json:"isDone"`
		Id     int    `json:"id"`
	}

	createTodoRequest struct {
		Name   string `json:"name"`
		IsDone bool   `json:"isDone"`
	}

	FindByIDRequest struct {
		Id int `param:"id" query:"id" form:"id" json:"id"`
	}

	UpdateTodoRequest struct {
		Id     int    `param:"id" query:"id" form:"id" json:"id"`
		Name   string `json:"name"`
		IsDone bool   `json:"isDone"`
	}
)
