// validators/attribute.go
package validators

import (
	"goVueBlog/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterAttributeValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册枚举选项验证
		v.RegisterValidation("validoptions", func(fl validator.FieldLevel) bool {
			if attrType := fl.Parent().FieldByName("Type"); attrType.IsValid() {
				if attrType.String() == "enum" {
					return fl.Field().Len() > 0
				}
			}
			return true
		})
	}
}

func ContainsDuplicates(arr models.OptionList) bool {
	// 创建map用于记录已出现的值
	seen := make(map[string]bool)

	for _, item := range arr {
		// 直接访问Options结构体的Value字段
		if seen[item.Value] {
			return true
		}
		seen[item.Value] = true
	}
	return false
}

func ValidateNumberRange(c *gin.Context) {
	var input struct {
		Min *float64 `json:"min"`
		Max *float64 `json:"max"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "无效输入"})
		return
	}

	if input.Min != nil && input.Max != nil && *input.Min > *input.Max {
		c.AbortWithStatusJSON(400, gin.H{"error": "最小值不能大于最大值"})
		return
	}

	c.Next()
}
