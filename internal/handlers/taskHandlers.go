package handlers

import (
	"Projects/internal/taskService"
	"Projects/internal/web/tasks"
	"context"
)

type TaskHandler struct {
	Service *taskService.TaskService
}

func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: *taskRequest.UserId,
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	return response, nil
}

func (h *TaskHandler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := request.Id
	updateData := request.Body

	taskToUpdate := taskService.Task{
		Task:   *updateData.Task,
		IsDone: *updateData.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskByID(uint(id), taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.Task{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return tasks.PatchTasksId200JSONResponse(response), nil
}

func (h *TaskHandler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	err := h.Service.DeleteTaskByID(uint(id))
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *TaskHandler) GetTasksByUserID(ctx context.Context, request tasks.GetTasksByUserIDRequestObject) (tasks.GetTasksByUserIDResponseObject, error) {
	userID := request.Id

	tasksList, err := h.Service.GetTasksByUserID(uint(userID))
	if err != nil {
		return nil, err
	}

	var response tasks.GetTasksByUserID200JSONResponse

	for _, t := range tasksList {
		task := tasks.Task{
			Id:     &t.ID,
			Task:   &t.Task,
			IsDone: &t.IsDone,
			UserId: &t.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}
