package domain

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// shallow copy of the person struct
func (person *Person) Copy() *Person {
	personCopy := *person
	return &personCopy
}
