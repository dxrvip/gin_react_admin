package service

import (
	"goVueBlog/models"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var secondHandSkuService *SecondHandSkuService

type SecondHandSkuService struct {
	BaseService
}

func NewSecondHandSkuService() *SecondHandSkuService {
	if secondHandSkuService == nil {
		return &SecondHandSkuService{
			BaseService: NewBaseApi(&models.SecondHandSku{}),
		}
	}
	return secondHandSkuService
}

// CreateSecondHandSkuRequest 创建二手商品SKU请求
type CreateSecondHandSkuRequest struct {
	// 商品标题默认空

	Title           string               `json:"title" binding:"max=60"` //商品标题不强制填写，默认取商品空
	ProductID       uint                 `json:"productId" binding:"required"`
	Price           float64              `json:"price" binding:"required,gt=0"`
	Stock           uint                 `json:"stock" binding:"required,gte=0"`
	Condition       string               `json:"condition" binding:"required,oneof=new like_new light obvious serious damaged"`
	Function        string               `json:"function" binding:"required,oneof=perfect repaired usable unusable"`
	UsageDuration   string               `json:"usageDuration" binding:"required,oneof=unused half_year one_year three_years"`
	Accessories     bool                 `json:"accessories"`
	AccessoriesList string               `json:"accessoriesList"`
	FreeShipping    bool                 `json:"freeShipping"`
	ShippingFee     float64              `json:"shippingFee" binding:"gte=0"`
	Description     string               `json:"description"`
	Images          models.PictureList   `json:"picture" binding:"required"`                   // 商品图片
	Cost            models.Float64String `json:"cost"`                                         // 成本价格
	ProductsType    string               `json:"productsType" binding:"required,max=30,min=5"` //货号
	IsRepair        bool                 `json:"isRepair"`                                     //是否维修过
	RepairEndDate   CustomDate           `json:"repairEndDate"`                                //维修结束时间
	Status          models.ProductStatus `json:"status" binding:"required"`                    // 商品状态
	BatteryLife     uint                 `json:"batteryLife"`                                  //电池工作时间（单位： 分钟）
	WorkingTime     uint                 `json:"workingTime"`                                  // 维修后电池工作时间（单位：分钟）
}

// CreateSecondHandSku 创建二手商品SKU
func (s *SecondHandSkuService) CreateSecondHandSku(req *CreateSecondHandSkuRequest) (*models.SecondHandSku, error) {
	// 检查商品是否存在
	var product models.Product
	if err := s.DB.First(&product, req.ProductID).Error; err != nil {
		return nil, errors.Wrap(err, "商品不存在")
	}
	// 生成一个商品id
	productSkuID := uuid.New().String()

	// 创建SKU
	sku := &models.SecondHandSku{
		ProductSkuID:    productSkuID,
		ProductID:       req.ProductID,
		Title:           req.Title,
		Cost:            req.Cost,
		ProductsType:    req.ProductsType,
		AccessoriesList: req.AccessoriesList,
		Price:           req.Price,
		Stock:           req.Stock,
		Condition:       req.Condition,
		Function:        req.Function,
		UsageDuration:   req.UsageDuration,
		Accessories:     req.Accessories,
		FreeShipping:    req.FreeShipping,
		ShippingFee:     req.ShippingFee,
		Description:     req.Description,
		Images:          req.Images,
		IsRepair:        req.IsRepair,
		RepairEndDate:   time.Time(req.RepairEndDate), // 默认值为零值
		Status:          req.Status,
		BatteryLife:     req.BatteryLife,
		WorkingTime:     req.WorkingTime,
	}

	if err := s.DB.Create(sku).Error; err != nil {
		return nil, errors.Wrap(err, "创建二手商品SKU失败")
	}

	return sku, nil
}

// UpdateSecondHandSkuRequest 更新二手商品SKU请求
type UpdateSecondHandSkuRequest struct {
	Price         *float64              `json:"price" validate:"omitempty,gt=0"`
	Stock         *uint                 `json:"stock" validate:"omitempty,gte=0"`
	Condition     *string               `json:"condition" validate:"omitempty,oneof=new like_new light obvious serious damaged"`
	Function      *string               `json:"function" validate:"omitempty,oneof=perfect repaired usable unusable"`
	UsageDuration *string               `json:"usageDuration" validate:"omitempty,oneof=unused half_year one_year three_years"`
	Accessories   *bool                 `json:"accessories"`
	FreeShipping  *bool                 `json:"freeShipping"`
	ShippingFee   *float64              `json:"shippingFee" validate:"omitempty,gte=0"`
	Description   *string               `json:"description"`
	Images        *models.PictureList   `json:"images"`
	Status        *models.ProductStatus `json:"status" validate:"omitempty"`
	// 工作时间记录
	WorkingTime *uint `json:"workingTime" validate:"omitempty,gte=0"` // 维修后电池工作时间（单位：分钟）

}

// UpdateSecondHandSku 更新二手商品SKU
func (s *SecondHandSkuService) UpdateSecondHandSku(id uint, req *UpdateSecondHandSkuRequest) error {
	updates := make(map[string]interface{})

	if req.Price != nil {
		updates["price"] = *req.Price
	}
	if req.Stock != nil {
		updates["stock"] = *req.Stock
	}
	if req.Condition != nil {
		updates["condition"] = *req.Condition
	}
	if req.Function != nil {
		updates["function"] = *req.Function
	}
	if req.UsageDuration != nil {
		updates["usage_duration"] = *req.UsageDuration
	}
	if req.Accessories != nil {
		updates["accessories"] = *req.Accessories
	}
	if req.FreeShipping != nil {
		updates["free_shipping"] = *req.FreeShipping
	}
	if req.ShippingFee != nil {
		updates["shipping_fee"] = *req.ShippingFee
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.Images != nil {
		updates["images"] = *req.Images
	}

	if err := s.DB.Model(&models.SecondHandSku{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return errors.Wrap(err, "更新二手商品SKU失败")
	}

	return nil
}

// GetSecondHandSkuByID 获取二手商品SKU详情
func (s *SecondHandSkuService) GetSecondHandSkuByID(id uint) (*models.SecondHandSku, error) {
	var sku models.SecondHandSku
	if err := s.DB.Preload("Product").First(&sku, id).Error; err != nil {
		return nil, errors.Wrap(err, "获取二手商品SKU失败")
	}
	return &sku, nil
}

// GetProductSecondHandSkus 获取商品的所有二手SKU
func (s *SecondHandSkuService) GetProductSecondHandSkus(productID uint) ([]models.SecondHandSku, error) {
	var skus []models.SecondHandSku
	if err := s.DB.Where("product_id = ?", productID).Find(&skus).Error; err != nil {
		return nil, errors.Wrap(err, "获取商品二手SKU列表失败")
	}
	return skus, nil
}

// DeleteSecondHandSku 删除二手商品SKU
func (s *SecondHandSkuService) DeleteSecondHandSku(id uint) error {
	if err := s.DB.Delete(&models.SecondHandSku{}, id).Error; err != nil {
		return errors.Wrap(err, "删除二手商品SKU失败")
	}
	return nil
}
