package mysql_repo

import (
	"stand_up/internal/core/domain"
	"stand_up/internal/ports"

	"gorm.io/gorm/clause"
)

type TaskRepository struct {
}

func NewTaskRepository() ports.Repository {
	return &TaskRepository{}
}

func (repo *TaskRepository) GetAll(param map[string]interface{}) (interface{}, error) {
	model := repo.Model()

	q := db.Preload(clause.Associations).Find(&model)
	return model, q.Error
}

func (repo *TaskRepository) Find(id string) (interface{}, error) {
	model := repo.Model()
	q := db.Preload(clause.Associations).Where("id = ?", id).First(model)

	return model, q.Error
}

func (repo *TaskRepository) Create(data interface{}) (interface{}, error) {
	model := repo.Model()
	q := db.Model(data).Create(data)
	return model, q.Error

}

func (repo *TaskRepository) Update(id string, data interface{}) (interface{}, error) {
	model := repo.Model()
	q := db.Model(&model).Where("id = ?", id).Updates(model)
	if q.Error != nil {
		return nil, q.Error
	}
	return repo.Find(id)
}

func (repo *TaskRepository) Delete(id string) (interface{}, error) {
	model := repo.Model()
	q := db.Model(&model).Where("id = ?", id).Delete(model)
	return model, q.Error
}

func (repo *TaskRepository) Model() domain.Task {
	return domain.Task{}
}

func (repo *TaskRepository) ArrayModel() []domain.Task {
	return []domain.Task{}
}
