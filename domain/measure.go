package domain

import (
	"fmt"
	"strings"
)

type MeasureRepository interface {
	Repository
	SaveMeasure(x Measure) (string, error)
	GetMeasure(id string) (Measure, error)
	AllMeasures() ([]Measure, error)
}

type Measure struct {
	Name    string   `json:"name",bson:"name",xml:"measure-name"`
	Metrics []Metric `json:"metrics",bson:"metrics",xml:"measure-metrics"`
}

func NewMeasure(name string, metrics ...Metric) Measure {
	return Measure{Name: name, Metrics: metrics}
}

func (m Measure) GetName() string      { return m.Name }
func (m Measure) GetMetrics() []Metric { return m.Metrics }
func (m Measure) String() string       { return fmt.Sprintf("%s [%s]", m.Name, (metrics)(m.Metrics)) }

type metrics []Metric

func (m metrics) String() string {
	r := make([]string, len(m))
	for i, j := range m {
		r[i] = fmt.Sprintf("%s", j)
	}
	return strings.Join(r, ", ")
}
