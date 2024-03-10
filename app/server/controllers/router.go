package api

import (
	"app/server/config"
	cartApi "app/server/controllers/cart"
	categoryApi "app/server/controllers/category"
	orderApi "app/server/controllers/order"
	productApi "app/server/controllers/product"
	userApi "app/server/controllers/user"
	"app/server/domain/cart"
	"app/server/domain/order"
	"app/server/domain/product"
	"app/server/utils/middleware"
	"fmt"
	"log"
	"net/http"

	"app/server/domain/category"
	"app/server/domain/user"
	"app/server/utils/database_handler"

	"github.com/gin-gonic/gin"
)

// Databases 结构体
type Databases struct {
	categoryRepository    *category.Repository
	userRepository        *user.Repository
	productRepository     *product.Repository
	cartRepository        *cart.Repository
	cartItemRepository    *cart.ItemRepository
	orderRepository       *order.Repository
	orderedItemRepository *order.OrderedItemRepository
}

// 配置文件全局对象
var AppConfig = &config.Configuration{}
var StaticSourcePath = "./client/src/assets"

// 根据配置文件创建数据库
func CreateDBs() *Databases {
	cfgFile := "./server/config/config.yaml"
	conf, err := config.GetAllConfigValues(cfgFile)
	AppConfig = conf
	if err != nil {
		return nil
	}
	if err != nil {
		log.Fatalf("读取配置文件失败. %v", err.Error())
	}
	db := database_handler.NewMySQLDB(AppConfig.DatabaseSettings.DatabaseURI)
	return &Databases{
		categoryRepository:    category.NewCategoryRepository(db),
		cartRepository:        cart.NewCartRepository(db),
		userRepository:        user.NewUserRepository(db),
		productRepository:     product.NewProductRepository(db),
		cartItemRepository:    cart.NewCartItemRepository(db),
		orderRepository:       order.NewOrderRepository(db),
		orderedItemRepository: order.NewOrderedItemRepository(db),
	}
}

// 注册所有控制器
func RegisterHandlers(r *gin.Engine) {

	dbs := *CreateDBs()
	RegisterStaticFiles(r, StaticSourcePath)
	RegisterUserHandlers(r, dbs)
	RegisterCategoryHandlers(r, dbs)
	RegisterCartHandlers(r, dbs)
	RegisterProductHandlers(r, dbs)
	RegisterOrderHandlers(r, dbs)
}

func RegisterStaticFiles(r *gin.Engine, s string) {
	r.StaticFS("/static", http.Dir(s))
}

// 注册分类控制器
func RegisterCategoryHandlers(r *gin.Engine, dbs Databases) {
	categoryService := category.NewCategoryService(*dbs.categoryRepository)
	fmt.Printf("categoryService: %v\n", categoryService)
	categoryController := categoryApi.NewCategoryController(categoryService)
	r.POST(
		"/addcategory", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), categoryController.CreateCategory)
	r.GET("/getallcategories", categoryController.GetAllCategories)
	r.GET("/getcategoriesinpage", categoryController.GetCategoriesInPage)
	r.POST(
		"/uploadcsv", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey),
		categoryController.BulkCreateCategory)
	r.DELETE("/deleteselectedcategories", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), categoryController.DeleteCategory)
}

// 注册用户控制器
func RegisterUserHandlers(r *gin.Engine, dbs Databases) {
	userService := user.NewUserService(*dbs.userRepository)
	userController := userApi.NewUserController(userService, AppConfig)
	userGroup := r.Group("/user")
	userGroup.POST("/sign", userController.CreateUser)
	userGroup.POST("/login", userController.Login)
	userGroup.GET("/cookielogin", middleware.AuthUserMiddleware(AppConfig.SecretKey), userController.CookieLogin)
	userGroup.GET("/log_out", middleware.AuthUserMiddleware(AppConfig.SecretKey), userController.Log_out)

}

// 注册购物车控制器
func RegisterCartHandlers(r *gin.Engine, dbs Databases) {
	cartService := cart.NewService(*dbs.cartRepository, *dbs.cartItemRepository, *dbs.productRepository)
	cartController := cartApi.NewCartController(cartService)
	cartGroup := r.Group("/cart", middleware.AuthUserMiddleware(AppConfig.JwtSettings.SecretKey))
	cartGroup.POST("/item", cartController.AddItem)
	cartGroup.PATCH("/item", cartController.UpdateItem)
	cartGroup.GET("/info", cartController.GetCart)
	cartGroup.DELETE("/delete", cartController.DeleteCart)
}

// 注册商品控制器
func RegisterProductHandlers(r *gin.Engine, dbs Databases) {
	productService := product.NewService(*dbs.productRepository)
	productController := productApi.NewProductController(*productService)
	r.GET("/getallproducts", productController.GetAllProducts)
	r.GET("/getproductsinpage", productController.GetProducts)
	r.GET("/getproductinfo", productController.GetProductInfo)
	r.POST(
		"/createproduct", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.CreateProduct)
	r.POST(
		"/uploadproductimg", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.UploadProductimg)
	r.DELETE(
		"/deleteselectedproducts", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.DeleteProduct)
	// productGroup.PATCH(
	// 	"", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.UpdateProduct)

}

// 注册订单控制器
func RegisterOrderHandlers(r *gin.Engine, dbs Databases) {
	orderService := order.NewService(
		*dbs.orderRepository, *dbs.orderedItemRepository, *dbs.productRepository, *dbs.cartRepository,
		*dbs.cartItemRepository)
	productController := orderApi.NewOrderController(orderService)
	orderGroup := r.Group("/order", middleware.AuthUserMiddleware(AppConfig.JwtSettings.SecretKey))
	orderGroup.POST("/pay", productController.CompleteOrder)
	orderGroup.DELETE("/delete", productController.CancelOrder)
	orderGroup.GET("/get", productController.GetOrders)
}
