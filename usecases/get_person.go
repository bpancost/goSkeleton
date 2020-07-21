package usecases

import (
	"github.com/pkg/errors"

	"goSkeleton/domain"
)

func (handler UsecaseHandler) GetPersonCase(id string) (*domain.Person, error) {
	person, err := handler.peopleRepository.GetPerson(id)
	if err != nil {
		return nil, errors.Wrap(err, "repository error fetching person")
	}
	return person, nil
}
