package storm

import (
	"encoding/json"
	"go-todolist/repository"
	"time"

	"github.com/asdine/storm/v3"
)

type todoRepository struct {
	DB *storm.DB
}

func NewTodoRepository(db *storm.DB) repository.TodoRepository {
	return &todoRepository{db}
}

func (t *todoRepository) GetAll([]model.Task, error) {
	panic("implement me")
}

func Create(message string) error {
	t := todo{string(rune(time.Now().Unix())), message, time.Now().Format("2006-01-02 03:04:05")}
	value, err := json.Marshal(t)
	if err != nil {
		insert(bucket, []byte(t.Id), value)
	}
	return err
}
