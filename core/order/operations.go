package order

import (
	"context"

	"github.com/sarathsp06/gorest/utils/metrics"
	"github.com/sarathsp06/gorest/utils/routes"
)

// Create creates an order
func Create(ctx context.Context, origin, destination routes.Location) (*Order, error) {
	metrics.CaptureDelay("CoreCreate")()
	return order.Create(ctx, origin, destination)
}

// UpdateStatus updates the status of the order
func UpdateStatus(ctx context.Context, ID string, status Status) error {
	metrics.CaptureDelay("CoreUpdateStatus")()
	return order.UpdateStatus(ctx, ID, status)
}

// List lists all the orders
func List(ctx context.Context, page int, limit int) ([]Order, error) {
	metrics.CaptureDelay("CoreList")()
	return order.List(ctx, page, limit)
}
