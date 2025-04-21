package routes

import (
	"goVueBlog/api"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.RouterGroup) {

	brand := r.Group("/brand")
	brandApi := api.NewBrandApi()
	brand.GET("", brandApi.List)
	brand.POST("", brandApi.Create)
	brand.GET("/:id", brandApi.InfoById)
	brand.PUT("/:id", brandApi.Update)
	brand.DELETE("/:id", brandApi.Delete)
	// 商品分类
	categories := r.Group("/categories")
	categoriesApi := api.NewCategoriesServiceApi()
	categories.GET("", categoriesApi.List)
	categories.POST("", categoriesApi.Create)
	categories.GET("/:id", categoriesApi.InfoById)
	categories.PUT("/:id", categoriesApi.Update)
	categories.DELETE("/:id", categoriesApi.Del)
	// 属性
	attributeKey := r.Group("/attribute")
	attributeKeyApi := api.NewAttributeApi()
	attributeKey.GET("/:id", attributeKeyApi.InfoAttribute)
	attributeKey.PUT("/:id", attributeKeyApi.UpdateAttribute)
	attributeKey.POST("", attributeKeyApi.CreateAttribute)
	attributeKey.GET("", attributeKeyApi.ListAttribute)
	attributeKey.DELETE("/:id", attributeKeyApi.DelAttribute)

	// 商品
	product := r.Group("/product")
	productApi := api.NewProductApi()
	product.GET("", productApi.ProductList)
	product.POST("", productApi.ProductCreate)
	product.GET("/:id", productApi.ProductInfo)
	product.PUT("/:id", productApi.ProductUpdate)
	product.DELETE("/:id", productApi.ProductDelete)

	// 二手商品SKU
	secondHandSku := r.Group("/secondHandSkus")
	secondHandSkuApi := api.NewSecondHandSkuApi()
	secondHandSku.GET("", secondHandSkuApi.ListSecondHandSkus)
	secondHandSku.POST("", secondHandSkuApi.CreateSecondHandSku)
	secondHandSku.GET("/:id", secondHandSkuApi.GetSecondHandSku)
	secondHandSku.PUT("/:id", secondHandSkuApi.UpdateSecondHandSku)
	secondHandSku.DELETE("/:id", secondHandSkuApi.DeleteSecondHandSku)

	// 获取商品的二手SKU列表
	// product.GET("/:id/secondHandSkus", secondHandSkuApi.GetProductSecondHandSkus)

	// 订单
	order := r.Group("/order")
	orderApi := api.NewOrderApi()
	order.GET("", orderApi.GetUserOrders)
	order.POST("", orderApi.CreateOrder)
	order.GET("/:id", orderApi.GetOrderDetail)
	order.PUT("/:id", orderApi.UpdateOrder)
	order.DELETE("/:id", orderApi.DelteOrder)

	// sku
	// sku := r.Group("/productsku")
	// skuApi := api.NewSecondHandSkuApi()
	// sku.GET("", skuApi.ListSecondHandSkus)
	// sku.POST("", skuApi.CreateSecondHandSku)
	// sku.GET("/:id", skuApi.GetProductSecondHandSkus)
	// sku.PUT("/:id", skuApi.UpdateSecondHandSku)
	// sku.DELETE("/:id", skuApi.DeleteSecondHandSku)

}
