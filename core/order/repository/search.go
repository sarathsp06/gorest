package repository

import (
	"context"

	"github.com/sarathsp06/gorest/core/order"
)

// List lists all the orders given page number and limit
func (rm OrderRepository) List(
	ctx context.Context,
	pageSize int,
	pageNum int,
) (
	[]order.Order,
	error,
) {
	return nil, nil
}
