package controllers

import (
	taskv1 "blrpc/gen/task/v1"
	"blrpc/gen/task/v1/taskv1connect"
	"blrpc/models"
	"blrpc/services"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
)

type TaskController struct {
	service services.TaskServicer
}

func NewTaskController(s services.TaskServicer) *TaskController {
	return &TaskController{service: s}
}

func (c *TaskController) HttpTaskListHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	log.Println(queryParam)
	title := queryParam.Get("title")

	t, err := c.service.GetTaskService(context.Background(), title)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var taskResponse []models.TaskResponse
	for _, v := range t {
		taskResponse = append(taskResponse, models.TaskResponse{
			ID:     v.ID,
			Title:  v.Title,
			Detail: v.Detail,
		})
	}

	var tasksRespose models.TasksResponse
	tasksRespose.Tasks = taskResponse

	res, _ := json.MarshalIndent(tasksRespose.Tasks, "", "\t\t")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
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

func (c *TaskController) CreateTask(
	ctx context.Context, req *connect.Request[taskv1.CreateTaskRequest],
) (*connect.Response[taskv1.CreateTaskResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	task := models.Task{
		Title:  req.Msg.Title,
		Detail: req.Msg.Detail,
	}

	if err := c.service.CreateTaskService(ctx, task); err != nil {
		return nil, err
	}

	res := connect.NewResponse(&taskv1.CreateTaskResponse{
		Creating: fmt.Sprintf("Creating %s ", req.Msg.Title),
	})
	res.Header().Set("Task-Version", "v1")
	return res, nil
}
