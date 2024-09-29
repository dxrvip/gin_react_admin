package routers

import (
	"goVueBlog/api"

	"github.com/gin-gonic/gin"
)

func InitCategory(r *gin.RouterGroup) {
	// 分类
	url := r.Group("/category")
	{
		url.GET("", api.GetCategoryList)

		url.POST("", api.AddCategory)
		url.GET("/:id", api.GetCategoryById)
		url.PUT("/:id", api.UpdateCategory)

		url.DELETE("/:id", api.DeleteCategory)
	}
}
