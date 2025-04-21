package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type AttributeMap map[string]interface{}

func (a *AttributeMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid database type")
	}
	return json.Unmarshal(bytes, a)
}

func (a AttributeMap) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// 商品分类（支持多级）
type ProductCategory struct {
	BaseModel
	Name        string `gorm:"type:varchar(100);not null;unique;comment:分类名称" json:"name" validate:"required,min=2,max=100"`
	Description string `gorm:"type:varchar(100);comment:分类描述" json:"description" validate:"omitempty,max=500"`
	ParentID    *uint  `gorm:"index;comment:父分类ID" json:"parentId,omitempty"`
	Order       int    `gorm:"default:0;comment:排序" json:"order"`
	IsEnabled   bool   `gorm:"default:true;comment:是否启用" json:"isEnabled"`
	// Attributes  []*Attribute `gorm:"many2many:category_attributes;joinForeignKey:product_category_id;joinReferencesKey:attribute_id"`
	Attributes []*Attribute `gorm:"many2many:category_attributes;" json:"attributes"`
}

// 商品信息
type Product struct {
	BaseModel
	Title             string          `gorm:"type:varchar(255);not null" json:"title" validate:"required"`
	Price             Float64String   `gorm:"type:decimal(10,2);not null" json:"price" validate:"gt=0"`
	Stock             uint            `gorm:"not null" json:"stock" validate:"gte=0"`
	Description       string          `gorm:"type:text" json:"description"`
	Images            PictureList     `gorm:"type:json" json:"picture"`
	ProductCategoryID uint            `gorm:"not null" json:"productCategoryID"`
	ProductCategory   ProductCategory `gorm:"foreignKey:ProductCategoryID" json:"category"`
	Status            string          `gorm:"type:varchar(20);default:'active'" json:"status"`

	BrandID    uint         `gorm:"not null" json:"brandID"`
	Brand      Brand        `gorm:"foreignKey:BrandID" json:"brand,omitempty"`
	Attributes AttributeMap `gorm:"type:json" json:"attributes"`
	// SecondHandSku 二手商品字段
	SecondHandSku []SecondHandSku `gorm:"foreignKey:ProductID" json:"secondHandSku,omitempty"`
}

// AfterFind 钩子用于处理 JSON 字段的反序列化
func (p *Product) AfterFind(tx *gorm.DB) error {
	// 处理属性字段

	// 处理状态
	switch p.Status {
	case "active":
		p.Status = "发布"
	case "disabled":
		p.Status = "草稿"
	case "pulled":
		p.Status = "下架"
	default:
		return nil

	}
	return nil
}

// BeforeSave 钩子用于处理图片数据的序列化
func (p *Product) BeforeSave(tx *gorm.DB) error {
	// 如果 Images 为空，确保它是一个空数组而不是 nil
	if p.Images == nil {
		p.Images = make(PictureList, 0)
	}
	return nil
}

// 商品sku数据库 二手商品
// 成色。全新，几乎全新，轻微痕迹，明显痕迹，严重痕迹，破损
// 功能，功能完好无维修，维修过，可以正常使用，无法正常使用
// 已使用年限，未使用，6个月内，6个月-1年，1-3年
// 配件，配件齐全，配件缺失
// 运费，包邮，不包邮

type ProductSku struct {
	BaseModel
	ProductID uint    `gorm:"not null" json:"productId"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	Sku       string  `gorm:"type:varchar(100);not null" json:"sku"`
}

// 二手商品成色枚举
const (
	ConditionNew     = "new"      // 全新
	ConditionLikeNew = "like_new" // 几乎全新
	ConditionLight   = "light"    // 轻微痕迹
	ConditionObvious = "obvious"  // 明显痕迹
	ConditionSerious = "serious"  // 严重痕迹
	ConditionDamaged = "damaged"  // 破损
)

// 二手商品功能状态枚举
const (
	FunctionPerfect  = "perfect"  // 功能完好无维修
	FunctionRepaired = "repaired" // 维修过
	FunctionUsable   = "usable"   // 可以正常使用
	FunctionUnusable = "unusable" // 无法正常使用
)

// 使用年限枚举
const (
	UsageUnused     = "unused"      // 未使用
	UsageHalfYear   = "half_year"   // 6个月内
	UsageOneYear    = "one_year"    // 6个月-1年
	UsageThreeYears = "three_years" // 1-3年
)

// 商品状态枚举
// 定义商品状态枚举类型
type ProductStatus string

const (
	StatusActive ProductStatus = "active" // 发布
	StatusDraft                = "draft"  // 草稿
	StatusPulled               = "pulled" // 下架
)

// 二手商品SKU
type SecondHandSku struct {
	BaseModel
	Title         string  `gorm:"type:varchar(60);" json:"title"`
	ProductID     uint    `gorm:"not null" json:"productId"`                      // 关联主商品ID
	Price         float64 `gorm:"type:decimal(10,2);not null" json:"price"`       // SKU价格
	Stock         uint    `gorm:"not null" json:"stock"`                          // 库存
	Condition     string  `gorm:"type:varchar(20);not null" json:"condition"`     // 成色
	Function      string  `gorm:"type:varchar(20);not null" json:"function"`      // 功能状态
	UsageDuration string  `gorm:"type:varchar(20);not null" json:"usageDuration"` // 使用年限
	Accessories   bool    `gorm:"default:true" json:"accessories"`                // 配件是否齐全
	// 缺少的配件
	AccessoriesList string        `gorm:"vachar(255);default:nill" json:"accessoriesList"` // 缺少配件信息
	FreeShipping    bool          `gorm:"default:false" json:"freeShipping"`               // 是否包邮
	ShippingFee     float64       `gorm:"type:decimal(10,2);default:0" json:"shippingFee"` // 运费金额（不包邮时的运费）
	Description     string        `gorm:"type:text" json:"description"`                    // SKU描述
	Images          PictureList   `gorm:"type:json" json:"picture"`                        // SKU图片
	Product         Product       `gorm:"foreignKey:ProductID" json:"product"`             // 关联主商品
	Cost            Float64String `gorm:"decimal(10,2);default:0" json:"cost"`             // 成本价格
	ProductsType    string        `gorm:"varchar(20);not null" json:"productsType"`        //货号
	ProductSkuID    string        `gorm:"varchar(100);not null" json:"productSkuID"`       // 商品SKU ID
	// 是否再保
	IsRepair      bool      `gorm:"default:false" json:"isRepair"`           // 是否维修过
	RepairEndDate time.Time `gorm:"autoCreateTime:int" json:"repairEndDate"` // 维修结束时间
	// 商品状态
	Status ProductStatus `gorm:"type:varchar(20);default:'active'" json:"status"` // 商品状态
	// 电池工作时间
	BatteryLife uint `gorm:"default:30" json:"batteryLife"` // 电池工作时间（单位： 分钟）
	// 机器工作时间记录
	WorkingTime uint `gorm:"default:0" json:"workingTime"` // 维修后电池工作时间（单位：分钟）
}
