package app

import (
	"net/http"

	"github.com/gorilla/mux"

	"goSkeleton/adapters/repository/person"
	"goSkeleton/adapters/rest"
	"goSkeleton/internal/config"
	"goSkeleton/internal/logging"
	"goSkeleton/usecases"
)

func (service *PeopleServerService) init(config config.Config) {
	service.PeopleRepository = person.NewPeopleInMemory()
	service.Usecases = usecases.NewUsecasesHandler(service.PeopleRepository)
	service.RestAdapter = rest.NewAdapter(service.Usecases)

	endpoints := config.Get("server.endpoints")
	var routes []Route
	if endpoints, ok := endpoints.([]interface{}); ok {
		routes = make([]Route, len(endpoints))
		for index, endpointRaw := range endpoints {
			if endpoint, ok := endpointRaw.(map[interface{}]interface{}); ok {
				handler, err := service.RestAdapter.GetHandler(rest.AdapterName(endpoint["handler"].(string)))
				if err != nil {
					logging.Panic(err)
				}
				routes[index] = Route{
					Name:    endpoint["name"].(string),
					Method:  endpoint["method"].(string),
					Path:    endpoint["path"].(string),
					Handler: handler,
				}
			} else {
				logging.Panic("server.endpoints configuration has no objects in the list")
			}
		}
	} else {
		logging.Panic("server.endpoints configuration is not a list")
	}
	service.Router = initRouter(routes)
}

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func initRouter(routes []Route) *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.Name(route.Name).Methods(route.Method).Path(route.Path).HandlerFunc(route.Handler)
	}
	return router
}
