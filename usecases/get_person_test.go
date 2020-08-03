package usecases_test

import (
	"goSkeleton/domain"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"goSkeleton/adapters/repository/person"
	"goSkeleton/usecases"
)

func TestUsecaseHandler_GetPersonCase(t *testing.T) {
	name := "Alberta Bobbeth Charleson"
	Convey("Given a Usecase Handler with an in-memory people repo", t, func() {
		peopleRepo := person.NewPeopleInMemory()
		handler := usecases.NewUsecasesHandler(peopleRepo)
		Convey("And there is no record matching the ID in the repository, an error is returned", func() {
			personResult, err := handler.GetPersonCase("some-test-id")
			So(err, ShouldNotBeNil)
			So(personResult, ShouldBeNil)
		})
		Convey("And a person record exists in the repository", func() {
			id, err := peopleRepo.AddPerson(name)
			So(err, ShouldBeNil)
			So(id, ShouldNotBeNil)
			Convey("And an empty ID is passed, an error is returned", func() {
				personResult, err := handler.GetPersonCase("")
				So(err, ShouldNotBeNil)
				So(personResult, ShouldBeNil)
			})
			Convey("And a matching ID is passed, the person data is returned", func() {
				personResult, err := handler.GetPersonCase(id)
				So(err, ShouldBeNil)
				So(personResult, ShouldResemble, &domain.Person{
					ID:   id,
					Name: name,
				})
			})
		})
	})
}
