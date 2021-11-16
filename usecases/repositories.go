package usecases

import "goSkeleton/domain"

type People interface {
	// GetPerson Gets a person by their ID
	GetPerson(id string) (*domain.Person, error)
	// AddPerson Creates a new person with the given name and returns the ID to the new record
	AddPerson(name string) (string, error)
	// Close the repository
	Close() error
}
