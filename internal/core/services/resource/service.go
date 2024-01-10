package services

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s Service) HealthCheck() error {
	return s.HealthCheck()
}
