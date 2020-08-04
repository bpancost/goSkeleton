package app

import (
	"context"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	"time"

	grpcAdapter "goSkeleton/adapters/grpc"
	"goSkeleton/adapters/repository/person"
	"goSkeleton/adapters/rest"
	"goSkeleton/internal/config"
	"goSkeleton/internal/logging"
	"goSkeleton/usecases"
)

type PeopleServerService struct {
	PeopleRepository usecases.People
	Usecases         usecases.Usecases
	RestAdapter      rest.Adapter
	GrpcAdapter      grpcAdapter.Adapter
	RestServer       *http.Server
	GrpcServer       *grpc.Server
}

func NewPeopleServerService() PeopleServerService {
	return PeopleServerService{}
}

func (service *PeopleServerService) Start() {
	hostname, err := os.Hostname()
	if err != nil {
		logging.Panic(err)
	}
	logging.GetLoggerWithBaseFields("people_server", hostname)

	conf, err := config.NewViperConfig("people_server")
	if err != nil {
		logging.Warn(err)
	}

	service.PeopleRepository = person.NewPeopleInMemory()
	service.Usecases = usecases.NewUsecasesHandler(service.PeopleRepository)

	service.initRest(conf)
	service.initGrpc(conf)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	<-done

	logging.Info("shutting down")
	service.Shutdown()
}

func (service *PeopleServerService) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	service.GrpcServer.GracefulStop()
	if err := service.RestServer.Shutdown(ctx); err != nil {
		logging.Error(err)
	}
	if err := service.PeopleRepository.Close(); err != nil {
		logging.Error(err)
	}
}
