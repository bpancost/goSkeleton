package grpc_test

import (
	"context"
	"errors"
	"goSkeleton/proto"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	"goSkeleton/adapters/grpc"
	"goSkeleton/internal/logging"
	"goSkeleton/usecases/mocks"
)

func TestAdapter_AddPerson(t *testing.T) {
	name := "Alberta Bobbeth Charleson"
	id := "0190fa37-eb2f-49ab-bb20-29c7033adbb7"
	Convey("Given the gRPC adapter exists", t, func() {
		ctrl := gomock.NewController(t)
		usecases := mocks.NewMockUsecases(ctrl)
		grpcAdapter := grpc.NewAdapter(usecases)
		ctx := logging.AddGrpcContextLogger(context.Background(), "/People/AddPerson")
		Convey("And a bad request, an error is returned", func() {
			request := &proto.AddPersonRequest{
				Name: "",
			}
			response, err := grpcAdapter.AddPerson(ctx, request)
			So(err, ShouldNotBeNil)
			So(response, ShouldResemble, &proto.AddPersonResponse{})
		})
		Convey("And a good request", func() {
			request := &proto.AddPersonRequest{
				Name: name,
			}
			Convey("And an error is returned by the use case, an error is returned", func() {
				usecases.EXPECT().AddPersonCase(name).Return("", errors.New("it broke"))
				response, err := grpcAdapter.AddPerson(ctx, request)
				So(err, ShouldNotBeNil)
				So(response, ShouldResemble, &proto.AddPersonResponse{})
			})
			Convey("And the use case is successful, a valid response is returned", func() {
				usecases.EXPECT().AddPersonCase(name).Return(id, nil)
				response, err := grpcAdapter.AddPerson(ctx, request)
				So(err, ShouldBeNil)
				So(response, ShouldResemble, &proto.AddPersonResponse{Id: id})
			})
		})
	})
}
