package models

// 订单主表
type Order struct {
	BaseModel
	UserID     uint        `gorm:"not null" json:"user_id"`
	User       User        `gorm:"foreignKey:UserID" json:"user"`
	TotalPrice float64     `gorm:"type:decimal(10,2);not null" json:"total_price"`
	Status     string      `gorm:"type:varchar(20);default:'pending'" json:"status"` // pending/paid/shipped/completed
	Items      []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

// 订单项
type OrderItem struct {
	BaseModel
	OrderID   uint    `gorm:"not null" json:"order_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  uint    `gorm:"not null" json:"quantity" validate:"gt=0"`
	Price     float64 `gorm:"type:decimal(10,2);not null" json:"price"`
}
