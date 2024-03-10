package product

import (
	"app/server/domain/product"
	"app/server/utils/api_helper"
	"app/server/utils/pagination"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	productService product.Service
}

// 实例化
func NewProductController(service product.Service) *Controller {
	return &Controller{
		productService: service,
	}
}

func (c *Controller) GetProducts(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	queryText := g.Query("qt")
	categoryname := g.Query("Cname")
	if categoryname != "" {
		if categoryname == "全部" {
			page = c.productService.GetAllInPage(page)
		} else {
			c.productService.GetSortedProduct(categoryname, page)
		}
	} else {
		if queryText != "" {
			page = c.productService.SearchProduct(queryText, page)
		} else {
			page = c.productService.GetAllInPage(page)

		}
	}

	g.JSON(http.StatusOK, page)

}

func (c *Controller) GetAllProducts(g *gin.Context) {
	page := c.productService.GetAll()
	g.JSON(http.StatusOK, page)
}

func (c *Controller) GetProductInfo(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	cname := g.Query("Pname")
	if cname != "" {
		page, err := c.productService.GetProInfo(cname, page)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		g.JSON(http.StatusOK, page)
	}
}

func (c *Controller) CreateProduct(g *gin.Context) {
	var req CreateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	Price, e := req.Price.Float64()
	Count, e2 := req.Count.Int64()
	if e != nil {
		api_helper.HandleError(g, e)
		return
	}
	if e2 != nil {
		api_helper.HandleError(g, e2)
		return
	}
	err := c.productService.CreateProduct(req.Name, req.Desc, int(Count), Price, req.CategoryName)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, api_helper.Response{
			Message: "Product Created",
		})
}

func (c *Controller) DeleteProduct(g *gin.Context) {
	var req DeleteProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	var t = []string{}
	for _, v := range req.Seleteditems {
		err := c.productService.DeleteProduct(v)
		if err != nil {
			api_helper.HandleError(g, err)
			continue
		}
		t = append(t, v)
	}
	g.JSON(http.StatusOK, DeleteProductResponse{Deleteditems: t})
}

func (c *Controller) UpdateProduct(g *gin.Context) {
	var req UpdateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	err := c.productService.UpdateProduct(req.ToProduct())
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusOK, CreateProductResponse{
			Message: "Product Updated",
		})
}

func (c *Controller) UploadProductimg(g *gin.Context) {
	// var ur UploadRequest
	// err := g.ShouldBind(&ur)
	// if err != nil {
	// 	api_helper.HandleError(g, err)
	// 	return
	// }
	productname := g.PostForm("Name")
	file, _ := g.FormFile("file")
	s := strings.Split(file.Filename, ".")
	suffix := "." + s[len(s)-1]
	err := c.productService.UploadProductimg(productname, suffix)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	err1 := g.SaveUploadedFile(file, "client/src/assets/"+productname+suffix)
	if err1 != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, UploadResponse{Message: "上传成功"})

}
