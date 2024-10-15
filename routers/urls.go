package routers

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
	v1.Use(middleware.AuthMiddleware)
	// swag
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 图片上传路由

	//=================================================
	//=> 用户路由
	user := v1.Group("/user")
	user.GET("/info", apiUser.Info)
	user.DELETE("/:id", apiUser.Delete)
	user.PUT("/:id", apiUser.Update)
	//=================================================
	//=> 文章路由
	articleApi := api.NewArticleApi()
	article := v1.Group("/article")
	{
		// 添加路由
		article.GET("", articleApi.ArticleList)
		article.PUT("/:id", articleApi.ArticleUpdate)
		article.GET("/:id", articleApi.ArticleInfo)

		article.DELETE("/:id", articleApi.ArticleDelete)
		article.POST("", articleApi.CreateArticle)
	}
	//=================================================
	//=> 分类路由
	categoryApi := api.NewCategoryApi()
	category := v1.Group("/category")
	{
		category.GET("", categoryApi.GetCategoryList)

		category.POST("", categoryApi.AddCategory)
		category.GET("/:id", categoryApi.GetCategoryById)
		category.PUT("/:id", categoryApi.UpdateCategory)

		category.DELETE("/:id", categoryApi.DeleteCategory)
	}
	//
	v1.GET("/auto/upload", api.UploadImage)
	return r
}
