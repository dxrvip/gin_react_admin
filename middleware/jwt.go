package middleware

import (
	"fmt"
	"goVueBlog/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// AuthMiddleware是我们用来检查令牌是否有效的中间件。如果返回401状态无效，则返回给客户。
func AuthMiddleware(c *gin.Context) {
	// 检查请求头中是否包含token
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的token"})
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	// 验证token
	appKey := viper.GetString("app.Key")
	claims, err := utils.VerifyJWT(tokenString, []byte(appKey))
	if err != nil {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.Set("username", claims.Foo)
	c.Set("id", claims.Id)

	fmt.Printf("claims: %v\n", claims)
	c.Next()
}
