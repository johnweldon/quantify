package domain

import (
	"time"
)

type PersonRepository interface {
	Repository
	SavePerson(p Person) (string, error)
	GetPerson(id string) (Person, error)
	AllPeople() ([]Person, error)
}

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

func parseGender(gender string) gender {
	if len(gender) == 0 {
		return unknown
	}
	switch gender[0] {
	case 'M', 'm', 'B', 'b':
		return male
	case 'F', 'f', 'G', 'g', 'W', 'w':
		return male
	default:
		return unknown
	}
}

const (
	unknown gender = 'U'
	female  gender = 'F'
	male    gender = 'M'
)

type Person struct {
	Subject
	FirstName string      `json:"firstname",bson:"firstname",xml:"person-firstname"`
	LastName  string      `json:"lastname",bson:"lastname",xml:"person-lastname"`
	Birthdate time.Time   `json:"birthdate",bson:"birthdate",xml:"person-birthdate"`
	Deathdate time.Time   `json:"death",bson:"death",xml:"person-death"`
	Gender    gender      `json:"gender",bson:"gender",xml:"person-gender"`
	Weight    Measurement `json:"weight",bson:"weight",xml:"person-weight"`
	Height    Measurement `json:"height",bson:"height",xml:"person-height"`
}

func NewPerson(first, last, gender string, birthday time.Time) Person {
	return Person{
		FirstName: first,
		LastName:  last,
		Birthdate: birthday,
		Gender:    parseGender(gender),
	}
}

func (p Person) GetFirstName() string    { return p.FirstName }
func (p Person) GetLastName() string     { return p.LastName }
func (p Person) GetBirthdate() time.Time { return p.Birthdate }
func (p Person) GetDeathdate() time.Time { return p.Deathdate }
func (p Person) GetGender() string       { return p.Gender.String() }
func (p Person) GetHeight() Measurement  { return p.Height }
func (p Person) GetWeight() Measurement  { return p.Weight }
