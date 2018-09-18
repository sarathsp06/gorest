package repository

//OrderRepository mocks IOrder interface
type OrderRepository struct{}

//New returns an instance of OrderRepository
func New() *OrderRepository {
	return new(OrderRepository)
}
