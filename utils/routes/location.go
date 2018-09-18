package routes

import (
	"encoding/json"
	"fmt"
)

// Location stores latitude and longitude of a location
type Location struct {
	Latitude  float64
	Longitude float64
}

//String implements log.Stringer
func (l Location) String() string {
	return fmt.Sprintf("%f,%f", l.Latitude, l.Longitude)
}

// UnmarshalJSON implements json.Unmarshaller
func (l *Location) UnmarshalJSON(data []byte) error {
	var lst = make([]float64, 0)
	if err := json.Unmarshal(data, &lst); err != nil {
		return err
	}
	if len(lst) != 2 {
		return fmt.Errorf("Invalid parameters,len(lst) = %d", len(lst))
	}

	l.Latitude = lst[0]
	l.Longitude = lst[1]
	return nil
}

// IsValid Validates the location
func (l Location) IsValid() bool {
	// TODO : check location boundaries and return validity status
	return true
}
