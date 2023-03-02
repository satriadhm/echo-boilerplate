package model

import (
	"errors"
)

var (
	Todos           = map[int]*Todo{}
	Seq             = 1
	ErrtodoNotFound = errors.New("todo not found")
)

type (
	Todo struct {
		Name   string `json:"name"`
		IsDone bool   `json:"isDone"`
		Id     int    `json:"id"`
	}

	CreateTodoRequest struct {
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
