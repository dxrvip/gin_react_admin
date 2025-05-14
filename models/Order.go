package models

import (
	"time"
)

// 订单主表
type Order struct {
	BaseModel
	OrderNo       string      `gorm:"type:varchar(50);not null;unique" json:"orderNo"` // 订单编号
	UserID        uint        `gorm:"not null" json:"userId"`                          // 用户ID
	Address       string      `gorm:"type:varchar(255);not null" json:"address"`       // 收货地址ID
	TotalAmount   float64     `gorm:"type:decimal(10,2);not null" json:"totalAmount"`  // 订单总金额
	CostPrice     float64     `gorm:"type:decimal(10,2);not null" json:"costPrice"`    // 成本价格
	Status        string      `gorm:"type:varchar(20);not null" json:"status"`         // 订单状态：pending/paid/shipping/completed/cancelled
	Note          string      `gorm:"type:varchar(500)" json:"note"`                   // 订单备注
	OrderItems    []OrderItem `gorm:"foreignKey:OrderID" json:"orderItems"`            // 订单项
	PaymentTime   *time.Time  `json:"paymentTime,omitempty"`                           // 支付时间
	ShippingTime  *time.Time  `json:"shippingTime,omitempty"`                          // 发货时间
	CompletedTime *time.Time  `json:"completedTime,omitempty"`                         // 完成时间
	Weight        uint        `gorm:"not nul" json:"weight"`                           // 重量
}

/*
	    { id: 'pending', name: '待付款' },
		{ id: 'paid', name: '已付款' },
		{ id: 'shipping', name: '已发货' },
		{ id: 'completed', name: '已完成' },
		{ id: 'cancelled', name: '已取消' },
*/
// func (o *Order) AfterFind(tx *gorm.DB) (err error) {
// 	switch o.Status {
// 	case "pending":
// 		o.Status = "待支付"
// 	case "paid":
// 		o.Status = "已付款"
// 	case "shipping":
// 		o.Status = "已发货"
// 	case "completed":
// 		o.Status = "已完成"
// 	case "cancelled":
// 		o.Status = "已取消"
// 	default:
// 		return
// 	}
// 	return
// }

/*
订单表设计：
	1. 订单id，通过雪花算法不重复
	4. 下单商品，一个订单可下多个商品
	5. 下单用户，用户id
	6. 收货地址，地址id
	3. 订单总价，通过订单项的商品价格和数量计算得出
	2. 订单状态，待支付，待发货，待收货，已完成，已取消
	7. 订单留言，下单备注
*/

// 订单项
type OrderItem struct {
	BaseModel
	OrderID   uint          `gorm:"not null" json:"orderId"`                                                                          // 订单ID
	ProductID uint          `gorm:"not null" json:"productId" binding:"required"`                                                     // 商品ID
	Quantity  uint          `gorm:"not null" json:"quantity" binding:"required"`                                                      // 购买数量
	Price     float64       `gorm:"type:decimal(10,2);not null" json:"price" binding:"required"`                                      // 商品单价
	Product   SecondHandSku `gorm:"foreignKey:ID;references:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product"` // 关联商品信息
}
