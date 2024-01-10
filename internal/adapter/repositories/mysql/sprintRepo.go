package mysql_repo

import (
	"stand_up/internal/core/domain"
	"stand_up/internal/ports"

	"gorm.io/gorm/clause"
)

type SprintRepository struct {
}

func NewSprintRepository() ports.Repository {
	return &SprintRepository{}
}

func (repo *SprintRepository) GetAll(param map[string]interface{}) (interface{}, error) {
	model := repo.Model()

	q := db.Preload(clause.Associations).Find(model)
	return model, q.Error
}

func (repo *SprintRepository) Find(id string) (interface{}, error) {
	model := repo.Model()
	q := db.Preload(clause.Associations).Where("id = ?", id).First(model)

	return model, q.Error
}

func (repo *SprintRepository) Create(data interface{}) (interface{}, error) {
	model := repo.Model()
	q := db.Model(data).Create(data)
	return model, q.Error

}

func (repo *SprintRepository) Update(id string, data interface{}) (interface{}, error) {
	model := repo.Model()
	q := db.Model(&model).Where("id = ?", id).Updates(model)
	if q.Error != nil {
		return nil, q.Error
	}
	return repo.Find(id)
}

func (repo *SprintRepository) Delete(id string) (interface{}, error) {
	model := repo.Model()
	q := db.Model(&model).Where("id = ?", id).Delete(model)
	return model, q.Error
}

func (repo *SprintRepository) Model() domain.Sprint {
	return domain.Sprint{}
}

func (repo *SprintRepository) ArrayModel() []domain.StandUpdate {
	return []domain.StandUpdate{}
}
