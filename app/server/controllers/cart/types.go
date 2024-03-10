package cart

import "encoding/json"

// item请求参数
type ItemCartRequest struct {
	SKU   string      `json:"sku"`
	Count json.Number `json:"count"`
}

// 创建分类响应
type CreateCategoryResponse struct {
	Message string `json:"message"`
}

// DeleteItem请求参数
type DeleteRequest struct {
	SKUs []string `json:"skus" form:"skus"`
}

// DeleteItem响应
type DeleteResponse struct {
	Message string `json:"message"`
}
