package entity

import "time"

type Task struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Type       int       `json:"type"`
	Status     int       `json:"status"`
	Desc       string    `json:"desc,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time"`
}
