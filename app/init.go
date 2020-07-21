package app

import (
	"net/http"

	"github.com/gorilla/mux"

	"goSkeleton/adapters/repository/person"
	"goSkeleton/adapters/rest"
	"goSkeleton/usecases"
)

func (service *PeopleServerService) init() {
	service.PeopleRepository = person.NewPeopleInMemory()
	service.Usecases = usecases.NewUsecasesHandler(service.PeopleRepository)
	service.RestAdapter = rest.NewAdapter(service.Usecases)
	routes := []Route{
		{
			Name:    "GetPerson",
			Method:  "GET",
			Path:    "/person/{id}",
			Handler: service.RestAdapter.GetPerson,
		},
		{
			Name:    "AddPerson",
			Method:  "POST",
			Path:    "/person",
			Handler: service.RestAdapter.AddPerson,
		},
	}
	service.Router = initRouter(routes)
}

type Route struct {
	Name   string
	Method string
	Path   string
	//TODO: This should be string/enum reference that is mapped to a handler function, not the handler function itself
	// This would allow us to put the handler as a name in a config file or environment variable and make all this
	// configurable
	Handler http.HandlerFunc
}

func initRouter(routes []Route) *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.Name(route.Name).Methods(route.Method).Path(route.Path).HandlerFunc(route.Handler)
	}
	return router
}
