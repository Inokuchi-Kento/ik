package controllers

import (
	taskv1 "blrpc/gen/task/v1"
	"blrpc/gen/task/v1/taskv1connect"
	"blrpc/services"
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
)

type TaskController struct {
	service services.TaskServicer
}

func NewTaskController(s services.TaskServicer) *TaskController {
	return &TaskController{service: s}
}

func (c *TaskController) TaskListHandler() (string, http.Handler) {
	return taskv1connect.NewTaskServiceHandler(c)
}

func (c *TaskController) ListTasks(
	ctx context.Context, req *connect.Request[taskv1.ListTasksRequest],
) (*connect.Response[taskv1.ListTasksResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	title := req.Msg.Title
	if title == "" {
		return nil, fmt.Errorf("no request")
	}

	t, err := c.service.GetTaskService(ctx, title)
	if err != nil {
		return nil, err
	}

	var tasks []*taskv1.ListTasksResponse_Task
	for _, v := range t {
		taskRes := &taskv1.ListTasksResponse_Task{
			Id:     int32(v.ID),
			Title:  v.Title,
			Detail: v.Detail,
		}
		tasks = append(tasks, taskRes)
	}

	res := connect.NewResponse(&taskv1.ListTasksResponse{
		Tasks: tasks,
	})

	return res, nil
}
