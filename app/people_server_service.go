package app

import (
	"context"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"goSkeleton/adapters/rest"
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
	service.init()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./cmd/peopleServer")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Panic(err)
	}

	address := viper.GetString("server.address.ip") + ":" + viper.GetString("server.address.port")
	service.Server = &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * viper.GetDuration("server.timeout.write"),
		ReadTimeout:  time.Second * viper.GetDuration("server.timeout.read"),
		IdleTimeout:  time.Second * viper.GetDuration("server.timeout.idle"),
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
