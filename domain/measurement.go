package quant

import (
	"fmt"
	"strings"
)

type Measurement interface {
	fmt.Stringer
	GetSubject() Subject
	GetMeasurements() []Measure
}

var _ Measurement = (*measurement)(nil)

type measurement struct {
	Event
	Subject      Subject   `json:"subject",bson:"subject",xml:"measurement-subject"`
	Measurements []Measure `json:"measurements",bson:"measurements",xml:"measurement-measurements"`
}

func NewMeasurement(event Event, subject Subject, measurements ...Measure) Measurement {
	return measurement{Event: event, Subject: subject, Measurements: measurements}
}

func (m measurement) GetSubject() Subject        { return m.Subject }
func (m measurement) GetMeasurements() []Measure { return m.Measurements }
func (m measurement) String() string             { return fmt.Sprintf("%s", (measurements)(m.Measurements)) }

type measurements []Measure

func (m measurements) String() string {
	r := make([]string, len(m))
	for i, j := range m {
		r[i] = fmt.Sprintf("%s", j)
	}
	return strings.Join(r, ", ")
}
