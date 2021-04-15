package repository

import "go-todolist/model"

type TodoRepository interface {
	GetAll() ([]model.Todo, error)
	GetByID(ID string) (model.Todo, error)
	Create(title, details, UUID string, dueDate int64) (model.Todo, error)
	Update(t *model.Todo) error
	UpdateField(t *model.Todo, field string, value interface{}) error
	Delete(t *model.Todo) error
}
