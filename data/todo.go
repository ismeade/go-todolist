// todo.go
// liyang
//
package data

import (
	"encoding/json"
	"time"
)

var (
	bucket []byte = []byte("data")
)

type todo struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Created string `json:"created"`
}

func Add(message string) error {
	t := todo{string(time.Now().Unix()), message, time.Now().Format("2006-01-02 03:04:05")}
	value, err := json.Marshal(t)
	if err != nil {
		insert(bucket, []byte(t.Id), value)
	}
	return err
}

func GetAll() {

}

func (t todo) Show() string {
	return t.Created + " " + t.Message + "\n"
}

func init() {
	initBucket(bucket)
}
