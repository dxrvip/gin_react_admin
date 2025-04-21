package service

import (
	"fmt"
	"goVueBlog/models"
	"goVueBlog/service/serializer"
	"reflect"

	"gorm.io/gorm"
)

// Key AttributeKey
var attributeService *AttributeService

type AttributeService struct {
	BaseService
}

// 公共基础结构体
type BaseAttribute struct {
	Name         string               `json:"name" binding:"required"`
	Type         models.AttributeType `json:"type" binding:"required,oneof=string number enum"`
	IsRequired   bool                 `json:"isRequired"`
	DefaultValue string               `json:"defaultValue"`
	MinValue     *float64             `json:"min"`
	MaxValue     *float64             `json:"max"`
	Options      models.OptionList    `json:"options"`
	Unit         string               `json:"unit"`
}

// Request 结构体
type Request struct {
	BaseAttribute
	CategoryIDs []uint `json:"categoryIds"`
}

// ListAttribute 结构体
type ListAttribute struct {
	ID uint `json:"id"`
	BaseAttribute
	Categories []*models.ProductCategory `json:"categories"`
}

// InfoAttribute 结构体
type InfoAttribute struct {
	ID uint `json:"id"`
	BaseAttribute
	CategoryIDs []uint `json:"categoryIds"`
}

func NewAttributeKeyService() *AttributeService {
	if attributeService == nil {
		return &AttributeService{
			BaseService: NewBaseApi(&models.Attribute{}),
		}
	}
	return attributeService
}
func (m *AttributeService) Create(params interface{}) error {
	// 类型断言获取具体属性对象
	attr, ok := params.(*models.Attribute)
	if !ok {
		return fmt.Errorf("无效的参数类型")
	}

	// 开启事务
	tx := m.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 创建属性主体
	if err := tx.Create(attr).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 处理分类关联（当有分类ID时）
	if len(attr.CategoryIDs) > 0 {
		// 构建中间表数据
		var categoryAttributes []models.CategoryAttribute
		for _, cid := range attr.CategoryIDs {
			categoryAttributes = append(categoryAttributes, models.CategoryAttribute{
				ProductCategoryID: cid,
				AttributeID:       attr.ID,
			})
		}

		// 批量创建关联
		if err := tx.Create(&categoryAttributes).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
func (s *AttributeService) GetAttributeDetails(id uint, info *InfoAttribute) error {
	// 1. 预加载分类关联数据
	var attr models.Attribute
	// 使用 Debug() 输出完整 SQL 日志
	// query := s.DB.Model(s.Model).
	// 	Preload("Categories").
	// 	First(&attr, id)
	// if err := query.Error; err != nil {
	// 	return fmt.Errorf("查询失败: %v", err)
	// }
	query := s.DB.Debug().Model(s.Model).
		Preload("Categories", func(db *gorm.DB) *gorm.DB {
			// 可选：忽略关联模型的软删除
			return db.Unscoped()
		}).
		Where("id = ?", id).
		First(&attr) // 使用 First 替代 Find，避免切片处理

	if err := query.Error; err != nil {
		return fmt.Errorf("查询失败: %v", err)
	}

	// 提取分类 ID
	// var categoryIDs []uint
	for _, category := range attr.Categories {
		attr.CategoryIDs = append(attr.CategoryIDs, category.ID)
	}
	// attr.CategoryIDs = categoryIDs // 回填到前端字段
	// 2. 构建返回数据
	*info = InfoAttribute{
		ID: attr.ID,
		BaseAttribute: BaseAttribute{
			Name:         attr.Name,
			Type:         attr.Type,
			IsRequired:   attr.IsRequired,
			DefaultValue: attr.DefaultValue,
			MinValue:     attr.MinValue,
			MaxValue:     attr.MaxValue,
			Options:      attr.Options,
			Unit:         attr.Unit,
		},
		CategoryIDs: attr.CategoryIDs,
	}
	// 3. 返回结果
	return nil
}

func (s *AttributeService) Updates(id uint, data *models.Attribute) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 更新基础字段（强制包含零值）
		if err := tx.Model(&models.Attribute{}).
			Where("id = ?", id).
			Updates(data).
			Error; err != nil {
			return err
		}

		// 2. 处理多对多关联分类
		// 初始化空切片（即使无分类也确保非 nil）
		var productCategory []*models.ProductCategory

		// 只有当 CategoryIDs 非空时才查询数据库
		if len(data.CategoryIDs) > 0 {
			// 查询分类是否存在
			if err := tx.Where("id IN (?)", data.CategoryIDs).Find(&productCategory).Error; err != nil {
				return fmt.Errorf("分类查询失败: %v", err)
			}
		} else {
			// 显式赋值为空切片（非 nil）
			productCategory = make([]*models.ProductCategory, 0)
		}

		// 3. 替换关联（确保操作在正确的模型上）
		attr := &models.Attribute{BaseModel: models.BaseModel{ID: id}}
		if err := tx.Model(attr).Association("Categories").Replace(productCategory); err != nil {
			return fmt.Errorf("关联分类更新失败: %v", err)
		}

		return nil
	})
}

func (s *AttributeService) List(datas interface{}, params *serializer.CommonQueryOtpones) (string, error) {
	// 初始化查询器并预加载分类信息 // 预加载分类信息限制 name和 id 字段

	query := s.DB.Model(s.Model).
		Preload("Categories", func(tx *gorm.DB) *gorm.DB {
			// 可选：忽略关联模型的软删除
			return tx.Select("id, name").Unscoped()
		})
		// Joins("LEFT JOIN category_attributes ON attributes.id = category_attributes.attribute_id")

	// 应用过滤条件
	for key, value := range params.Filter {
		if key == "categories_id" {
			// 处理分类过滤的特殊逻辑
			query = query.
				Joins("JOIN category_attributes ON category_attributes.attribute_id = attributes.id").
				Where("category_attributes.product_category_id IN (?)", value).
				Distinct()
			continue
		}
		if err := applyFilter(query, key, value); err != nil {
			return "", err
		}
	}

	// 保留基础分页排序逻辑
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return Empty, err
	}

	if params.Sort.Field != "" {
		query = query.Order(fmt.Sprintf("%s %s", params.Sort.Field, params.Sort.Md))
	}

	if err := query.Offset(params.Ranges.Skip).
		Limit(params.Ranges.Limit).
		Find(datas).Error; err != nil {
		return Empty, err
	}

	// 生成分页标识
	v := reflect.ValueOf(datas).Elem()
	return fmt.Sprintf("%d-%d/%d",
		params.Ranges.Skip,
		params.Ranges.Skip+v.Len(),
		total), nil
}
