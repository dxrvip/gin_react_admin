package models

import (
	utils "goVueBlog/utils"

	"gorm.io/gorm"
)

// / 定义性别类型
type GenderType string

const (
	Male    GenderType = "男"
	Female  GenderType = "女"
	Unknown GenderType = "未知"
)

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(20);not null;index;unique;" json:"username,omitempty" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password,omitempty" validate:"required,min=6,max=120" label:"密码"`
	Roles    []Role `gorm:"many2many:role_user;comment: 权限ID" json:"roles,omitempty"`
	NikeName string `gorm:"type:varchar(50);default:null" json:"nike_name,omitempty" validate:"min=2,max=50" label:"昵称"`
	Email    string `gorm:"type:varchar(50);default:null" json:"email,omitempty" validate:"usage=email" label:"邮箱"`
	Active   bool   `gorm:"type:tinyint(1);default:0" json:"status,omitempty"  label:"状态"`
	// Active   bool       `gorm:"default:true" json:"status,omitempty" validate:"required" label:"状态"`
	IsSuper bool       `gorm:"default:false" label:"是否超级管理员"`
	Gender  GenderType `gorm:"type:enum('男', '女', '未知');default:'未知'" label:"性别"`

	DepartmentID *uint      `gorm:"index" json:"department_id"` // 允许为空
	Department   Department `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
}

// 插入之前进行加密
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = utils.EncryptPassword(u.Password)
	return
}

// 获取密码
func (u *User) GetPassword() string {
	return "****"
}

// 获取数据对性别修改
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	switch u.Gender {
	case "male":
		u.Gender = "男"
	case "female":
		u.Gender = "女"
	default:
		u.Gender = "未知"
	}
	return
}
