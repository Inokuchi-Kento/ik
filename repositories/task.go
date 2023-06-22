package repositories

import (
	"blrpc/ent"
	"blrpc/ent/task"
	"blrpc/models"
	"context"
)

func SelectTask(ctx context.Context, title string) ([]*ent.Task, error) {
	t, err := client.Task.Query().Where(task.Title(title)).All(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func InsertTask(ctx context.Context, task models.Task) error {
	_, err := client.Task.Create().SetTitle(task.Title).SetDetail(task.Detail).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTask(ctx context.Context, title string, newtask models.Task) error {
	_, err := client.Task.
		Update().
		Where(task.Title(title)).
		SetTitle(newtask.Title).
		SetDetail(newtask.Detail).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil

}

func DeleteTask(ctx context.Context, title string) error {
	_, err := client.Task.Delete().Where(task.Title(title)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
