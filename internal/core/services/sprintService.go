package services

import (
	"stand_up/internal/adapter/api/requests"
	mysql_repo "stand_up/internal/adapter/repositories/mysql"
	"stand_up/internal/core/domain"
	"stand_up/internal/core/helper"
	ports "stand_up/internal/ports"
)

type SprintService struct {
	repo ports.Repository
}

func NewSprintService() *SprintService {
	return &SprintService{
		repo: mysql_repo.NewSprintRepository(),
	}
}

func (s SprintService) Save(input requests.SprintRequest) (interface{}, error) {
	sprint := domain.Sprint{}
	helper.Copy(input, &sprint)
	return s.repo.Create(&sprint)
}

func (s SprintService) GetAll(param map[string]interface{}) (interface{}, error) {

	return s.repo.GetAll(param)
}
