package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"go.uber.org/dig"

	"goshop/app/api"
)

func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		category *api.Category,
		product *api.Product,
	) error {
		apiV1 := r.Group("api/v1")
		{
			apiV1.GET("/users/:uuid", userService.GetUserByID)
		}
		{
			apiV1.GET("/products", product.GetProducts)
			//apiV1.POST("/products", product.CreateProduct)
			//apiV1.GET("/products/:uuid", product.GetProductByID)
			//apiV1.PUT("/products/:uuid", product.UpdateProduct)
		}
		{
			apiV1.GET("/categories", category.GetCategories)
			apiV1.POST("/categories", category.CreateCategory)
			apiV1.GET("/categories/:uuid", category.GetCategoryByID)
			//apiV1.GET("/categories/:uuid/products", productService.GetProductByCategory)
			apiV1.PUT("/categories/:uuid", category.UpdateCategory)
		}
		{
			apiV1.GET("/warehouses", warehouseService.GetWarehouses)
			apiV1.POST("/warehouses", warehouseService.CreateWarehouse)
			apiV1.GET("/warehouses/:uuid", warehouseService.GetWarehouseByID)
			apiV1.PUT("/warehouses/:uuid", warehouseService.UpdateWarehouse)
		}
		{
			apiV1.GET("/quantities", quantityService.GetQuantities)
			apiV1.POST("/quantities", quantityService.CreateQuantity)
			apiV1.GET("/quantities/:uuid", quantityService.GetQuantityByID)
			apiV1.PUT("/quantities/:uuid", quantityService.UpdateQuantity)
		}
		{
			apiV1.GET("/orders", orderService.GetOrders)
			apiV1.POST("/orders", orderService.CreateOrder)
			apiV1.GET("/orders/:uuid", orderService.GetOrderByID)
			apiV1.PUT("/orders/:uuid", orderService.UpdateOrder)
		}

		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return err
}