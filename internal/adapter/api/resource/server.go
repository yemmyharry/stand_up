package resource

import (
	"stand_up/internal/core/services"
)

type HTTPHandler struct {
	standupUpdate *services.StandupUpdateService
	sprintService *services.SprintService
	taskService   *services.TaskService
}

func NewHTTPHandler() *HTTPHandler {
	handler := &HTTPHandler{
		standupUpdate: services.NewStandupUpdateService(),
		sprintService: services.NewSprintService(),
		taskService:   services.NewTaskService(),
	}
	return handler
}
