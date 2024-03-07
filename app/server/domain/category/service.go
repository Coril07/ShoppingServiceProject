package category

import (
	"app/server/utils/csv_helper"
	"app/server/utils/pagination"
	"mime/multipart"
)

type Service struct {
	r Repository
}

// 实例化商品分类service
func NewCategoryService(r Repository) *Service {
	// // 生成表
	// r.Migration()
	// // 插入测试数据
	// r.InsertSampleData()
	return &Service{
		r: r,
	}
}

// 创建分类
func (c *Service) Create(category *Category) error {
	existCity, _ := c.r.GetByName(category.Name)
	if existCity != nil {
		return ErrCategoryExistWithName
	}

	err := c.r.Create(category)
	if err != nil {
		return err
	}

	return nil
}

// 批量创建分类
func (c *Service) BulkCreate(fileHeader *multipart.FileHeader) (int, error) {
	categories := make([]*Category, 0)
	bulkCategory, err := csv_helper.ReadCsv(fileHeader)
	if err != nil {
		return 0, err
	}
	for _, categoryVariables := range bulkCategory {
		categories = append(categories, NewCategory(categoryVariables[0], categoryVariables[1]))
	}
	count, err := c.r.BulkCreate(categories)
	if err != nil {
		return count, err
	}
	return count, nil
}

// 获得分页内分类
func (c *Service) GetAllInPage(page *pagination.Pages) *pagination.Pages {
	categories, count, pageCount := c.r.GetAllInPage(page.Page, page.PageSize)
	page.Items = categories
	page.TotalCount = count
	page.PageCount = pageCount
	return page
}

// 获得所有分类
func (c *Service) GetAll() *pagination.Pages {
	categories, count := c.r.GetAll()
	page := new(pagination.Pages)
	page.Items = categories
	page.TotalCount = count
	return page
}

// 删除分类
func (c *Service) DeleteCategory(name string) error {
	err := c.r.Delete(name)
	return err
}
