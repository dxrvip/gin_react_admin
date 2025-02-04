package routers

import (
	"goVueBlog/api"

	"github.com/gin-gonic/gin"
)

func RegisterUserUrl(r *gin.RouterGroup, userApi *api.UserApi) {
	user := r.Group("/user")
	{
		user.GET("", userApi.List)
		user.POST("", userApi.Register)
		user.GET("/:id", userApi.Info)
		user.DELETE("/:id", userApi.Delete)
		user.PUT("/:id", userApi.Update)
	}

}
