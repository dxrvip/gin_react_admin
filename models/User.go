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

func CreateUser(user *User) error {
	return Db.Create(user).Error
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

// 判断用户名是否存在
func GetUserByUsername(username string) (*User, error) {
	var user User
	result := Db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return &User{}, nil
		}
		return &User{}, result.Error
	}
	return &user, nil
}

// 删除用户
func DeleteUser(id uint) error {
	return Db.Delete(&User{}, id).Error
}

// 修改用户信息
func UpdateUser(id uint, username string) error {
	return Db.Where("id = ?", id).Update("username", username).Error
}
