package services

import (
	"blrpc/ent"
	"context"
)

type GreetServicer interface {
	GreetService()
}

type TaskServicer interface {
	GetTaskService(ctx context.Context, title string) ([]*ent.Task, error)
}
