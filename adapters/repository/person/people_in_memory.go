package person

import (
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"goSkeleton/domain"
)

type People struct {
	PeopleByIDs map[string]*domain.Person
}

func NewPeopleInMemory() *People {
	return &People{
		PeopleByIDs: make(map[string]*domain.Person),
	}
}

func (p *People) GetPerson(id string) (*domain.Person, error) {
	if id == "" {
		return nil, errors.New("id can not be empty")
	}
	person, ok := p.PeopleByIDs[id]
	if !ok {
		return nil, errors.New("person doesn't exist")
	}
	return person.Copy(), nil
}

func (p *People) AddPerson(name string) (string, error) {
	if name == "" {
		return "", errors.New("name can not be empty")
	}
	uuidObj, err := uuid.NewV4()
	if err != nil {
		return "", errors.Wrap(err, "uuid generation error")
	}
	id := uuidObj.String()
	p.PeopleByIDs[id] = &domain.Person{
		ID:   id,
		Name: name,
	}
	return id, nil
}

func (p *People) Close() error {
	// Nothing to do for the in memory version
	return nil
}
