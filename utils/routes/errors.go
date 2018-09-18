package routes

import "errors"

var (
	// ErrNotFound error thrown when there is no possible routes between two locations
	ErrNotFound = errors.New("NO_ROUTES_FOUND")
	// ErrInvalidLocation thrown when the location validatin fails
	ErrInvalidLocation = errors.New("LOCATION_INVALID")
)
