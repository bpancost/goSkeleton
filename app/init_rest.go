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

	endpoints, err := config.GetListOfMaps("rest.endpoints")
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
	router.Use(loggingMiddleware)

	address := config.GetString("rest.address.ip") + ":" + config.GetString("rest.address.port")
	service.RestServer = &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * config.GetDuration("rest.timeout.write"),
		ReadTimeout:  time.Second * config.GetDuration("rest.timeout.read"),
		IdleTimeout:  time.Second * config.GetDuration("rest.timeout.idle"),
		Handler:      router,
	}
	logging.Infof("starting RESTful server on: %s", address)
	go func() {
		if err := service.RestServer.ListenAndServe(); err != nil {
			logging.Error(err)
		}
	}()
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		next.ServeHTTP(w, logging.AddRestRequestLogger(req))
	})
}
