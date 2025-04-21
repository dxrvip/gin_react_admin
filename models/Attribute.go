package models

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type AttributeType string

const (
	TypeString AttributeType = "string"
	TypeNumber AttributeType = "number"
	TypeEnum   AttributeType = "enum"
)

// 分类-属性中间表
type CategoryAttribute struct {
	ProductCategoryID uint `gorm:"primaryKey"`
	AttributeID       uint `gorm:"primaryKey"`
}

type Options struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// 重命名类型避免歧义
type OptionList []Options

func (l OptionList) Value() (driver.Value, error) {
	return json.Marshal(l)
}
func (l *OptionList) Scan(value interface{}) error {
	// 实现数据库驱动接口，用于从数据库读取JSON数据
	// 1. 类型检查：确认数据库返回的是字节切片格式
	bytes, ok := value.([]byte)
	if !ok {
		return nil // 非字节类型直接返回（适用于NULL值情况）
	}

	// 2. 将JSON字节数据反序列化为OptionList对象
	// 注意：这里使用指针接收器，确保能修改原始对象
	return json.Unmarshal(bytes, l)
}

// 属性模型
type Attribute struct {
	BaseModel
	Name       string        `gorm:"type:varchar(100)" json:"name"`
	Type       AttributeType `gorm:"type:varchar(20);not null" json:"type"`
	IsRequired bool          `json:"isRequired"`
	// 公共字段
	DefaultValue string `gorm:"type:varchar(255)" json:"defaultValue,omitempty"`
	// 数字类型专用
	MinValue *float64 `gorm:"type:decimal(10,2);default(0)" json:"min,omitempty"`
	MaxValue *float64 `gorm:"type:decimal(10,2);default(10000)" json:"max,omitempty"`
	// 枚举类型专用

	Options      OptionList         `gorm:"type:json" json:"-"`
	OptionsArray []string           `gorm:"-" json:"options"`
	Categories   []*ProductCategory `gorm:"many2many:category_attributes;joinForeignKey:attribute_id"`
	// Categories   []*ProductCategory `gorm:"many2many:category_attributes;joinForeignKey:attribute_id;joinReferencesKey:product_category_id"`

	CategoryIDs []uint `gorm:"-" json:"categoryIds"`

	// 属性单位
	Unit string `gorm:"type:varchar(20);default:''" json:"unit"` // 单位
}

func (a *Attribute) BeforeSave(tx *gorm.DB) error {
	// 将 OptionsArray 转换为 OptionList
	if a.OptionsArray != nil {
		var options OptionList
		for _, v := range a.OptionsArray {
			options = append(options, Options{Value: v, Label: v})
		}
		a.Options = options
	}
	return nil
}
