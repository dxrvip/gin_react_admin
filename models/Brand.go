package models

// Brand 品牌模型
type Brand struct {
	BaseModel
	Name        string `gorm:"type:varchar(255);not null;unique" json:"name"`
	Logo        string `gorm:"type:varchar(255)" json:"logo"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}
