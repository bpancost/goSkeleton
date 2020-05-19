package usecases

import "skeleton/domain"

type People interface {
	// Gets a person by their ID
	GetPerson(id string) (*domain.Person, error)
	// Creates a new person with the given name and returns the ID to the new record
	AddPerson(name string) (string, error)
	// Close the repository
	Close() error
}
