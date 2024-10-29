package routers

import "github.com/gin-gonic/gin"

func RegisterNewsUrls(r *gin.RouterGroup) {
	url := r.Group("news")
	{
		url.GET("")
	}
}
