package models

type Taskrequest struct {
	ID     int    `json:"task_id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type TaskResponse struct {
	ID     int    `json:"task_id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}
type TasksResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}
