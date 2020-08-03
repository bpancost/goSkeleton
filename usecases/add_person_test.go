package usecases_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"goSkeleton/adapters/repository/person"
	"goSkeleton/usecases"
)

func TestUsecaseHandler_AddPersonCase(t *testing.T) {
	name := "Alberta Bobbeth Charleson"
	Convey("Given a Usecase Handler with an in-memory people repo", t, func() {
		peopleRepo := person.NewPeopleInMemory()
		handler := usecases.NewUsecasesHandler(peopleRepo)
		Convey("And no name is passed, then an error is returned", func() {
			id, err := handler.AddPersonCase("")
			So(err, ShouldNotBeNil)
			So(id, ShouldBeEmpty)
		})
		Convey("And a valid name is passed, then an ID is returned", func() {
			id, err := handler.AddPersonCase(name)
			So(err, ShouldBeNil)
			So(id, ShouldNotBeNil)
		})
	})
}
