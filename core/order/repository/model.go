package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/sarathsp06/gorest/core/order"
	"github.com/sarathsp06/gorest/db/crud"
	"github.com/sarathsp06/gorest/utils"
)

const itemName = "order"

// Order is the model for order
// TODO: the `bson` makes the model little bit coupled to the mongo db
// Change it if required later
type Order struct {
	ID          string       `bson:"_id"`
	Distance    float64      `bson:"distance"`
	Status      order.Status `bson:"status"`
	DateCreated *time.Time   `bson:"date_created"`
	DateUpdated *time.Time   `bson:"date_updated"`
}

func create(order *Order) error {
	if order.ID == "" {
		order.ID = utils.GetUniqueID()
	}
	dateCreated := time.Now()
	order.DateCreated = &dateCreated
	if err := crud.Insert(itemName, order); err != nil {
		return fmt.Errorf("Error saving Order : %s", err.Error())
	}
	return nil
}

func update(ctx context.Context, ID string, fields *Order) (*Order, error) {
	dateUpdated := time.Now()
	fields.DateUpdated = &dateUpdated
	var ord Order
	if err := crud.Update(itemName, ID, fields, &ord); err != nil {
		if err == crud.ErrNotFound {
			return nil, order.ErrNotFound
		}
		return nil, fmt.Errorf("Error deleting order : %s", err.Error())
	}
	return &ord, nil
}

func search(filter Order, pageNum, pageSize int) ([]Order, error) {
	orders := []Order{}
	if err := crud.Get(itemName, filter, &orders, pageNum, pageSize); err != nil {
		if err == crud.ErrNotFound {
			return nil, order.ErrNotFound
		}
		return nil, fmt.Errorf("Error searching for order: %s", err.Error())
	}
	if len(orders) == 0 {
		return nil, order.ErrNotFound
	}
	return orders, nil
}

func get(ctx context.Context, ID string) (*Order, error) {
	orders := []Order{}
	if err := crud.Get(itemName, Order{ID: ID}, &orders, 1, 1); err != nil {
		if err == crud.ErrNotFound {
			return nil, order.ErrNotFound
		}
		return nil, fmt.Errorf("Error searching for order: %s", err.Error())
	}
	if len(orders) == 0 {
		return nil, order.ErrNotFound
	}
	return &orders[0], nil
}
