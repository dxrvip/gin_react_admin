package routers

import (
	"goVueBlog/api"

	"github.com/gin-gonic/gin"
)

func RegisterRoleUrls(r *gin.RouterGroup) {
	roleApi := api.NewRoleApi()
	roleUrl := r.Group("/role")
	{
		roleUrl.GET("", roleApi.ListRole)
		roleUrl.GET("/:id", roleApi.GetRoleById)
		roleUrl.PUT("/:id", roleApi.UpdateRole)
		roleUrl.DELETE("/:id", roleApi.DelRole)
		roleUrl.POST("", roleApi.CreateRole)

	}
}
