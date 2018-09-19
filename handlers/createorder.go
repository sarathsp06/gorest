package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sarathsp06/gorest/core/order"
	"github.com/sarathsp06/gorest/utils/metrics"
	"github.com/sarathsp06/gorest/utils/routes"
)

// CreateOrderRequest stores a request to create order
type CreateOrderRequest struct {
	Origin      routes.Location `json:"origin"`
	Destination routes.Location `json:"destination"`
}

// Validate Checks for  the Validity of CreateOrderReque
func (cro CreateOrderRequest) Validate() (int, *Error) {
	if (cro.Origin == routes.Location{}) ||
		(cro.Destination == routes.Location{}) {
		return ErrBadRequestParametersMissing("origin and destination are mandatory")
	}

	if !cro.Origin.IsValid() || !cro.Destination.IsValid() {
		return ErrBadRequestInvalidBody("Origin or destination not in proper format")
	}
	return 0, nil
}

// CreateOrder handles create order http requets
func CreateOrder(ctx echo.Context) error {
	metrics.CaptureDelay("CreateOrderHandler")()
	var req CreateOrderRequest
	if err := ctx.Bind(&req); err != nil {
		log.Println("Error binding request", err)
		return ctx.JSON(ErrBadRequestInvalidBody())
	}
	//validate request
	if httpCode, err := req.Validate(); err != nil {
		return ctx.JSON(httpCode, err)
	}

	// execute
	ord, err := order.Create(ctx.Request().Context(), req.Origin, req.Destination)
	if err != nil {
		log.Println("Error creating order ", err)
		// check for error and return appropriately
		return ctx.JSON(ErrInternalServerError(err.Error()))
	}
	return ctx.JSON(http.StatusOK, ord)
}
