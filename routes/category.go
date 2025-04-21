package routes

import (
	"goVueBlog/api"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryUrls(r *gin.RouterGroup) {

	categoryApi := api.NewCategoryApi()

	category := r.Group("/category")
	{
		category.GET("", categoryApi.GetCategoryList)

		category.POST("", categoryApi.AddCategory)
		category.GET("/:id", categoryApi.GetCategoryById)
		category.PUT("/:id", categoryApi.UpdateCategory)

		category.DELETE("/:id", categoryApi.DeleteCategory)
	}

}
