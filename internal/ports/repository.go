package ports

type HealthCheckRepository interface {
	HealthCheck() error
}

type Repository interface {
	Create(data interface{}) (interface{}, error)
	Find(id string) (interface{}, error)
	GetAll(param map[string]interface{}) (interface{}, error)
	Update(id string, data interface{}) (interface{}, error)
	Delete(id string) (interface{}, error)
}
