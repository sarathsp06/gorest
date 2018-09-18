package routes

import (
	"context"
)

// Interface declares all the functions to be exposed for recipe service
type Interface interface {
	GetDistance(ctx context.Context, origin, destination Location) (float64, error)
}

// order implements the interface
var route Interface

// SetDefault sets default IRecipe executer
func SetDefault(routeImpl Interface) {
	route = routeImpl
}
