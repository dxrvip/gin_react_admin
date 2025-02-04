package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AppendError(existErr, newErr error) error {
	if existErr == nil {
		return newErr
	}

	return fmt.Errorf("%v, %w", existErr, newErr)
}

func GetUserId(c *gin.Context) (uint, error) {
	userId, exists := c.Get("id")
	if !exists {
		return 0, fmt.Errorf("获取用户ID失败")
	}
	id, exists := userId.(int)
	if !exists {
		return 0, fmt.Errorf("用户ID类型错误")
	}
	return uint(id), nil
}
