package usecases

import "goSkeleton/domain"

//go:generate mockgen -destination=./mocks/mocks.go -package=mocks goSkeleton/usecases Usecases

type Usecases interface {
	// AddPersonCase Adds a person by their name to the database and returns the new ID
	AddPersonCase(name string) (string, error)
	// GetPersonCase Gets a person from the database by their ID
	GetPersonCase(id string) (*domain.Person, error)
}

type UsecaseHandler struct {
	peopleRepository People
}

func NewUsecasesHandler(peopleRepository People) UsecaseHandler {
	return UsecaseHandler{
		peopleRepository: peopleRepository,
	}
}
