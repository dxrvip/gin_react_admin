package models

import "fmt"

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name" validate:"required,min(3),max(20)" label:"名称"`
}

// 创建分类,返回id或错误
func CreateCategory(data *Category) error {
	return Db.Create(data).Error
}

// 根据页码数量查询条件获取分类列表
func GetCategoryList(skip int, limit int, stroArr []string) (*[]Category, int64, error) {
	var categories []Category
	var totalCount int64
	// 统计总记录数
	if err := Db.Model(&Category{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}
	return &categories, totalCount, Db.Order(fmt.Sprintf("%s %s", stroArr[0], stroArr[1])).Offset(skip).Limit(limit).Find(&categories).Error

}

// 根据ID获取分类信息
func GetCategoryByID(id uint) (*Category, error) {
	var category Category
	return &category, Db.First(&category, id).Error
}

// 查询id是否存在
func IsCategoryExistByID(id uint) bool {
	var category Category
	return Db.First(&category, id).Error == nil
}

// 根据ID删除分类
func DeleteCategoryByID(id uint) error {
	return Db.Delete(&Category{}, id).Error
}

// 跟新分类信息
func UpdateCategoryByID(id uint, data *Category) error {
	return Db.Model(&Category{}).Where("id = ?", id).Updates(data).Error
}
