package grpc_test

import (
	"context"
	"errors"
	"goSkeleton/domain"
	"goSkeleton/internal/logging"
	"goSkeleton/proto"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	"goSkeleton/adapters/grpc"
	"goSkeleton/usecases/mocks"
)

func TestAdapter_GetPerson(t *testing.T) {
	name := "Alberta Bobbeth Charleson"
	id := "0190fa37-eb2f-49ab-bb20-29c7033adbb7"
	Convey("Given the gRPC adapter exists", t, func() {
		ctrl := gomock.NewController(t)
		usecases := mocks.NewMockUsecases(ctrl)
		grpcAdapter := grpc.NewAdapter(usecases)
		ctx := logging.AddGrpcContextLogger(context.Background(), "/People/GetPerson")
		request := &proto.GetPersonRequest{
			Id: id,
		}
		Convey("And the person's ID can't be found, return an error", func() {
			usecases.EXPECT().GetPersonCase(id).Return(nil, errors.New("it broke"))
			response, err := grpcAdapter.GetPerson(ctx, request)
			So(err, ShouldNotBeNil)
			So(response, ShouldResemble, &proto.GetPersonResponse{})
		})
		Convey("And the person's ID exists, return the person's data", func() {
			expectedPerson := domain.Person{
				ID:   id,
				Name: name,
			}
			usecases.EXPECT().GetPersonCase(id).Return(&expectedPerson, nil)
			response, err := grpcAdapter.GetPerson(ctx, request)
			So(err, ShouldBeNil)
			So(response, ShouldResemble, &proto.GetPersonResponse{
				Id:   id,
				Name: name,
			})
		})
	})
}
