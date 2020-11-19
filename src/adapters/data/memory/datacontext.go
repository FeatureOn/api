package memory

// DataContext represents a struct that holds concrete repositories
type DataContext struct {
	UserRepository   UserRepository
	HealthRepository HealthRepository
}

// NewDataContext returns a new memory backed DataContext
func NewDataContext() DataContext {
	dataContext := DataContext{}
	dataContext.UserRepository = newUserRepository()
	dataContext.HealthRepository = newHealthRepository()
	return dataContext
}
