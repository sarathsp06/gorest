package order

import "errors"

var (
	// ErrAlreadyTaken error is thrown upon trying to take an already taken order
	ErrAlreadyTaken = errors.New("ORDER_ALREADY_BEEN_TAKEN")
	// ErrNotFound error thrown when trying to operate on non existing order
	ErrNotFound = errors.New("ORDER_NOT_FOUND")
	// ErrInvalidLocation thrown when the location validatin fails
	ErrInvalidLocation = errors.New("LOCATION_INVALID")
	// ErrNoRoutes error is thrown when orign and destination does not have a route
	ErrNoRoutes = errors.New("NO_ROUTES_FOUND")
	// ErrUnavailableOperation can not udate state to anything other than "taken"
	ErrUnavailableOperation = errors.New("UNSUPPORTED_STATUS_TRANSITION")
)
