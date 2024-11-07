package service

import (
	"fmt"
	"goVueBlog/models"
	"reflect"

	"gorm.io/gorm"
)

var roleService *RoleService

type RoleService struct {
	BaseService
}

type RoleParams struct {
	Name   string `json:"name" binding:"required,min=2,max=50"`
	Key    string `json:"key" binding:"omitempty,max=50,min=3" label:"权限标识符"`
	Sort   uint   `json:"sort" binding:"omitempty" label:"排序顺序"`
	Active bool   `json:"active" binding:"omitempty" label:"是否启用"`
}
type RoleResponse struct {
	RoleParams
	ID    uint              `json:"id" binding:"omitempty" label:"id"`
	Menus models.JSONString `json:"menus,omitempty" binding:"omitempty" label:"权限菜单"`
	User  []uint            `json:"user,omitempty"`
}

type UpdateParams struct {
	ID     uint              `json:"id,omitempty"`
	Name   string            `json:"name,omitempty" binding:"omitempty,min=2,max=50"`
	Key    string            `json:"key,omitempty" binding:"omitempty,max=50,min=3" validate:"omitempty"`
	Sort   uint              `json:"sort,omitempty" label:"排序顺序"`
	Active bool              `json:"active,omitempty" label:"是否启用"`
	Menus  models.JSONString `gorm:"type:string" json:"menus,omitempty" binding:"omitempty" label:"权限菜单"`
	User   []uint            `json:"user"`
}

func NewRoleService() *RoleService {
	if roleService == nil {
		return &RoleService{
			BaseService: NewBaseApi(&models.Role{}),
		}
	}
	return roleService
}

func (m *RoleService) GetUsersById(id uint, data interface{}) ([]models.User, error) {
	volue := reflect.ValueOf(data).Elem()
	userSlip := volue.FieldByName("User")

	var mapUser []models.User
	if userSlip.IsValid() && userSlip.Kind() == reflect.Slice {
		// 循环拿出user对象
		for i := 0; i < userSlip.Len(); i++ {
			userId := userSlip.Index(i).Uint()
			// 重数据库中拿出用户信息
			var thisUser models.User
			if err := m.DB.Where("id = ?", userId).Find(&thisUser).Error; err != nil {
				return mapUser, fmt.Errorf("获取id为：%d 失败！", userId)
			}
			mapUser = append(mapUser, thisUser)
		}
	}

	// 循环取出userId，重数据库中取出，赋值给user，
	return mapUser, nil
}

// 关联插入
func (m *RoleService) UpdateUserAndRoleDataByID(datas *models.Role) error {
	result := m.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(datas)
	return result.Error

}

// 获取数据根据ID
func (m *RoleService) GetDataByID(id uint) (models.Role, error) {
	// 先查询 role
	var role models.Role

	// 查询角色及其关联的用户字段
	if err := m.DB.Model(&m.Model).Preload("Users", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username") // 选择 Users 中的 id 和 username
	}).First(&role, id).Error; err != nil {
		return role, fmt.Errorf("获取角色失败: %w", err)
	}

	return role, nil
}
