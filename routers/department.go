package routers

import (
	"goVueBlog/api"

	"github.com/gin-gonic/gin"
)

func RegisterDepartmentUrls(r *gin.RouterGroup) {
	api := api.NewDepartmentApi()
	departmentUrl := r.Group("/department")
	{
		departmentUrl.GET("", api.ListDepartment)
		departmentUrl.GET("/:id", api.GetDepartmentById)
		departmentUrl.PUT("/:id", api.UpdateDepartment)
		departmentUrl.DELETE("/:id", api.DeleteDepartment)
		departmentUrl.POST("", api.CreateDepartment)

	}
}
