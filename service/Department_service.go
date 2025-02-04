package service

import (
	"fmt"
	"goVueBlog/models"
)

var departmentService *DepartmentService

type DepartmentService struct {
	BaseService
}

func NewDepartmentService() *DepartmentService {
	if departmentService == nil {
		return &DepartmentService{
			BaseService: NewBaseApi(&models.Department{}),
		}
	}
	return departmentService
}

func (s *DepartmentService) CreateDepartment(department *models.Department) error {
	return s.DB.Create(department).Error
}

// 删除部门
func (s *DepartmentService) DeleteDepartment(id uint) error {
	return s.DB.Delete(&models.Department{}, id).Error
}

// 更新部门
func (s *DepartmentService) UpdateDepartment(id uint, data map[string]interface{}) error {
	return s.DB.Model(&models.Department{}).Where("id = ?", id).Updates(data).Error
}

// 查询部门
func (s *DepartmentService) GetDepartment(id uint) (models.Department, error) {
	var department models.Department
	if err := s.DB.First(&department, id).Error; err != nil {
		return department, fmt.Errorf("获取部门数据失败: %w", err)
	}
	return department, nil
}

// 查询所有部门
func (s *DepartmentService) GetAllDepartments() ([]models.Department, error) {
	var departments []models.Department
	if err := s.DB.Find(&departments).Error; err != nil {
		return departments, fmt.Errorf("获取所有部门数据失败: %w", err)
	}
	return departments, nil
}

// 获取父部门 ID
func (s *DepartmentService) GetParentDepartmentID(id uint) (*uint, error) {
	var department models.Department
	if err := s.DB.Select("parent_id").First(&department, id).Error; err != nil {
		return nil, fmt.Errorf("获取父部门ID失败: %w", err)
	}
	return department.ParentID, nil
}
