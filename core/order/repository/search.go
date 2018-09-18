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
	orders, err := search(Order{}, pageSize, pageNum)
	if err != nil {
		return nil, err
	}
	result := make([]order.Order, len(orders))
	for idx, ord := range orders {
		result[idx] = order.Order{
			ID:       ord.ID,
			Status:   ord.Status,
			Distance: ord.Distance,
		}
	}
	return result, nil
}
