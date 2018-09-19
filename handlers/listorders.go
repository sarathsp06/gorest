package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sarathsp06/gorest/core/order"
	"github.com/sarathsp06/gorest/utils/metrics"
)

const (
	defaultPageNum  = 1
	defaultPageSize = 20
)

func getPageInfo(ctx echo.Context) (int, int) {
	var pageSize, pageNum int64
	// TODO: May be move this to a middle ware
	pageNumStr, pageSizeStr := ctx.QueryParam("page"), ctx.QueryParam("limit")
	if pageSizeStr != "" {
		pageSize, _ = strconv.ParseInt(pageSizeStr, 10, 64)
	}
	if pageNumStr != "" {
		pageNum, _ = strconv.ParseInt(pageNumStr, 10, 64)
	}

	if pageNum <= 1 {
		pageNum = defaultPageNum
	}
	if pageSize <= 1 {
		pageSize = defaultPageSize
	}
	return int(pageNum), int(pageSize)
}

// ListOrders lists all the orders
func ListOrders(ctx echo.Context) error {
	metrics.CaptureDelay("ListOrders")()
	page, limit := getPageInfo(ctx)
	orders, err := order.List(ctx.Request().Context(), page, limit)
	if err != nil {
		log.Println("Error listing orders ", err)
		// check for error and return appropriately
		return ctx.JSON(ErrInternalServerError(err.Error()))
	}
	return ctx.JSON(http.StatusOK, orders)
}
