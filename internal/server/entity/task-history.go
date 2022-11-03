package entity

import "time"

type TaskHistory struct {
	ID         int       `json:"id"`
	TaskID     int       `json:"task_id"`
	Status     int       `json:"status"`
	ChangeTime time.Time `json:"change_time"`
	Desc       string    `json:"desc"`
}
