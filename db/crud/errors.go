package crud

import "errors"

var (
	//ErrNotFound thrown if requested item is not found
	ErrNotFound = errors.New("Record not Found")
)
