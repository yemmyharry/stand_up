package services

import (
	"stand_up/internal/adapter/api/requests"
	mysql_repo "stand_up/internal/adapter/repositories/mysql"
	"stand_up/internal/core/domain"
	"stand_up/internal/core/helper"
	ports "stand_up/internal/ports"
)

type TaskService struct {
	repo ports.Repository
}

func NewTaskService() *TaskService {
	return &TaskService{
		repo: mysql_repo.NewStandupRepository(),
	}
}

func (s TaskService) Save(input requests.TaskRequest) (interface{}, error) {
	task := domain.Task{}
	helper.Copy(input, &task)
	return s.repo.Create(task)
}

func (s TaskService) GetAll(param map[string]interface{}) (interface{}, error) {

	return s.repo.GetAll(param)
}
