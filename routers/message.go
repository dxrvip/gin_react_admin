package routers

import (
	"goVueBlog/api"

	"github.com/gin-gonic/gin"
)

func RegisterMessageUrls(r *gin.RouterGroup) {
	api := api.NewMessageApi()
	message := r.Group("/message")
	{
		message.GET("", api.ListMessage)
		message.GET("/:id", api.GetMessageById)
		message.PUT("/:id", api.UpdateMessage)
		message.DELETE("/:id", api.DeleteMessage)
		message.POST("", api.CreateMessage)

	}
}
