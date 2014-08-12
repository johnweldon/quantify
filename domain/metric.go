package domain

import (
	"fmt"
)

type Metric interface {
	fmt.Stringer
	GetType() string
	GetSI() string
	GetAmount() int64
}

type MetricRepository interface {
	Repository
	SaveMetric(x Metric) (string, error)
	GetMetric(id string) (Metric, error)
	AllMetrics() ([]Metric, error)
}

var _ Metric = (*metric)(nil)

type metric struct {
	Type   string `json:"type",bson:"type",xml:"metric-type"`
	SI     string `json:"si",bson:"si",xml:"metric-si"`
	Amount int64  `json:"amount",bson:"amount",xml:"metric-amount"`
}

func NewMetric(metrictype, si string, amount int64) Metric {
	return metric{Type: metrictype, Amount: amount}
}

func (m metric) String() string   { return fmt.Sprintf("%d %s", m.Amount, m.SI) }
func (m metric) GetType() string  { return m.Type }
func (m metric) GetSI() string    { return m.SI }
func (m metric) GetAmount() int64 { return m.Amount }
