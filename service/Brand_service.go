package service

import "goVueBlog/models"

var brandService *BrandService

type BrandService struct {
	BaseService
}

func NewBrandService() *BrandService {
	if brandService == nil {
		return &BrandService{
			BaseService: NewBaseApi(&models.Brand{}),
		}
	}
	return brandService
}

type BrandCreateData struct {
	Name        string `json:"name" binding:"required,min=2,max=255" label:"品牌名称"` // 品牌名称
	Logo        string `json:"logo,omitempty" binding:"omitempty,max=255"`         // Logo URL
	Description string `json:"description,omitempty" binding:"omitempty,max=255"`  // 品牌描述
}

type BrandResponse struct {
	BrandCreateData
	Id string `json:"id"`
}
