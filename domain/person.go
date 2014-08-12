package domain

import (
	"fmt"
	"time"
)

type Person interface {
	fmt.Stringer
	GetFirstName() string
	GetLastName() string
	GetBirthdate() time.Time
	GetDeathdate() time.Time
	GetGender() string
	GetHeight() Measurement
	GetWeight() Measurement
}

type PersonRepository interface {
	Repository
	SavePerson(x Person) (string, error)
	GetPerson(id string) (Person, error)
	AllPeople() ([]Person, error)
}

var _ Person = (*person)(nil)

type gender byte

func (g gender) String() string {
	switch g {
	case female:
		return "female"
	case male:
		return "male"
	default:
		return "unspecified"
	}
}

const (
	unknown gender = 'U'
	female  gender = 'F'
	male    gender = 'M'
)

type person struct {
	Subject
	FirstName string      `json:"firstname",bson:"firstname",xml:"person-firstname"`
	LastName  string      `json:"lastname",bson:"lastname",xml:"person-lastname"`
	Birthdate time.Time   `json:"birthdate",bson:"birthdate",xml:"person-birthdate"`
	Deathdate time.Time   `json:"death",bson:"death",xml:"person-death"`
	Gender    gender      `json:"gender",bson:"gender",xml:"person-gender"`
	Weight    Measurement `json:"weight",bson:"weight",xml:"person-weight"`
	Height    Measurement `json:"height",bson:"height",xml:"person-height"`
}

func (p person) GetFirstName() string    { return p.FirstName }
func (p person) GetLastName() string     { return p.LastName }
func (p person) GetBirthdate() time.Time { return p.Birthdate }
func (p person) GetDeathdate() time.Time { return p.Deathdate }
func (p person) GetGender() string       { return p.Gender.String() }
func (p person) GetHeight() Measurement  { return p.Height }
func (p person) GetWeight() Measurement  { return p.Weight }
