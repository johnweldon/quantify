package quant

import (
	"fmt"
)

type Subject interface {
	fmt.Stringer
	GetName() string
}

var _ Subject = (*subject)(nil)

type subject struct {
	Name string `json:"name",bson:"name",xml:"subject-name"`
}

func NewSubject(name string) Subject {
	return subject{Name: name}
}

func (s subject) GetName() string { return s.Name }
func (s subject) String() string  { return s.Name }
