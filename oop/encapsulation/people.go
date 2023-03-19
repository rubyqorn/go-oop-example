package encapsulation

import "fmt"

type People struct {
	Id                        int
	Name, Surname, Occupation string
}

func (people *People) SetPeople(id int, name, surname, occupation string) {
	people.SetId(id)
	people.SetName(name)
	people.SetSurname(surname)
	people.SetOccupation(occupation)
}

func (people *People) Introduce() {
	fmt.Printf(
		"Id: %d\nName: %s\nSurname: %s\nOccupation: %s\n",
		people.GetId(),
		people.GetName(),
		people.GetSurname(),
		people.GetOccupation(),
	)
}

func (people *People) SetId(id int) {
	people.Id = id
}

func (people *People) GetId() int {
	return people.Id
}

func (people *People) SetName(name string) {
	people.Name = name
}

func (people *People) GetName() string {
	return people.Name
}

func (people *People) SetSurname(surname string) {
	people.Surname = surname
}

func (people *People) GetSurname() string {
	return people.Surname
}

func (people *People) SetOccupation(occupation string) {
	people.Occupation = occupation
}

func (people *People) GetOccupation() string {
	return people.Occupation
}
