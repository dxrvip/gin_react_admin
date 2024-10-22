package service

import (
	"goVueBlog/models"
)

var categoryService *CategoryService

type CategoryService struct {
	BaseService
}

func NewCateGoryService() *CategoryService {
	if categoryService == nil {
		return &CategoryService{
			BaseService: NewBaseApi(models.Category{}),
		}
	}
	return categoryService
}

// 创建分类,返回id或错误
func (m *CategoryService) CreateCategory(data *models.Category) error {
	return m.DB.Create(data).Error
}

// 查询id是否存在
func (m *CategoryService) IsCategoryExistByID(id uint) bool {
	var category models.Category
	return m.DB.First(&category, id).Error == nil
}
