package service

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		// 添加全部
		AllowHeaders: []string{
			"*",
		},
		ExposeHeaders:    []string{"Content-Length", "Content-Range"},
		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	})
}
