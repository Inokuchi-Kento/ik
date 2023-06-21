package services

import (
	"blrpc/ent"
	"blrpc/models"
	"context"
)

type GreetServicer interface {
	GreetService()
}

type TaskServicer interface {
	GetTaskService(ctx context.Context, title string) ([]*ent.Task, error)
	CreateTaskService(ctx context.Context, task models.Task) error
}
