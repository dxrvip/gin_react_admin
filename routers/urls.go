package routers

import (
	"goVueBlog/api"
	"goVueBlog/docs"
	"goVueBlog/middleware"
	"goVueBlog/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitUrlsRouter() *gin.Engine {
	gin.SetMode(viper.GetString("servers.debug"))
	r := gin.Default()
	// 解决跨域问题
	r.Use(service.Cors())
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	// 用户登陆，注册路由
	v1.POST("/user/login", api.Login)
	v1.POST("/user/register", api.Register)

	// 注册中间件r.Use(middleware.AuthMiddleware)
	v1.Use(middleware.AuthMiddleware)
	// swag
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 图片上传路由

	UserUrls(v1)     // 注册用户路由
	InitRouter(v1)   // 注册文章路由
	InitCategory(v1) // 注册分类路由
	v1.GET("/auto/upload", api.UploadImage)
	return r
}
