package category

// 创建分类请求参数类型
type CreateCategoryRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// 创建分类响应参数类型
type CreateCategoryResponse struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// 删除类请求参数类型
type DeleteCategoryRequest struct {
	Seleteditems []string `json:"seleteditems" form:"seleteditems"`
}

// 删除分类响应参数类型
type DeleteCategoryResponse struct {
	Deleteditems []string `json:"deleteditems" form:"deleteditems"`
}
