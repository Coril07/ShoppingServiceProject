package product

import (
	"app/server/utils/pagination"
)

type Service struct {
	productRepository Repository
}

// 实例化
func NewService(productRepository Repository) *Service {
	// productRepository.Migration()
	return &Service{
		productRepository: productRepository,
	}

}

// 获得商品分页
func (c *Service) GetAllInPage(page *pagination.Pages) *pagination.Pages {
	products, count, pageCount := c.productRepository.GetAllInPage(page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	page.PageCount = pageCount
	return page

}

// 获得所有商品
func (c *Service) GetAll() *pagination.Pages {
	products, count := c.productRepository.GetAll()
	page := new(pagination.Pages)
	page.Items = products
	page.TotalCount = count
	return page

}

// 创建商品
func (c *Service) CreateProduct(name string, desc string, count int, price float64, cname string) error {
	newProduct := NewProduct(name, desc, count, price, cname)
	err := c.productRepository.Create(newProduct)
	return err
}

// 删除商品
func (c *Service) DeleteProduct(sku string) error {
	err := c.productRepository.Delete(sku)
	return err
}

// 更新商品
func (c *Service) UpdateProduct(product *Product) error {
	err := c.productRepository.Update(*product)
	return err
}

// 上传商品照片
func (c *Service) UploadProductimg(pname string, suffix string) error {
	Product, err := c.productRepository.FindByName(pname)
	if err != nil {
		return err
	}
	Product.Url = "http://localhost:8080/static/" + pname + suffix
	err1 := c.productRepository.Update(*Product)
	if err1 != nil {
		return err1
	}
	return nil
}

// 查找商品
func (c *Service) SearchProduct(text string, page *pagination.Pages) *pagination.Pages {
	products, count := c.productRepository.SearchByString(text, page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}

// 查找某类商品
func (c *Service) GetSortedProduct(Cname string, page *pagination.Pages) *pagination.Pages {
	products, count := c.productRepository.SearchByCategory(Cname, page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}

func (c *Service) GetProInfo(Cname string, page *pagination.Pages) (*pagination.Pages, error) {
	product, err := c.productRepository.FindByName(Cname)
	page.Items = product
	return page, err
}
