package repositories

import (
	"blrpc/ent"
	"blrpc/ent/task"
	"blrpc/models"
	"context"
	"fmt"
)

func SelectTask(ctx context.Context, title string) ([]*ent.Task, error) {
	t, err := client.Task.Query().Where(task.Title(title)).All(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func InsertTask(ctx context.Context, task models.Task) error {
	t, err := client.Task.Create().SetTitle("").SetDetail("").Save(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("created schema %s", t)
	return nil
}
