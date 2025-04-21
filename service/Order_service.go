package service

import (
	"encoding/json"
	"fmt"
	"goVueBlog/models"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var orderService *OrderService

type CustomDateTime time.Time

// UnmarshalJSON 反序列化
func (ct *CustomDateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(time.RFC3339Nano, s) // 支持带毫秒的 ISO 8601
	if err != nil {
		return fmt.Errorf("invalid ISO 8601 time format: %v", err)
	}
	*ct = CustomDateTime(t.UTC()) // 强制转为 UTC 时间
	return nil
}

// MarshalJSON 序列化（可选）
func (ct CustomDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(ct).Format(time.RFC3339Nano))

}

type OrderService struct {
	BaseService
}
type OrderQueryBase struct {
	Address   string         `json:"address" binding:"required"`
	Note      string         `json:"note"`
	ClienTime CustomDateTime `json:"createdAt" binding:"required"`
	CostPrice float64        `json:"costPrice" binding:"required,min=0.01"`
	Weight    uint           `json:"weight" binding:"required,min=0"`
	UserId    uint           `json:"userId" binding:"required,min=1"`
	Status    string         `json:"status" binding:"required,oneof=pending paid shipping completed cancelled"`
}

// CreateOrderRequest 创建订单请求
type CreateOrderRequest struct {
	OrderQueryBase
	Items []CreateOrderItemRequest `json:"items" binding:"required"`
}

// UpdateOrderRequest 更新订单请求
type UpdateOrderRequest struct {
	ID uint `json:"id" binding:"required,min=1"`
	OrderQueryBase
	Items []models.OrderItem `json:"items" binding:"required"` // 更新订单时必须提供订单项
}

// CreateOrderItemRequest 创建订单项请求
type CreateOrderItemRequest struct {
	ProductID uint    `json:"productId" binding:"required"`
	Quantity  uint    `json:"quantity" binding:"required,gt=0"`
	Price     float64 `json:"price" binding:"required,gt=0"`
}

// OrderResponse 订单响应
type OrderResponse struct {
	ID            uint                `json:"id"`
	OrderNo       string              `json:"orderNo"`
	UserID        uint                `json:"userId"`
	Address       string              `json:"address"`
	TotalAmount   float64             `json:"totalAmount"`
	Status        string              `json:"status"`
	Note          string              `json:"note"`
	OrderItems    []OrderItemResponse `json:"items"`
	PaymentTime   *time.Time          `json:"paymentTime,omitempty"`
	ShippingTime  *time.Time          `json:"shippingTime,omitempty"`
	CompletedTime *time.Time          `json:"completedTime,omitempty"`
	CreatedAt     time.Time           `json:"createdAt"`
	CostPrice     float64             `json:"costPrice"`
	Weight        uint                `json:"weight"`
	Description   string              `json:"description"`
}

// OrderItemResponse 订单项响应
type OrderItemResponse struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"productId"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
}

func NewOrderService() *OrderService {
	if orderService == nil {
		return &OrderService{
			BaseService: NewBaseApi(&models.Order{}),
		}
	}
	return orderService
}

// 判断用户是否存在
func (s *OrderService) isUserId(userId uint) error {
	var user = models.User{}
	user.ID = userId
	return s.DB.First(&user).Error
}

// CreateOrder 创建订单
func (s *OrderService) CreateOrder(req *CreateOrderRequest) (*OrderResponse, error) {
	// 开启事务
	tx := s.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 判断用户是否存在
	if err := s.isUserId(req.UserId); err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "用户不存在")
	}
	// 生成订单号
	orderNo := fmt.Sprintf("%d%d", time.Now().UnixNano(), req.UserId)

	// 创建订单项并计算总金额
	var totalAmount float64
	var orderItems []models.OrderItem

	for _, item := range req.Items {
		// 查询商品信息，是查二手商品了
		var secondHandSku models.SecondHandSku
		if err := tx.First(&secondHandSku, item.ProductID).Error; err != nil {
			tx.Rollback()
			return nil, errors.Wrap(err, "商品不存在")
		}

		// 检查库存
		if secondHandSku.Stock < item.Quantity {
			tx.Rollback()
			return nil, fmt.Errorf("货号 %s 库存不足", secondHandSku.ProductsType)
		}

		// 扣减库存， 二手商品默认扣减1
		if err := tx.Model(&secondHandSku).Update("stock", gorm.Expr("stock - ?", item.Quantity)).Error; err != nil {
			tx.Rollback()
			return nil, errors.Wrap(err, "更新库存失败")
		}

		// 创建订单项
		orderItem := models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
		orderItems = append(orderItems, orderItem)
		totalAmount += item.Price * float64(item.Quantity)
	}

	// 创建订单
	order := &models.Order{
		OrderNo:     orderNo,
		UserID:      req.UserId,
		Address:     req.Address,
		TotalAmount: totalAmount,
		Status:      req.Status,
		Note:        req.Note,
		OrderItems:  orderItems,
		CostPrice:   req.CostPrice,
		Weight:      req.Weight,
	}
	// 对时间进行转换
	clientTime := time.Time(req.ClienTime)

	order.BaseModel.CreatedAt = clientTime
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "创建订单失败")
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, errors.Wrap(err, "提交事务失败")
	}

	// 构建响应
	response := &OrderResponse{
		ID:          order.ID,
		OrderNo:     order.OrderNo,
		UserID:      order.UserID,
		Address:     order.Address,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
		Note:        order.Note,
		CreatedAt:   order.CreatedAt,
	}

	for _, item := range order.OrderItems {
		response.OrderItems = append(response.OrderItems, OrderItemResponse{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}

	return response, nil
}

// UpdateOrderStatus 更新订单状态
func (s *OrderService) UpdateOrderStatus(orderID uint, status string) error {
	updates := map[string]interface{}{
		"status": status,
	}

	// 根据状态更新相应的时间字段
	switch status {
	case "paid":
		updates["payment_time"] = time.Now()
	case "shipping":
		updates["shipping_time"] = time.Now()
	case "completed":
		updates["completed_time"] = time.Now()
	}

	if err := s.DB.Model(&models.Order{}).Where("id = ?", orderID).Updates(updates).Error; err != nil {
		return errors.Wrap(err, "更新订单状态失败")
	}

	return nil
}
func (s *OrderService) Updates(id uint, updateReq interface{}) (*OrderResponse, error) {
	// 开启事务
	// 类型断言
	req, ok := updateReq.(*models.Order)
	if !ok {
		return nil, fmt.Errorf("无效的请求类型")
	}
	tx := s.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 判断用户是否存在
	if err := s.isUserId(req.UserID); err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "用户不存在")
	}
	// 更新订单
	/*
		1.获取订单商品信息与更新商品信息进行对比是否有商品信息发生变化
		2.如果有变化，更新库存
		3.如果没有变化，不更新库存
	*/
	var order models.Order
	if err := tx.Preload("OrderItems").First(&order, req.ID).Error; err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "订单不存在")
	}
	// 先删除订单商品信息
	if err := tx.Where("order_id =?", req.ID).Delete(&models.OrderItem{}).Error; err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "删除订单商品信息失败")
	}
	// 在查询二手商品是否存在，如果存在就进行库存增加】
	for _, item := range order.OrderItems {
		// 查询二手商品信息
		var secondHandSku models.SecondHandSku
		if err := tx.First(&secondHandSku, item.ProductID).Error; err != nil {
			continue // 如果商品不存在，跳过
		} else {
			// 否则加库存
			if err := tx.Model(&secondHandSku).Update("stock", gorm.Expr("stock + ?", item.Quantity)).Error; err != nil {
				tx.Rollback()
				return nil, errors.Wrap(err, "更新库存失败")
			}
		}
	}
	// 创建新的订单项并计算总金额
	var totalAmount float64

	var orderItems []models.OrderItem
	for _, item := range req.OrderItems {
		// 查询商品信息，是查二手商品了
		var secondHandSku models.SecondHandSku
		if err := tx.First(&secondHandSku, item.ProductID).Error; err != nil {
			tx.Rollback()
			return nil, errors.Wrap(err, "商品不存在")
		}
		// 检查库存
		if secondHandSku.Stock < item.Quantity {
			tx.Rollback()
			return nil, fmt.Errorf("货号 %s 库存不足", secondHandSku.ProductsType)
		}
		// 扣减库存， 二手商品默认扣减1
		if err := tx.Model(&secondHandSku).Update("stock", gorm.Expr("stock -?", item.Quantity)).Error; err != nil {
			tx.Rollback()
			return nil, errors.Wrap(err, "更新库存失败")
		}
		// 创建订单项
		orderItem := models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
		orderItems = append(orderItems, orderItem)
		// 计算总金额
		totalAmount = totalAmount + item.Price*float64(item.Quantity)
	}
	// 更新订单信息
	order.Address = req.Address
	order.TotalAmount = totalAmount
	order.Status = req.Status
	order.Note = req.Note
	order.OrderItems = orderItems
	order.CostPrice = req.CostPrice
	order.Weight = req.Weight
	// 对时间进行转换
	clientTime := time.Time(req.CreatedAt)
	order.BaseModel.CreatedAt = clientTime
	// 更新订单
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "更新订单失败")

	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, errors.Wrap(err, "提交事务失败")
	}
	// 构建响应
	response := &OrderResponse{
		ID:          order.ID,
		OrderNo:     order.OrderNo,
		UserID:      order.UserID,
		Address:     order.Address,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
		Note:        order.Note,
		CreatedAt:   order.CreatedAt,
		CostPrice:   order.CostPrice,
		Weight:      order.Weight,
	}
	return response, nil
}

// GetOrderByID 获取订单详情
func (s *OrderService) GetOrderByID(orderID uint) (*OrderResponse, error) {
	var order models.Order
	if err := s.DB.Preload("OrderItems.Product").First(&order, orderID).Error; err != nil {
		return nil, errors.Wrap(err, "获取订单失败")
	}

	response := &OrderResponse{
		ID:            order.ID,
		OrderNo:       order.OrderNo,
		UserID:        order.UserID,
		Address:       order.Address,
		TotalAmount:   order.TotalAmount,
		Status:        order.Status,
		Note:          order.Note,
		PaymentTime:   order.PaymentTime,
		ShippingTime:  order.ShippingTime,
		CompletedTime: order.CompletedTime,
		CreatedAt:     order.CreatedAt,
		CostPrice:     order.CostPrice, // 成本价格
		Weight:        order.Weight,    // 重量
	}

	for _, item := range order.OrderItems {
		response.OrderItems = append(response.OrderItems, OrderItemResponse{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}

	return response, nil
}

// GetUserOrders 获取用户订单列表
func (s *OrderService) GetUserOrders(userID uint, page, pageSize int) ([]OrderResponse, int64, error) {
	var orders []models.Order
	var total int64

	query := s.DB.Model(&models.Order{}).Where("user_id = ?", userID)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, errors.Wrap(err, "获取订单总数失败")
	}

	// 获取分页数据
	if err := query.Preload("OrderItems.Product").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&orders).Error; err != nil {
		return nil, 0, errors.Wrap(err, "获取订单列表失败")
	}

	var responses []OrderResponse
	for _, order := range orders {
		response := OrderResponse{
			ID:            order.ID,
			OrderNo:       order.OrderNo,
			UserID:        order.UserID,
			Address:       order.Address,
			TotalAmount:   order.TotalAmount,
			Status:        order.Status,
			Note:          order.Note,
			PaymentTime:   order.PaymentTime,
			ShippingTime:  order.ShippingTime,
			CompletedTime: order.CompletedTime,
			CreatedAt:     order.CreatedAt,
		}

		for _, item := range order.OrderItems {
			response.OrderItems = append(response.OrderItems, OrderItemResponse{
				ID:        item.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Price,
			})
		}

		responses = append(responses, response)
	}

	return responses, total, nil
}
