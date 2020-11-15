package application

// HealthRepository is the interface to interact with database
type HealthRepository interface {
	Ready() bool
}

type HealthService struct {
	HealthRepository HealthRepository
}

func NewHealthService(hr HealthRepository) HealthService {
	if hr == nil {
		panic("missing HealthRepository")
	}
	return HealthService{
		HealthRepository: hr,
	}
}

func (hs HealthService) Ready() bool {
	return hs.HealthRepository.Ready()
}
