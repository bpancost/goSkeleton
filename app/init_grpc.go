package app

import (
	"context"
	"fmt"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	grpcAdapter "goSkeleton/adapters/grpc"
	"goSkeleton/internal/config"
	"goSkeleton/internal/logging"
	"goSkeleton/proto"
)

func (service *PeopleServerService) initGrpc(config config.Config) {
	service.GrpcAdapter = grpcAdapter.NewAdapter(service.Usecases)

	port := config.GetInt("grpc.port")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logging.Panic(err)
	}
	service.GrpcServer = grpc.NewServer(grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(UnaryLoggingInterceptor())))
	proto.RegisterPeopleServer(service.GrpcServer, service.GrpcAdapter)

	logging.Infof("starting GRPC server on port: %d", port)
	go func() {
		if err := service.GrpcServer.Serve(listener); err != nil {
			logging.Error(err)
		}
	}()
}

func UnaryLoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		newCtx := logging.AddGrpcContextLogger(ctx, info.FullMethod)
		resp, err = handler(newCtx, req)
		return
	}
}
