package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"goSkeleton/internal/logging"
	"goSkeleton/proto"
)

func (adapter Adapter) AddPerson(ctx context.Context, request *proto.AddPersonRequest) (*proto.AddPersonResponse, error) {
	logger := logging.GetGrpcContextLogger(ctx)
	id, err := adapter.Usecases.AddPersonCase(request.Name)
	if err != nil {
		logger.Errorf("failed to add person: %v", err)
		return &proto.AddPersonResponse{}, status.Errorf(codes.Internal, "failed to add person")
	}
	return &proto.AddPersonResponse{Id: id}, nil
}
