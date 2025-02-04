package models

// Brand 品牌模型
type Brand struct {
	BaseModel
	Name        string    `gorm:"type:varchar(100);not null;unique" json:"name" validate:"required"` // 品牌名称
	Logo        string    `gorm:"type:varchar(255)" json:"logo"`                                     // Logo URL
	Description string    `gorm:"type:text" json:"description"`                                      // 品牌描述
	Products    []Product `gorm:"foreignKey:BrandID" json:"products,omitempty"`                      // 关联商品
}
