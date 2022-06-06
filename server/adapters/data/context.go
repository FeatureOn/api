package data

import "github.com/FeatureOn/api/server/application"

// DataContext represents a struct that holds concrete repositories
type DataContext struct {
	UserRepository    application.UserRepository
	HealthRepository  application.HealthRepository
	ProductRepository application.ProductRepository
	FlagRepository    application.FlagRepository
}
