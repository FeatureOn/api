package memory

type DataContext struct {
	UserRepository   UserRepository
	HealthRepository HealthRepository
}

func NewDataContext() DataContext {
	dataContext := DataContext{}
	dataContext.UserRepository = newUserRepository()
	dataContext.HealthRepository = newHealthRepository()
	return dataContext
}

// CheckConnection always returns true because there's no problem connecting to memory
func (dc DataContext) checkConnection() bool {
	return true
}
