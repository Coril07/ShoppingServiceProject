package order

import (
	"app/server/domain/order"
	"app/server/utils/api_helper"
	"app/server/utils/pagination"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	orderService *order.Service
}

// 实例订单
func NewOrderController(orderService *order.Service) *Controller {
	return &Controller{
		orderService: orderService,
	}
}

func (c *Controller) CompleteOrder(g *gin.Context) {
	var req CompleteOrderRequest
	if err := g.ShouldBind(&req); err != nil {
		fmt.Printf("err: %v\n", err)
		api_helper.HandleError(g, err)
		return
	}
	userId := api_helper.GetUserId(g)
	var pids = []int64{}
	for i := 0; i < len(req.PIDs); i++ {
		val, err := req.PIDs[i].Int64()
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		pids = append(pids, val)
	}
	err := c.orderService.CompleteOrder(userId, pids)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, api_helper.Response{
			Message: "Order Created",
		})
}

func (c *Controller) CancelOrder(g *gin.Context) {
	var req CancelOrderRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	userId := api_helper.GetUserId(g)
	err := c.orderService.CancelOrder(userId, req.OrderId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, api_helper.Response{
			Message: "Order Canceled",
		})
}

func (c *Controller) GetOrders(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	userId := api_helper.GetUserId(g)
	page = c.orderService.GetAll(page, userId)
	g.JSON(http.StatusOK, page)
}
