package domain

// Product is the main object of the system holding both features and environments
type Product struct {
	ID           string
	Name         string
	Features     []Feature
	Environments []Environment
}

// Feature is a basic flag (as for now) holding a key within a project and its default state
type Feature struct {
	Key          string
	DefaultState bool
}

// Environment is a struct that will hold the collection of flags for each of product's deployment
type Environment struct {
	ID   string
	Name string
}
