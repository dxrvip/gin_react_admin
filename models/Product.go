package models

// 商品分类（支持多级）
type ProductCategory struct {
	BaseModel
	Name     string     `gorm:"type:varchar(100);not null" json:"name"`
	ParentID *uint      `gorm:"index" json:"parent_id"`
	Children []Category `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}

// 属性键表（定义属性名称，如 "颜色"、"尺寸"）
type AttributeKey struct {
	BaseModel
	Name       string           `gorm:"type:varchar(100);not null;unique" json:"name" validate:"required"` // 属性名称
	CategoryID uint             `gorm:"not null" json:"category_id"`                                       // 所属分类（可选）
	Values     []AttributeValue `gorm:"foreignKey:KeyID" json:"values"`                                    // 可选预定义值（如颜色可选值）
}

// 属性值表（存储具体值，如 "红色"、"XL"）
type AttributeValue struct {
	BaseModel
	KeyID uint   `gorm:"not null" json:"key_id"`                  // 关联属性键
	Value string `gorm:"type:varchar(100);not null" json:"value"` // 属性值
}

// 商品属性关联表（动态属性）
type ProductAttribute struct {
	ProductID uint           `gorm:"primaryKey" json:"product_id"`   // 商品ID
	KeyID     uint           `gorm:"primaryKey" json:"key_id"`       // 属性键ID
	ValueID   *uint          `gorm:"index" json:"value_id"`          // 预定义值ID（可选）
	Value     string         `gorm:"type:varchar(255)" json:"value"` // 自定义值（当ValueID为空时使用）
	Key       AttributeKey   `gorm:"foreignKey:KeyID" json:"key"`
	Val       AttributeValue `gorm:"foreignKey:ValueID" json:"val,omitempty"`
}

// 商品信息
type Product struct {
	BaseModel
	Name              string   `gorm:"type:varchar(200);not null" json:"name" validate:"required"`
	Price             float64  `gorm:"type:decimal(10,2);not null" json:"price" validate:"gt=0"`
	Stock             uint     `gorm:"not null" json:"stock" validate:"gte=0"`
	Description       string   `gorm:"type:text" json:"description"`
	Images            []string `gorm:"type:json" json:"images"` // 图片URL列表
	ProductCategoryID uint     `gorm:"not null" json:"category_id"`
	ProductCategory   Category `gorm:"foreignKey:CategoryID" json:"category"`
	Status            string   `gorm:"type:varchar(20);default:'active'" json:"status"` // active/disabled

	BrandID    uint               `gorm:"not null" json:"brand_id"` // 关联品牌
	Brand      Brand              `gorm:"foreignKey:BrandID" json:"brand,omitempty"`
	Attributes []ProductAttribute `gorm:"foreignKey:ProductID" json:"attributes"` // 动态属性
}
