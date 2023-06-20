package services

import (
	"blrpc/ent"
	"blrpc/repositories"
	"context"
	"fmt"
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
