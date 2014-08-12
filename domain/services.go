package domain

import (
	"time"
)

type MeasurementService interface {
	AddMeasurement(name string, measures ...Measure) (Entity, error)
	GetMeasurements(name string) ([]Measurement, error)
}

type PeopleService interface {
	AddPerson(first, last, gender string, birthdate time.Time) (Person, error)
	GetPeople() ([]Person, error)
}
