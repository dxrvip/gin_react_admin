package routes

import (
	"goVueBlog/api"
	"goVueBlog/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterBlogUrls(r *gin.RouterGroup) {

	apis := api.NewArticleApi()
	url := r.Group("/article")
	{
		url.GET("", middleware.RoleMiddleware("hahaha"), apis.ArticleList)
		url.PUT("/:id", apis.ArticleUpdate)
		url.GET("/:id", apis.ArticleInfo)

		url.DELETE("/:id", apis.ArticleDelete)
		url.POST("", apis.CreateArticle)
	}
}
