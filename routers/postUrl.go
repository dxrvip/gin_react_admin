package routers

import (
	"goVueBlog/api"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	urls := r.Group("/posts")
	{
		// 添加路由
		urls.GET("", api.PostList)
		urls.PUT("/:id", api.PostUpdate)
		urls.GET("/:id", api.PostInfo)

		urls.DELETE("/:id", api.PostDelete)
		urls.POST("", api.CreatePost)
	}
}
