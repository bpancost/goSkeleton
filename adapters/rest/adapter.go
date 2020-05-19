package rest

import "skeleton/usecases"

type Adapter struct {
	Usecases usecases.Usecases
}

func NewAdapter(usecases usecases.Usecases) Adapter {
	return Adapter{Usecases: usecases}
}
