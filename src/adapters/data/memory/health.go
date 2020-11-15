package memory

// // Health object represents the Health model to hold in memory
// type Health struct {
// 	ID       string
// 	Name     string
// 	HealthName string
// 	Password string
// }

type HealthRepository struct{}

func newHealthRepository() HealthRepository {
	return HealthRepository{}
}

func (hr HealthRepository) Ready() bool {
	return true
}
