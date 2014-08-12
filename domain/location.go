package domain

import (
	"fmt"
)

type Location interface {
	fmt.Stringer
	GetName() string
	GetAltitude() Metric
	GetLatLong() (Degree, Degree)
	GetPosition() string
}

type LocationRepository interface {
	Repository
	SaveLocation(location Location) (string, error)
	GetLocation(id string) (Location, error)
	AllLocations() ([]Location, error)
}

var _ Location = (*location)(nil)

type Degree struct {
	Int      int8  `json:"int",bson:"int",xml:"degree-int"`
	Fraction int64 `json:"fraction",bson:"fraction",xml:"degree-fraction"`
}

func (d Degree) String() string { return fmt.Sprintf("%d.%d", d.Int, d.Fraction) }

type location struct {
	Name      string `json:"name",bson:"name",xml:"location-name"`
	Latitude  Degree `json:"lat",bson:"lat",xml:"location-lat"`
	Longitude Degree `json:"lon",bson:"lon",xml:"location-lon"`
	Altitude  Metric `json:"altitude",bson:"altitude",xml:"location-altitude"`
}

func NewLocation(name string, latitude, longitude Degree, altitude Metric) Location {
	return location{Name: name, Latitude: latitude, Longitude: longitude, Altitude: altitude}
}

func (l location) GetName() string              { return l.Name }
func (l location) GetAltitude() Metric          { return l.Altitude }
func (l location) GetLatLong() (Degree, Degree) { return l.Latitude, l.Longitude }
func (l location) GetPosition() string          { return fmt.Sprintf("%s,%s", l.Latitude, l.Longitude) }
func (l location) String() string               { return fmt.Sprintf("%s (%s)", l.GetPosition(), l.Name) }
