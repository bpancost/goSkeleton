package rest_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	"goSkeleton/adapters/rest"
	"goSkeleton/internal/logging"
	"goSkeleton/usecases/mocks"
)

func TestAdapter_AddPerson(t *testing.T) {
	name := "Alberta Bobbeth Charleson"
	id := "0190fa37-eb2f-49ab-bb20-29c7033adbb7"
	Convey("Given the REST adapter exists", t, func() {
		ctrl := gomock.NewController(t)
		usecases := mocks.NewMockUsecases(ctrl)
		restAdapter := rest.NewAdapter(usecases)
		Convey("And a bad request, an error is returned", func() {
			request := httptest.NewRequest("POST", "/person", nil)
			request = logging.AddRequestLogger(request)
			response := httptest.NewRecorder()
			restAdapter.AddPerson(response, request)
			So(response.Result().StatusCode, ShouldEqual, http.StatusBadRequest)
		})
		Convey("And a good request", func() {
			requestBody := rest.AddPersonRequest{Name: name}
			requestBodyBytes, err := json.Marshal(requestBody)
			So(err, ShouldBeNil)

			request := httptest.NewRequest("POST", "/person", bytes.NewReader(requestBodyBytes))
			request = logging.AddRequestLogger(request)
			response := httptest.NewRecorder()

			Convey("And an error is returned by the use case, an error is returned", func() {
				usecases.EXPECT().AddPersonCase(name).Return("", errors.New("it broke"))
				restAdapter.AddPerson(response, request)
				So(response.Result().StatusCode, ShouldEqual, http.StatusInternalServerError)
			})
			Convey("And the use case is successful, a valid response is returned", func() {
				usecases.EXPECT().AddPersonCase(name).Return(id, nil)
				restAdapter.AddPerson(response, request)
				So(response.Result().StatusCode, ShouldEqual, http.StatusOK)
				bodyBytes, err := ioutil.ReadAll(response.Result().Body)
				So(err, ShouldBeNil)
				var responseStruct rest.AddPersonResponse
				err = json.Unmarshal(bodyBytes, &responseStruct)
				So(err, ShouldBeNil)
				So(responseStruct, ShouldResemble, rest.AddPersonResponse{ID: id})
			})
		})
	})
}
