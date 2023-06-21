package services

import (
	"blrpc/ent"
	"blrpc/models"
	"blrpc/repositories"
	"context"
	"fmt"
	"log"
)

func (s *AppService) GetTaskService(ctx context.Context, title string) ([]*ent.Task, error) {
	taskList, err := repositories.SelectTask(ctx, title)
	if err != nil {
		return nil, err
	}

	if len(taskList) == 0 {
		err := fmt.Errorf("no data")
		return nil, err
	}
	return taskList, nil
}

func (s *AppService) CreateTaskService(ctx context.Context, task models.Task) error {
	log.Printf("title: %s, detail: %s", task.Title, task.Detail)
	err := repositories.InsertTask(ctx, task)
	if err != nil {
		return err
	}
	return nil
}
