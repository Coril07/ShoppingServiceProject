package order

import (
	"app/server/domain/cart"
	"app/server/domain/product"
	"app/server/utils/pagination"
	"time"
)

var day14ToHours float64 = 336

type Service struct {
	orderRepository       Repository
	orderedItemRepository OrderedItemRepository
	productRepository     product.Repository
	cartRepository        cart.Repository
	cartItemRepository    cart.ItemRepository
}

// 实例化
func NewService(
	orderRepository Repository,
	orderedItemRepository OrderedItemRepository,
	productRepository product.Repository,
	cartRepository cart.Repository,
	cartItemRepository cart.ItemRepository,
) *Service {
	// orderRepository.Migration()
	// orderedItemRepository.Migration()
	return &Service{
		orderRepository:       orderRepository,
		orderedItemRepository: orderedItemRepository,
		productRepository:     productRepository,
		cartRepository:        cartRepository,
		cartItemRepository:    cartItemRepository,
	}

}

// 完成订单
func (c *Service) CompleteOrder(userId uint, pids []int64) error {
	currentCart, err := c.cartRepository.FindOrCreateByUserID(userId)
	if err != nil {
		return err
	}
	cartItems, err := c.cartItemRepository.GetSelectedItemsByPID(currentCart.ID, pids)
	if err != nil {
		return err
	}
	if len(cartItems) == 0 {
		return ErrEmptyCartFound
	}
	orderedItems := make([]OrderedItem, 0)
	for _, item := range cartItems {
		order_item := NewOrderedItem(item.Count, item.ProductID)
		product, err := c.productRepository.FindByID(order_item.ProductID)
		if err != nil {
			return err
		}
		order_item.Product = *product
		orderedItems = append(orderedItems, *order_item)
	}
	err = c.orderRepository.Create(NewOrder(userId, orderedItems))
	return err
}

// 取消订单
func (c *Service) CancelOrder(uid, oid uint) error {
	currentOrder, err := c.orderRepository.FindByOrderID(oid)
	if err != nil {
		return err
	}
	if currentOrder.UserID != uid {
		return ErrInvalidOrderID
	}
	if currentOrder.CreatedAt.Sub(time.Now()).Hours() > day14ToHours {
		return ErrCancelDurationPassed
	}
	currentOrder.IsCanceled = true
	err = c.orderRepository.Update(*currentOrder)

	return err
}

// 获得订单
func (c *Service) GetAll(page *pagination.Pages, uid uint) *pagination.Pages {
	orders, count := c.orderRepository.GetAll(page.Page, page.PageSize, uid)
	page.Items = orders
	page.TotalCount = count
	return page
}