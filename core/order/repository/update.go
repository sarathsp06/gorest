package repository

import (
	"context"

	"github.com/sarathsp06/gorest/core/order"
)

// UpdateStatus updates the order with status set to the given status
func (rm OrderRepository) UpdateStatus(
	ctx context.Context,
	ID string,
	Status order.Status,
) error {
	if Status != order.StatusTaken {
		return order.ErrUnavailableOperation
	}
	return rm.Take(ctx, ID)
}

// Take tries to update the order status as taken
func (rm OrderRepository) Take(ctx context.Context, ID string) error {
	//TODO : use transactions or locks
	ord, err := get(ctx, ID)
	if err != nil {
		return err
	}
	if ord.Status == order.StatusTaken {
		return order.ErrAlreadyTaken
	}
	if _, err = update(ctx, ID, &Order{Status: order.StatusTaken}); err != nil {
		return err
	}
	return nil
}
