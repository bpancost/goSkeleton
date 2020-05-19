package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"skeleton/adapters/rest"
	"skeleton/usecases"
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
	service.init()

	address := "0.0.0.0:8080"
	service.Server = &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      service.Router,
	}
	logrus.Infof("starting server on: %s", address)
	go func() {
		if err := service.Server.ListenAndServe(); err != nil {
			logrus.Error(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	<-done

	logrus.Info("shutting down")
	service.Shutdown()
}

func (service *PeopleServerService) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if err := service.Server.Shutdown(ctx); err != nil {
		logrus.Error(err)
	}
	if err := service.PeopleRepository.Close(); err != nil {
		logrus.Error(err)
	}
}
