package order

import (
	"app/server/domain/product"
	"app/server/domain/user"

	"gorm.io/gorm"
)

// 订单结构体
type Order struct {
	gorm.Model
	UserID       uint
	User         user.User     `gorm:"foreignKey:ID;references:UserID" json:"-"`
	OrderedItems []OrderedItem `gorm:"foreignKey:OrderID"`
	TotalPrice   float64
	IsCanceled   bool
}

// 订单项结构体
type OrderedItem struct {
	gorm.Model
	Product    product.Product `gorm:"foreignKey:ProductID"`
	ProductID  uint
	Count      int
	OrderID    uint
	IsCanceled bool
}

// 实例化订单
func NewOrder(uid uint, items []OrderedItem) *Order {
	var totalPrice float64 = 0.0
	for _, item := range items {
		totalPrice += item.Product.Price * float64(item.Count)
	}
	return &Order{
		UserID:       uid,
		OrderedItems: items,
		TotalPrice:   totalPrice,
		IsCanceled:   false,
	}
}

// 实例化订单项
func NewOrderedItem(count int, pid uint) *OrderedItem {
	return &OrderedItem{
		Count:      count,
		ProductID:  pid,
		IsCanceled: false,
	}
}
