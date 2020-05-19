package person_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"skeleton/adapters/repository/person"
	"skeleton/domain"
)

func TestPeople_AddPerson(t *testing.T) {
	Convey("Given the People repository exists", t, func() {
		people := person.NewPeopleInMemory()
		Convey("And a name is passed in, an ID is returned", func() {
			name := "Alberta Bobbeth Charleson"
			id, err := people.AddPerson(name)

			So(err, ShouldBeNil)
			So(id, ShouldNotBeNil)
			So(id, ShouldNotBeEmpty)

			Convey("And the ID can be used to retrieve the name", func() {
				foundPerson, err := people.GetPerson(id)

				So(err, ShouldBeNil)
				So(foundPerson, ShouldResemble, &domain.Person{
					ID:   id,
					Name: name,
				})
			})
		})
		Convey("And a name is not passed in, an error is returned", func() {
			id, err := people.AddPerson("")

			So(err, ShouldNotBeNil)
			So(id, ShouldBeEmpty)
		})
	})
}

func TestPeople_GetPerson(t *testing.T) {
	Convey("Given the People repository exists", t, func() {
		people := person.NewPeopleInMemory()

		Convey("And a person exists within the repository", func() {
			name := "Alberta Bobbeth Charleson"
			id, err := people.AddPerson(name)

			So(err, ShouldBeNil)
			So(id, ShouldNotBeNil)
			So(id, ShouldNotBeEmpty)

			Convey("Then the person's ID can be used to retrieve the name", func() {
				foundPerson, err := people.GetPerson(id)

				So(err, ShouldBeNil)
				So(foundPerson, ShouldResemble, &domain.Person{
					ID:   id,
					Name: name,
				})
			})

			Convey("Then an ID that matches no record will return an error", func() {
				foundPerson, err := people.GetPerson(id + "GARBAGE")

				So(err, ShouldNotBeNil)
				So(foundPerson, ShouldBeNil)
			})

			Convey("Then an empty ID will always return an error", func() {
				foundPerson, err := people.GetPerson("")

				So(err, ShouldNotBeNil)
				So(foundPerson, ShouldBeNil)
			})
		})
	})
}
