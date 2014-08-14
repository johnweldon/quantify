package domain

import (
	"fmt"
	"time"
)

type EventRepository interface {
	Repository
	SaveEvent(event Event) (string, error)
	GetEvent(id string) (Event, error)
	AllEvents() ([]Event, error)
}

type Event struct {
	Time        time.Time `json:"time",bson:"time",xml:"event-time"`
	Location    Location  `json:"location",bson:"location",xml:"event-location"`
	Name        string    `json:"name",bson:"name",xml:"event-name"`
	Description string    `json:"description",bson:"description",xml:"event-description"`
}

func NewEvent(name, description string, time time.Time, location Location) Event {
	return Event{Time: time, Location: location, Name: name, Description: description}
}

func (e Event) GetName() string        { return e.Name }
func (e Event) GetTime() time.Time     { return e.Time }
func (e Event) GetLocation() Location  { return e.Location }
func (e Event) GetDescription() string { return e.Description }
func (e Event) String() string         { return fmt.Sprintf("%s, %s, %s", e.Name, e.Time, e.Location) }
