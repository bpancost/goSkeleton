package app

import (
	"context"
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcAdapter "goSkeleton/adapters/grpc"
	"goSkeleton/internal/config"
	"goSkeleton/internal/logging"
	"goSkeleton/proto"
	"google.golang.org/grpc"
	"net"
)

func (service *PeopleServerService) initGrpc(config config.Config) {
	service.GrpcAdapter = grpcAdapter.NewAdapter(service.Usecases)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		logging.Panic(err)
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(UnaryLoggingInterceptor())))
	proto.RegisterPeopleServer(server, service.GrpcAdapter)

	server.Serve(listener)
}

func UnaryLoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		newCtx := logging.AddGrpcContextLogger(ctx, info.FullMethod)
		resp, err = handler(newCtx, req)
		return
	}
}
