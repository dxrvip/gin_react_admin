package service

import (
	"fmt"
	"goVueBlog/models"
	"goVueBlog/service/serializer"
	"reflect"
	"time"

	"gorm.io/gorm"
)

var productService *ProductService

type ProductService struct {
	BaseService
}

type ProductRequry struct {
	Title             string               `json:"title" binding:"required"`
	Price             models.Float64String `json:"price" binding:"gt=0"`
	Stock             uint                 `json:"stock" binding:"gte=0"`
	Description       string               `json:"description"`
	Images            models.PictureList   `json:"picture"` // 图片URL列表
	ProductCategoryID uint                 `json:"productCategoryID"`
	Status            string               `json:"status"` // active/disabled

	BrandID    uint                `json:"brandID" binding:"numeric"` // 关联品牌
	Attributes models.AttributeMap `json:"attributes"`
}

type ProductResponse struct {
	ID                uint                   `json:"id"`
	Title             string                 `json:"title"`
	Price             models.Float64String   `json:"price"`
	Stock             uint                   `json:"stock"`
	Description       string                 `json:"description"`
	Images            models.PictureList     `json:"picture"`
	ProductCategoryID uint                   `json:"productCategoryID"`
	Status            string                 `json:"status"`
	BrandID           uint                   `json:"brandID"`
	Attributes        models.AttributeMap    `json:"attributes"`
	CreatedAt         time.Time              `json:"createdAt"`
	SecondHandSku     []models.SecondHandSku `json:"secondHandSku,omitempty"`
}

func NewProductService() *ProductService {
	if productService == nil {
		return &ProductService{
			BaseService: NewBaseApi(&models.Product{}),
		}
	}
	return productService
}

func (m *ProductService) List(datas interface{}, params *serializer.CommonQueryOtpones) (string, error) {
	// 添加查询条件
	query := m.DB.Model(m.Model)
	// 当查询条件中带二手商品时候加载二手商品，并且过滤掉库存0，状态不是active的商品

	// Preload("SecondHandSku") // 预加载关联分类
	// Preload("Values", func(db *gorm.DB) *gorm.DB {
	// 	return db.Order("id DESC")
	// })
	// 构建查询条件
	for key, value := range params.Filter {
		if key == "status" && value == "all" {
			// 如果状态是 "all"，则不添加任何过滤条件
			continue
		}
		// 当查询条件中带二手商品时候加载二手商品，并且过滤掉库存0，状态不是active的商品
		if key == "is_second_hand_sku" {
			query = query.Preload("SecondHandSku", func(db *gorm.DB) *gorm.DB {
				return db.Where("stock >?", 0).Where("status =?", "active")
			})
			continue
		}
		// 对q字段进行对title模糊查询
		if key == "q" && value != "" {
			query = query.Where("title LIKE ?", fmt.Sprintf("%%%s%%", value))
			continue
		}
		// 处理状态上active 且库存大于0的商品
		if key == "status" && value == "active" {
			query = query.Where("status = ? AND stock > ?", "active", 0)
			continue
		}

		if err := applyFilter(query, key, value); err != nil {
			return "", err
		}
	}

	var total int64
	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return Empty, err
	}
	if total <= 0 {
		return Empty, nil
	}
	v := reflect.ValueOf(datas)
	// 检查是否为指针，并解引用
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if params.Sort.Field == "" {

		query = query.Find(datas)
	} else {
		sort := fmt.Sprintf("%s %s", params.Sort.Field, params.Sort.Md)
		query = query.Order(sort).Offset(params.Ranges.Skip).Limit(params.Ranges.Limit).Find(datas)
	}

	rs := fmt.Sprintf("%d-%d/%d", params.Ranges.Skip, params.Ranges.Skip+v.Len(), total)

	return rs, query.Error

}
