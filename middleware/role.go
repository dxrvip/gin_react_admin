package middleware

import (
	"github.com/gin-gonic/gin"
)

type Role struct {
	Name    string
	Actions map[string]bool // 存储函数名及其权限
}

// 定义角色和其可访问的 API 方法
var roles = []Role{
	{
		Name: "admin",
		Actions: map[string]bool{
			"GetUser":    true,
			"UpdateUser": true,
			"DeleteUser": true,
		},
	},
	{
		Name: "user",
		Actions: map[string]bool{
			"GetUser":    true,
			"UpdateUser": false,
			"DeleteUser": false,
		},
	},
}

// 权限检查中间件
func RoleMiddleware(requiredAction string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 假设用户角色存储在请求的上下文中
		// userRole := c.MustGet("userRole").(string)
		// userRole := "userRole"

		// if !authorize(userRole, requiredAction) {
		// 	c.JSON(http.StatusNotFound, gin.H{"error": "权限不足"})
		// 	c.Abort()
		// 	return
		// }

		c.Next()
	}
}

// 检查用户角色的权限
func authorize(role string, action string) bool {
	for _, r := range roles {
		if r.Name == role {
			return r.Actions[action]
		}
	}
	return false
}
