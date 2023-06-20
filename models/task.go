package models

type Task struct {
	ID     int    `json:"task_id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}
