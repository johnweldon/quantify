package domain

import (
	"fmt"
)

type MetricRepository interface {
	Repository
	SaveMetric(x Metric) (string, error)
	GetMetric(id string) (Metric, error)
	AllMetrics() ([]Metric, error)
}

type Metric struct {
	Type   string `json:"type",bson:"type",xml:"metric-type"`
	SI     string `json:"si",bson:"si",xml:"metric-si"`
	Amount int64  `json:"amount",bson:"amount",xml:"metric-amount"`
}

func NewMetric(metrictype, si string, amount int64) Metric {
	return Metric{Type: metrictype, Amount: amount}
}

func (m Metric) String() string   { return fmt.Sprintf("%d %s", m.Amount, m.SI) }
func (m Metric) GetType() string  { return m.Type }
func (m Metric) GetSI() string    { return m.SI }
func (m Metric) GetAmount() int64 { return m.Amount }
