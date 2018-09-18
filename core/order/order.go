package order

// Order represents an order
type Order struct {
	ID       string  `json:"id"`
	Distance float64 `json:"distance"`
	Status   Status  `json:"status"`
}

// IsValid checks if the status is one of the valid statuses
func (status Status) IsValid() bool {
	switch status {
	case StatusTaken, StatusUnAssign:
		return true
	}
	return false
}

// Status represents an order status
type Status string

const (
	// StatusUnAssign orders yet to be assigned
	StatusUnAssign = "UNASSIGN"
	// StatusTaken orders taken
	StatusTaken = "taken"
)
