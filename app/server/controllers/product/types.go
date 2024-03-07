package product

import (
	"app/server/domain/product"
	"encoding/json"
)

// 创建商品请求参数
type CreateProductRequest struct {
	Name         string      `json:"name"`
	Desc         string      `json:"desc"`
	Price        json.Number `json:"price"`
	Count        json.Number `json:"count"`
	CategoryName string      `json:"categoryName"`
}

// 创建商品响应参数
type CreateProductResponse struct {
	Message string `json:"message"`
}

// 删除商品请求参数类型
type DeleteProductRequest struct {
	Seleteditems []string `json:"seleteditems" form:"seleteditems"`
}

// 删除商品响应参数类型
type DeleteProductResponse struct {
	Deleteditems []string `json:"deleteditems" form:"deleteditems"`
}

// 更新商品请求参数
type UpdateProductRequest struct {
	SKU          string  `json:"sku"`
	Name         string  `json:"name"`
	Desc         string  `json:"desc"`
	Price        float64 `json:"price"`
	Count        int     `json:"count"`
	CategoryName string  `json:"categoryName"`
}

// 类型转换 UpdateProductRequest to Product
func (p *UpdateProductRequest) ToProduct() *product.Product {
	return product.NewProduct(p.Name, p.Desc, p.Count, p.Price, p.CategoryName)
}

// 上传请求
type UploadRequest struct {
	Productname string `json:"Name"`
}

// 上传相应
type UploadResponse struct {
	Message string `json:"message"`
}

// 商品详情
type ProductInfo struct {
	SKU          string  `json:"sku"`
	Name         string  `json:"name"`
	Desc         string  `json:"desc"`
	Price        float64 `json:"price"`
	Count        int     `json:"count"`
	CategoryName string  `json:"categoryName"`
	StockCount   int     `json:"stockcount"`
	Url          string  `json:"url" `
	IsDeleted    bool    `json:"isdeleted"`
}
