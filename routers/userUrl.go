package routers

import (
	"goVueBlog/api"

	"github.com/gin-gonic/gin"
)

func UserUrls(v1 *gin.RouterGroup) {
	user := v1.Group("/user")
	user.GET("/info", api.Info)
	user.DELETE("/:id", api.Delete)
	user.PUT("/:id", api.Update)
}
