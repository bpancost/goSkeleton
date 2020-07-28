package rest

import (
	"github.com/pkg/errors"
	"net/http"
)

type AdapterName string

const (
	GetPerson AdapterName = "GetPerson"
	AddPerson AdapterName = "AddPerson"
)

func (adapter Adapter) GetHandler(name AdapterName) (http.HandlerFunc, error) {
	switch name {
	case GetPerson:
		return adapter.GetPerson, nil
	case AddPerson:
		return adapter.AddPerson, nil
	default:
		return nil, errors.Errorf("rest adapter %s doesn't exist", name)
	}
}
