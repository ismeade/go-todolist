package storm

import (
	"encoding/json"
	"fmt"
	"go-todolist/model"
	"time"

	"github.com/asdine/storm/v3"
)

type TodoRepository struct {
	DB *storm.DB
}

func NewTodoRepository(db *storm.DB) TodoRepository {
	return TodoRepository{db}
}

func (t *TodoRepository) GetAll([]model.Todo, error) {
	panic("implement me")
}

func Create(message string) error {
	t := model.Todo{string(rune(time.Now().Unix())), message, time.Now().Format("2006-01-02 03:04:05")}
	value, err := json.Marshal(t)
	if err != nil {
		fmt.Println(value)
		//insert(bucket, []byte(t.Id), value)
	}
	return err
}
