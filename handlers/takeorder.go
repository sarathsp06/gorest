package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sarathsp06/gorest/core/order"
)

// UpdateOrderRequest stores a request to update an order
type UpdateOrderRequest struct {
	Status order.Status `json:"status"`
}

// UpdateOrder handles take order http requets
func UpdateOrder(ctx echo.Context) error {
	var req UpdateOrderRequest
	if err := ctx.Bind(&req); err != nil {
		log.Println("Error binding request", err)
		return ctx.JSON(ErrBadRequestInvalidBody())
	}

	if string(req.Status) == "" {
		return ctx.JSON(ErrBadRequestParametersMissing("non empty status is mandatory"))
	}

	if !req.Status.IsValid() {
		return ctx.JSON(ErrBadRequestInvalidParameter("status invalid"))
	}

	ID := ctx.Param("id")
	reqCtx := ctx.Request().Context()

	if err := order.UpdateStatus(reqCtx, ID, req.Status); err != nil {
		switch err {
		case order.ErrNotFound:
			return ctx.JSON(ErrNotFound(err.Error()))
		case order.ErrAlreadyTaken:
			return ctx.JSON(ErrResourceConflict(err.Error()))
		case order.ErrUnavailableOperation:
			return ctx.JSON(ErrInvalidOperation(err.Error()))
		}
		log.Println("Error taking order ", err)
		// check for error and return appropriately
		return ctx.JSON(ErrInternalServerError(err.Error()))
	}
	return ctx.JSON(http.StatusOK, map[string]string{"status": "SUCCESS"})
}
