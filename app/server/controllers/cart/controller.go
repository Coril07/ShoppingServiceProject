package cart

import (
	"app/server/domain/cart"
	"app/server/utils/api_helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	cartService *cart.Service
}

// 实例化
func NewCartController(service *cart.Service) *Controller {
	return &Controller{
		cartService: service,
	}
}

// AddItem godoc
// @Summary 添加Item
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param ItemCartRequest body ItemCartRequest true "product information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart/item [post]
func (c *Controller) AddItem(g *gin.Context) {
	var req ItemCartRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		fmt.Printf("err: %v\n", err)
		return
	}
	val, err0 := req.Count.Int64()
	if err0 != nil {
		api_helper.HandleError(g, err0)
		return
	}
	userId := api_helper.GetUserId(g)
	err := c.cartService.AddItem(userId, req.SKU, int(val))
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(
		http.StatusCreated, CreateCategoryResponse{
			Message: "Item added to cart",
		})
}

// UpdateItem godoc
// @Summary 更新Item
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param ItemCartRequest body ItemCartRequest true "product information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart/item [patch]
func (c *Controller) UpdateItem(g *gin.Context) {
	var req ItemCartRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	val, err0 := req.Count.Int64()
	if err0 != nil {
		api_helper.HandleError(g, err0)
		return
	}
	userId := api_helper.GetUserId(g)
	err := c.cartService.UpdateItem(userId, req.SKU, int(val))
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, CreateCategoryResponse{
			Message: "updated",
		})
}

// GetCart godoc
// @Summary 获得购物车商品列表
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Success 200 {array} cart.Item
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart [get]
func (c *Controller) GetCart(g *gin.Context) {
	userId := api_helper.GetUserId(g)

	result, err := c.cartService.GetCartItems(userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(200, result)
}

func (c *Controller) DeleteCart(g *gin.Context) {
	var req DeleteRequest
	if err := g.ShouldBind(&req); err != nil {
		fmt.Printf("err: %v\n", err)
		api_helper.HandleError(g, err)
		return
	}
	userId := api_helper.GetUserId(g)
	fmt.Printf("req: %v\n", req)
	err1 := c.cartService.DeleteItem(userId, req.SKUs)
	fmt.Printf("err1: %v\n", err1)
	if err1 != nil {
		api_helper.HandleError(g, err1)
		return
	}
	g.JSON(http.StatusOK, "")

}
