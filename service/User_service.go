package service

import (
	"goVueBlog/models"

	"gorm.io/gorm"
)

var userService *UserService

type UserService struct {
	BaseService
}

func NewUserService() *UserService {
	if userService == nil {
		return &UserService{
			BaseService: NewBaseApi(),
		}
	}
	return userService
}

// 判断用户名是否存在
func (m *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := m.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return &models.User{}, nil
		}
		return &models.User{}, result.Error
	}
	return &user, nil
}

// 删除用户
func (m *UserService) DeleteUser(id uint) error {
	return m.DB.Delete(&models.User{}, id).Error
}

// 修改用户信息
func (m *UserService) UpdateUser(id uint, username string) error {
	return m.DB.Where("id = ?", id).Update("username", username).Error
}

func (m *UserService) CreateUser(user *models.User) error {
	return m.DB.Create(user).Error
}
