package models

import (
	utils "goVueBlog/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;index;unique;" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
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
