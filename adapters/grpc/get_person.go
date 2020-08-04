package grpc

import (
	"context"
	"goSkeleton/internal/logging"
	"goSkeleton/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (adapter Adapter) GetPerson(ctx context.Context, request *proto.GetPersonRequest) (*proto.GetPersonResponse, error) {
	logger := logging.GetGrpcContextLogger(ctx)
	person, err := adapter.Usecases.GetPersonCase(request.Id)
	if err != nil {
		logger.Errorf("failed to fetch person: %v", err)
		return &proto.GetPersonResponse{}, status.Errorf(codes.Internal, "failed to fetch person")
	}
	return &proto.GetPersonResponse{
		Id:   person.ID,
		Name: person.Name,
	}, nil
}
