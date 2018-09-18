package repository

import (
	"context"

	"github.com/sarathsp06/gorest/core/order"
	"github.com/sarathsp06/gorest/utils/routes"
)

// Create adds items
func (rm OrderRepository) Create(
	ctx context.Context,
	origin routes.Location,
	destination routes.Location,
) (
	*order.Order,
	error,
) {
	distance, err := routes.GetDistance(ctx, origin, destination)
	if err != nil {
		switch err {
		case routes.ErrInvalidLocation:
			return nil, order.ErrInvalidLocation
		case routes.ErrNotFound:
			return nil, order.ErrInvalidLocation
		default:
			return nil, err
		}
	}

	ord := Order{
		Distance: distance,
		Status:   order.StatusUnAssign,
	}

	if err = create(&ord); err != nil {
		return nil, err
	}

	result := &order.Order{
		ID:       ord.ID,
		Distance: ord.Distance,
		Status:   ord.Status,
	}
	return result, nil
}
