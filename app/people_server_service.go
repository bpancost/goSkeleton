package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"goSkeleton/adapters/rest"
	"goSkeleton/app/config"
	"goSkeleton/internal/logging"
	"goSkeleton/usecases"
)

type PeopleServerService struct {
	PeopleRepository usecases.People
	Usecases         usecases.Usecases
	RestAdapter      rest.Adapter
	Router           *mux.Router
	Server           *http.Server
}

func NewPeopleServerService() PeopleServerService {
	return PeopleServerService{}
}

func (service *PeopleServerService) Start() {
	hostname, err := os.Hostname()
	if err != nil {
		logging.Panic(err)
	}
	logging.GetLoggerWithBaseFields("peopleServer", hostname)

	conf, err := config.NewViperConfig("peopleServer")
	if err != nil {
		logging.Panic(err)
	}

	service.init(conf)

	address := conf.GetString("server.address.ip") + ":" + conf.GetString("server.address.port")
	service.Server = &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * conf.GetDuration("server.timeout.write"),
		ReadTimeout:  time.Second * conf.GetDuration("server.timeout.read"),
		IdleTimeout:  time.Second * conf.GetDuration("server.timeout.idle"),
		Handler:      service.Router,
	}
	logging.Infof("starting server on: %s", address)
	go func() {
		if err := service.Server.ListenAndServe(); err != nil {
			logging.Error(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	<-done

	logging.Info("shutting down")
	service.Shutdown()
}

func (service *PeopleServerService) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if err := service.Server.Shutdown(ctx); err != nil {
		logging.Error(err)
	}
	if err := service.PeopleRepository.Close(); err != nil {
		logging.Error(err)
	}
}
