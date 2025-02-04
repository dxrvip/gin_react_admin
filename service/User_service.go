package service

import (
	"errors"
	"fmt"
	"goVueBlog/models"
	"goVueBlog/utils"
	"log"
)

var userService *UserService

type UserService struct {
	BaseService
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}
type EditUser struct {
	NikeName string            `json:"nike_name" binding:"min=2,max=50" label:"昵称"`
	Email    string            `json:"email" binding:"email" label:"邮箱"`
	Active   bool              `json:"status,omitempty" label:"状态"`
	Gender   models.GenderType `json:"gender,omitempty" label:"性别"`
}

type UpdateUser struct {
	EditUser
	Username string `json:"username" binding:"required,min=3,max=50"`
}

type RegisterData struct {
	LoginRequest
	EditUser
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ResponseUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username" validate:"required,min=4,max=12" label:"用户名"`
	NikeName string `json:"nike_name" validate:"min=2,max=50" label:"昵称"`
	Email    string `json:"email" validate:"usage=email" label:"邮箱"`
	Active   bool   `json:"status" validate:"required" label:"状态"`
	Gender   string `json:"gender" label:"性别"`
}

func NewUserService() *UserService {

	if userService == nil {
		return &UserService{
			BaseService: NewBaseApi(&models.User{}),
		}
	}
	return userService
}

func (m *UserService) CreateUser(p *models.User) (any, error) {
	// 验证用户名是否存在
	user, _ := m.GetUserByUsername(p.Username)
	if user.Username != "" {
		return nil, errors.New("用户名已存在")
	}

	// 存入数据库
	if err := m.DB.Create(&p).Error; err != nil {
		return nil, err
	}

	// 生成token
	token, err := utils.GenerateToken(p.Username, int(user.ID))
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	return token, nil
}

// 判断用户名是否存在
func (m *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := m.DB.Where("username = ?", username).First(&user)
	return &user, result.Error
}

// 删除用户
func (m *UserService) DeleteUser(id uint) error {
	return m.DB.Model(&models.User{}).Delete(&models.User{}, id).Error
}

// 修改用户信息
func (m *UserService) UpdateUser(id uint, data UpdateUser) error {
	// 使用map更新数据
	activeValue := 0
	if data.Active {
		activeValue = 1
	}
	result := m.DB.Model(&models.User{}).Where("id = ?", id).Updates(
		map[string]interface{}{
			"username":  data.Username,
			"nike_name": data.NikeName,
			"email":     data.Email,
			"active":    activeValue,
			"gender":    models.GenderType(data.Gender),
		},
	)
	// 检查受影响的行数
	if result.RowsAffected == 0 {
		log.Printf("No rows were updated for user ID: %d", id)
		return fmt.Errorf("no rows were updated")
	}
	return result.Error
}
