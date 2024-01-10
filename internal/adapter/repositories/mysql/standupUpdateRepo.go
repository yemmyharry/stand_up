package mysql_repo

import (
	"stand_up/internal/core/domain"
	"stand_up/internal/database"
	"stand_up/internal/ports"

	"gorm.io/gorm/clause"
)

type StandupRepository struct {
}

var db = database.ConnectDB()

func NewStandupRepository() ports.Repository {
	return &StandupRepository{}
}

func (repo *StandupRepository) GetAll(param map[string]interface{}) (interface{}, error) {
	model := repo.Model()

	q := db.Preload(clause.Associations)
	if param["day"] != nil {
		q.Where("DAY(created_at) == ?", param["day"])

	}
	if param["sprint_id"] != nil {
		q.Where("sprint_id == ?", param["sprint_id"])

	}
	if param["user_id"] != nil {
		q.Where("user_id == ?", param["user_id"])

	}
	if param["week"] != nil {
		q.Where("WEEK(created_at) == ?", param["week"])

	}

	q.Find(model)
	return model, q.Error
}

func (repo *StandupRepository) Find(id string) (interface{}, error) {
	model := repo.Model()
	q := db.Preload(clause.Associations).Where("id = ?", id).First(model)

	return model, q.Error
}

func (repo *StandupRepository) Create(data interface{}) (interface{}, error) {
	model := repo.Model()
	q := db.Model(data).Create(data)
	return model, q.Error

}

func (repo *StandupRepository) Update(id string, data interface{}) (interface{}, error) {
	model := repo.Model()
	q := db.Model(&model).Where("id = ?", id).Updates(model)
	if q.Error != nil {
		return nil, q.Error
	}
	return repo.Find(id)
}

func (repo *StandupRepository) Delete(id string) (interface{}, error) {
	model := repo.Model()
	q := db.Model(&model).Where("id = ?", id).Delete(model)
	return model, q.Error
}

func (repo *StandupRepository) Model() domain.StandUpdate {
	return domain.StandUpdate{}
}

func (repo *StandupRepository) ArrayModel() domain.StandUpdate {
	return domain.StandUpdate{}
}
