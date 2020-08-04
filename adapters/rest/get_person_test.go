package rest_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"

	"goSkeleton/adapters/rest"
	"goSkeleton/domain"
	"goSkeleton/internal/logging"
	"goSkeleton/usecases/mocks"
)

func TestAdapter_GetPerson(t *testing.T) {
	name := "Alberta Bobbeth Charleson"
	id := "0190fa37-eb2f-49ab-bb20-29c7033adbb7"
	Convey("Given the REST adapter exists", t, func() {
		ctrl := gomock.NewController(t)
		usecases := mocks.NewMockUsecases(ctrl)
		restAdapter := rest.NewAdapter(usecases)
		request := httptest.NewRequest("GET", "/person/"+id, nil)
		request = mux.SetURLVars(request, map[string]string{
			"id": id,
		})
		request = logging.AddRestRequestLogger(request)
		response := httptest.NewRecorder()

		Convey("And the person's ID can't be found, return an error", func() {
			usecases.EXPECT().GetPersonCase(id).Return(nil, errors.New("it broke"))
			restAdapter.GetPerson(response, request)
			So(response.Result().StatusCode, ShouldEqual, http.StatusInternalServerError)
		})
		Convey("And the person's ID exists, return the person's data", func() {
			expectedPerson := domain.Person{
				ID:   id,
				Name: name,
			}
			usecases.EXPECT().GetPersonCase(id).Return(&expectedPerson, nil)
			restAdapter.GetPerson(response, request)
			So(response.Result().StatusCode, ShouldEqual, http.StatusOK)

			bodyBytes, err := ioutil.ReadAll(response.Result().Body)
			So(err, ShouldBeNil)
			var responseStruct domain.Person
			err = json.Unmarshal(bodyBytes, &responseStruct)
			So(err, ShouldBeNil)
			So(responseStruct, ShouldResemble, expectedPerson)
		})
	})
}
