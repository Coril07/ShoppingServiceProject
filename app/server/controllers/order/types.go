package order

import "encoding/json"

// 完成订单请求
type CompleteOrderRequest struct {
	PIDs []json.Number `json:"PIDs"`
}

// 取消订单
type CancelOrderRequest struct {
	OrderIds []json.Number `json:"orderIds" form:"orderIds"`
}
