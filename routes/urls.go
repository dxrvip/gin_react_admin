package routes

import (
	"goVueBlog/api"
	"goVueBlog/docs"
	"goVueBlog/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitUrlsRouter() *gin.Engine {

	// 设置模式
	gin.SetMode(viper.GetString("servers.debug"))
	r := gin.Default()
	// 解决跨域问题
	r.Use(middleware.Cors())
	//添加根路由
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	// 用户登陆，注册路由
	apiUser := api.NewUserApi()
	v1.POST("/user/login", apiUser.Login)
	v1.POST("/user/register", apiUser.Register)

	// 注册中间件r.Use(middleware.AuthMiddleware)
	v1.Use(middleware.JwtToken())
	// swag
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//=================================================
	//=> 用户路由
	RegisterUserUrl(v1, apiUser)
	//=================================================
	//=> 文章路由
	RegisterBlogUrls(v1)
	//=================================================
	//=> 分类路由
	RegisterCategoryUrls(v1)

	//=> 菜单路由
	RegisterSystemUrls(v1)
	//=> 角色路由
	RegisterRoleUrls(v1)
	//=> 部门路由
	RegisterDepartmentUrls(v1)
	//=> 消息路由
	RegisterMessageUrls(v1)
	//=> 商城路由
	RegisterRouters(v1)
	//=> 文件上传
	v1.GET("/auto/upload", api.UploadImage)
	return r
}
