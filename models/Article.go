package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"goVueBlog/serializer"

	"gorm.io/gorm"
)

type Picture struct {
	Src   string `gorm:"type:varchar(100)" json:"src" binding:"omitempty,url"`
	Title string `json:"title" binding:"max=100"`
}

// Value实现数据库序列化的driver.Valuer接口
func (p Picture) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan实现数据库反序列化的sql.Scanner接口
func (p *Picture) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, p)
}

type Article struct {
	Category Category `gorm:"foreignKey:Cid"`
	gorm.Model
	ID           uint    `gorm:"primarykey" json:"id"`
	Title        string  `gorm:"type:varchar(100);not null"`
	Cid          int     `gorm:"type:int;not null"`
	Desc         string  `gorm:"type:varchar(200)"`
	Content      string  `gorm:"type:longtext"`
	Picture      Picture `gorm:"type:json"`
	CommentCount int     `gorm:"type:int;not null;default:0"`
	ReadCount    int     `gorm:"type:int;not null;default:0"`
}

// 添加数据
func CreatePost(data *serializer.PostRequry) (err error) {

	return Db.Model(&Article{}).Create(&data).Error
}

// 查询所有数据
func PostList(limit, skip int, stroArr []string) (*[]Article, int64, error) {
	var posts []Article
	var total int64

	// 计算总数
	if err := Db.Model(&Article{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 验证排序参数
	if len(stroArr) != 2 {
		return nil, 0, errors.New("invalid sort parameters")
	}

	// 分页和排序查询
	return &posts, total, Db.Order(stroArr[0] + " " + stroArr[1]).Offset(skip).Limit(limit).Find(&posts).Error
}

// 根据id查询文章
func GetPostById(id int) (*Article, error) {
	var post Article
	return &post, Db.First(&post, id).Error
}

// 更新文章
func UpdatePost(id int, data *Article) error {
	return Db.Model(&Article{}).Where("id = ?", id).Updates(data).Error
}

// 删除文章
func DeletePost(id int) error {
	return Db.Where("id = ?", id).Delete(&Article{}).Error
}
