package services

import (
	"stand_up/internal/adapter/api/requests"
	mysql_repo "stand_up/internal/adapter/repositories/mysql"
	"stand_up/internal/core/domain"
	"stand_up/internal/core/helper"
	ports "stand_up/internal/ports"
	"time"
)

type StandupUpdateService struct {
	repo        ports.Repository
	standupRepo ports.Repository
}

func NewStandupUpdateService() *StandupUpdateService {
	return &StandupUpdateService{
		repo:        mysql_repo.NewStandupRepository(),
		standupRepo: mysql_repo.NewSprintRepository(),
	}
}

func (s StandupUpdateService) Save(input requests.StandUpdateRequest) (interface{}, error) {
	standupRepo, err := s.standupRepo.Find(input.StandupID)

	standup := domain.Standup{}
	helper.Copy(standupRepo, &standup)
	if err != nil {
		return nil, err
	}

	standupUpdate := domain.StandUpdate{}

	helper.Copy(input, &standupUpdate)
	currHr, curMin, _ := time.Now().Clock()
	strHr, strMin, _ := standup.EndTime.Clock()
	_, endMin, _ := standup.EndTime.Clock()

	if currHr <= strHr {
		if currHr == strHr && curMin >= strMin && curMin <= endMin {
			standupUpdate.Status = domain.Within
		} else {
			standupUpdate.Status = domain.Before
		}
	} else {
		standupUpdate.Status = domain.After
	}

	return s.repo.Create(standup)
}

func (s StandupUpdateService) GetAll(param map[string]interface{}) (interface{}, error) {

	return s.repo.GetAll(param)
}
