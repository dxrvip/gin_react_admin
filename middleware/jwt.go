package middleware

import (
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// JWT中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			utils.ResponseAuthError(c, errmsg.ERROR_TOKEN_NOT_EXIST, nil)
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			utils.ResponseAuthError(c, errmsg.ERROR_TOKEN_TYPE_WRONG, nil)
			c.Abort()
			return
		}

		signKey := []byte(viper.GetString("app.Key"))
		claims, err := utils.VerifyJWT(checkToken[1], signKey)
		if err != nil {
			utils.ResponseAuthError(c, http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		c.Set("id", claims.Id)
		c.Next()
	}
}
