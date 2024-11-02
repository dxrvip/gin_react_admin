package service

import (
	"errors"
	"goVueBlog/models"
	"goVueBlog/utils"
)

var userService *UserService

type UserService struct {
	BaseService
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

type RegisterData struct {
	LoginRequest
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	NikeName   string `json:"nikeName" binding:"min=2,max=50" label:"昵称"`
	Email      string `json:"email" binding:"email" label:"邮箱"`
	Active     bool   `json:"active" binding:"required" label:"状态"`
	Gender     string `json:"gender" binding:"required" label:"性别"`
}
type ResponseUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username" validate:"required,min=4,max=12" label:"用户名"`
	NikeName string `json:"nike_name" validate:"min=2,max=50" label:"昵称"`
	Email    string `json:"email" validate:"usage=email" label:"邮箱"`
	Active   bool   `json:"active" validate:"required" label:"状态"`
	Gender   string `json:"gender" label:"性别"`
}

func NewUserService() *UserService {
	if userService == nil {
		return &UserService{
			BaseService: NewBaseApi(models.User{}),
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
	return m.DB.Delete(&models.User{}, id).Error
}

// 修改用户信息
func (m *UserService) UpdateUser(id uint, username string) error {
	return m.DB.Where("id = ?", id).Update("username", username).Error
}
