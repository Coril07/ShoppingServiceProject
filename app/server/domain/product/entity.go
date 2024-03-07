package product

import (
	"app/server/domain/category"

	"gorm.io/gorm"
)

// 商品结构体
type Product struct {
	gorm.Model
	Name         string
	SKU          string
	Desc         string
	StockCount   int
	Price        float64
	Url          string
	CategoryName string            `gorm:"type:varchar(191)"`
	Category     category.Category `gorm:"references:Name;foreignKey:CategoryName" json:"-"` // 分类
	IsDeleted    bool
}

// 商品结构体实例
func NewProduct(name string, desc string, stockCount int, price float64, cname string) *Product {
	return &Product{
		Name:         name,
		Desc:         desc,
		StockCount:   stockCount,
		Price:        price,
		CategoryName: cname,
		IsDeleted:    false,
	}
}
