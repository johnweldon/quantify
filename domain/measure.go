package domain

import (
	"fmt"
	"strings"
)

type Measure interface {
	fmt.Stringer
	GetName() string
	GetMetrics() []Metric
}

type MeasureRepository interface {
	Repository
	SaveMeasure(x Measure) (string, error)
	GetMeasure(id string) (Measure, error)
	AllMeasures() ([]Measure, error)
}

var _ Measure = (*measure)(nil)

type measure struct {
	Name    string   `json:"name",bson:"name",xml:"measure-name"`
	Metrics []Metric `json:"metrics",bson:"metrics",xml:"measure-metrics"`
}

func NewMeasure(name string, metrics ...Metric) Measure {
	return measure{Name: name, Metrics: metrics}
}

func (m measure) GetName() string      { return m.Name }
func (m measure) GetMetrics() []Metric { return m.Metrics }
func (m measure) String() string       { return fmt.Sprintf("%s [%s]", m.Name, (metrics)(m.Metrics)) }

type metrics []Metric

func (m metrics) String() string {
	r := make([]string, len(m))
	for i, j := range m {
		r[i] = fmt.Sprintf("%s", j)
	}
	return strings.Join(r, ", ")
}
