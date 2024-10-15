package service

import (
	"fmt"
	"goVueBlog/models"
)

var categoryService *CategoryService

type CategoryService struct {
	BaseService
}

func NewCateGoryService() *CategoryService {
	if categoryService == nil {
		return &CategoryService{
			BaseService: NewBaseApi(),
		}
	}
	return categoryService
}

// 创建分类,返回id或错误
func (m *CategoryService) CreateCategory(data *models.Category) error {
	return m.DB.Create(data).Error
}

// 根据页码数量查询条件获取分类列表
func (m *CategoryService) GetCategoryList(skip int, limit int, stroArr []string) (*[]models.Category, int64, error) {
	var categories []models.Category
	var totalCount int64
	// 统计总记录数
	if err := m.DB.Model(&models.Category{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}
	return &categories, totalCount, m.DB.Order(fmt.Sprintf("%s %s", stroArr[0], stroArr[1])).Offset(skip).Limit(limit).Find(&categories).Error

}

// 根据ID获取分类信息
func (m *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	return &category, m.DB.First(&category, id).Error
}

// 查询id是否存在
func (m *CategoryService) IsCategoryExistByID(id uint) bool {
	var category models.Category
	return m.DB.First(&category, id).Error == nil
}

// 根据ID删除分类
func (m *CategoryService) DeleteCategoryByID(id uint) error {
	return m.DB.Delete(&models.Category{}, id).Error
}

// 跟新分类信息
func (m *CategoryService) UpdateCategoryByID(id uint, data *models.Category) error {
	return m.DB.Model(&models.Category{}).Where("id = ?", id).Updates(data).Error
}
