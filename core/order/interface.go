package order

import (
	"context"

	"github.com/sarathsp06/gorest/utils/routes"
)

// Interface declares all the functions to be exposed for recipe service
type Interface interface {
	Create(ctx context.Context, origin, destination routes.Location) (*Order, error)
	UpdateStatus(ctx context.Context, ID string, status Status) error
	List(ctx context.Context, page int, limit int) ([]Order, error)
}

// order implements the interface
var order Interface

// SetExecuter sets default IRecipe executer
func SetExecuter(orderImpl Interface) {
	order = orderImpl
}
