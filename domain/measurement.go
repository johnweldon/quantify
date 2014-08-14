package domain

import (
	"fmt"
	"strings"
)

type MeasurementRepository interface {
	Repository
	SaveMeasurement(x Measurement) (string, error)
	GetMeasurement(id string) (Measurement, error)
	AllMeasurements() ([]Measurement, error)
}

type Measurement struct {
	Event
	Subject      Subject   `json:"subject",bson:"subject",xml:"measurement-subject"`
	Measurements []Measure `json:"measurements",bson:"measurements",xml:"measurement-measurements"`
}

func NewMeasurement(event Event, subject Subject, measurements ...Measure) Measurement {
	return Measurement{Event: event, Subject: subject, Measurements: measurements}
}

func (m Measurement) GetSubject() Subject        { return m.Subject }
func (m Measurement) GetMeasurements() []Measure { return m.Measurements }
func (m Measurement) String() string             { return fmt.Sprintf("%s", (measurements)(m.Measurements)) }

type measurements []Measure

func (m measurements) String() string {
	r := make([]string, len(m))
	for i, j := range m {
		r[i] = fmt.Sprintf("%s", j)
	}
	return strings.Join(r, ", ")
}
