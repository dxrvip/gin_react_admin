package routes

import (
	"goVueBlog/api"

	"github.com/gin-gonic/gin"
)

func RegisterSystemUrls(r *gin.RouterGroup) {

	systemApi := api.NewSystemMenuApi()
	systemUrls := r.Group("/systemMenu")
	{
		systemUrls.GET("", systemApi.SystemMenuList)

	}
}
