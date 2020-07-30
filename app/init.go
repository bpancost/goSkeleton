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

	endpoints, err := config.GetListOfMaps("server.endpoints")
	if err != nil {
		logging.Panic(err)
	}
	router := mux.NewRouter()
	for _, endpoint := range endpoints {
		handler, err := service.RestAdapter.GetHandler(rest.AdapterName(endpoint["handler"]))
		if err != nil {
			logging.Panic(err)
		}
		router.Name(endpoint["name"]).Methods(endpoint["method"]).Path(endpoint["path"]).HandlerFunc(handler)
	}
	service.Router = router
	service.Router.Use(loggingMiddleware)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		next.ServeHTTP(w, logging.AddRequestLogger(req))
	})
}
