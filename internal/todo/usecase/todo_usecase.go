package usecase

import "github.com/satriadhm-echo-boilerplate/internal/todo/repository"

type TodoUsecase interface {
	Create(todo *Todo) error
	FindById(id int) (*Todo, error)
	Update(todo *Todo) error
	Delete(id int) error
}

type todoUsecase struct {
	repo repository.TodoRepository
}

func NewTodoUsecase(repo repository.TodoRepository) TodoUsecase {
	return &todoUsecase{repo: repo}
}

func (uc *todoUsecase) Create(todo *Todo) error {
	return uc.repo.Create(todo)
}

func (uc *todoUsecase) FindById(id int) (*Todo, error) {
	return uc.repo.FindById(id)
}

func (uc *todoUsecase) Update(todo *Todo) error {
	return uc.repo.Update(todo)
}

func (uc *todoUsecase) Delete(id int) error {
	return uc.repo.Delete(id)
}
