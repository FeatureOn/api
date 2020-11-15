package memory

type DataContext struct {
	userRepository UserRepository
}

func NewDataContext() DataContext {
	dataContext := DataContext{}
	dataContext.userRepository = newUserRepository()
	return dataContext
}

// CheckConnection always returns true because there's no problem connecting to memory
func (dc DataContext) CheckConnection() bool {
	return true
}
