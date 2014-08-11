package quant

import (
	"fmt"
	"time"
)

type Event interface {
	fmt.Stringer
	GetName() string
	GetTime() time.Time
	GetLocation() Location
	GetDescription() string
}

var _ Event = (*event)(nil)

type event struct {
	Time        time.Time `json:"time",bson:"time",xml:"event-time"`
	Location    Location  `json:"location",bson:"location",xml:"event-location"`
	Name        string    `json:"name",bson:"name",xml:"event-name"`
	Description string    `json:"description",bson:"description",xml:"event-description"`
}

func NewEvent(name, description string, time time.Time, location Location) Event {
	return event{Time: time, Location: location, Name: name, Description: description}
}

func (e event) GetName() string        { return e.Name }
func (e event) GetTime() time.Time     { return e.Time }
func (e event) GetLocation() Location  { return e.Location }
func (e event) GetDescription() string { return e.Description }
func (e event) String() string         { return fmt.Sprintf("%s, %s, %s", e.Name, e.Time, e.Location) }
