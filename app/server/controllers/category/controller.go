package category

import (
	"app/server/domain/category"
	"app/server/utils/api_helper"
	"app/server/utils/pagination"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	categoryService *category.Service
}

// 实例化控制器
func NewCategoryController(service *category.Service) *Controller {
	return &Controller{
		categoryService: service,
	}
}

func (c *Controller) CreateCategory(g *gin.Context) {
	var req CreateCategoryRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	newCategory := category.NewCategory(req.Name, req.Desc)
	err := c.categoryService.Create(newCategory)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(
		http.StatusCreated, CreateCategoryResponse{Name: newCategory.Name, Desc: newCategory.Desc})
}

func (c *Controller) BulkCreateCategory(g *gin.Context) {
	fileHeader, err := g.FormFile("file")
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	count, err := c.categoryService.BulkCreate(fileHeader)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(
		http.StatusOK, api_helper.Response{
			Message: fmt.Sprintf(
				"'%s' uploaded! '%d' new categories created", fileHeader.Filename, count)})
}

func (c *Controller) GetCategoriesInPage(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	page = c.categoryService.GetAllInPage(page)
	g.JSON(http.StatusOK, page)
}

func (c *Controller) GetAllCategories(g *gin.Context) {
	page := c.categoryService.GetAll()
	g.JSON(http.StatusOK, page)
}

func (c *Controller) DeleteCategory(g *gin.Context) {
	var req DeleteCategoryRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	var t = []string{}
	for _, v := range req.Seleteditems {
		err := c.categoryService.DeleteCategory(v)
		if err != nil {
			api_helper.HandleError(g, err)
			continue
		}
		t = append(t, v)
		g.JSON(http.StatusOK, DeleteCategoryResponse{Deleteditems: t})
	}
}
