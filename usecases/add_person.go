package usecases

import (
	"github.com/pkg/errors"
)

func (handler UsecaseHandler) AddPersonCase(name string) (string, error) {
	id, err := handler.peopleRepository.AddPerson(name)
	if err != nil {
		return "", errors.Wrap(err, "repository error adding person")
	}
	return id, nil
}
