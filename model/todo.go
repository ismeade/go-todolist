package model

var (
	bucket []byte = []byte("data")
)

type Todo struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Created string `json:"created"`
}
