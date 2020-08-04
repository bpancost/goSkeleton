package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"goSkeleton/adapters/rest"
	"goSkeleton/internal/config"
	"goSkeleton/internal/logging"
)

func (service *PeopleServerService) initRest(config config.Config) {
	service.RestAdapter = rest.NewAdapter(service.Usecases)

	endpoints, err := config.GetListOfMaps("server.endpoints")
	if err != nil {
		logging.Panic(err)
	}
	router := mux.NewRouter()
	if len(endpoints) == 0 {
		logging.Info("Skipping RESTful listener, no endpoints defined")
		return
	}
	for _, endpoint := range endpoints {
		handler, err := service.RestAdapter.GetHandler(rest.AdapterName(endpoint["handler"]))
		if err != nil {
			logging.Panic(err)
		}
		router.Name(endpoint["name"]).Methods(endpoint["method"]).Path(endpoint["path"]).HandlerFunc(handler)
	}
	service.Router = router
	service.Router.Use(loggingMiddleware)

	address := config.GetString("server.address.ip") + ":" + config.GetString("server.address.port")
	service.Server = &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * config.GetDuration("server.timeout.write"),
		ReadTimeout:  time.Second * config.GetDuration("server.timeout.read"),
		IdleTimeout:  time.Second * config.GetDuration("server.timeout.idle"),
		Handler:      service.Router,
	}
	logging.Infof("starting RESTful server on: %s", address)
	go func() {
		if err := service.Server.ListenAndServe(); err != nil {
			logging.Error(err)
		}
	}()
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		next.ServeHTTP(w, logging.AddRestRequestLogger(req))
	})
}
