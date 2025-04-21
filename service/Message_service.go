package service

import (
	"goVueBlog/models"
	"time"

	"gorm.io/gorm"
)

type CreateData struct {
	Title string `json:"title" binding:"required,omitempty"`
	// Creator uint   `json:"creator,omitempty"` // 创建人ID
	Content string `json:"content" binding:"required"`
}
type ResponseMessage struct {
	Title     string    `json:"title"`
	Creator   uint      `json:"creator"` // 创建人ID
	Content   string    `json:"content"`
	UserName  string    `gorm:"-" json:"userName"`
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createTime"`
}

// AfterFind 钩子方法，在查询消息后只返回 User 的 name 字段
func (m *ResponseMessage) AfterFind(tx *gorm.DB) (err error) {
	var user models.User
	if err := tx.Select("username").Where("id = ?", m.Creator).First(&user).Error; err != nil {
		return err
	}
	m.UserName = user.Username
	return nil
}

var messageService *MessageService

type MessageService struct {
	BaseService
}

func NewMessageService() *MessageService {
	if messageService == nil {
		return &MessageService{
			BaseService: NewBaseApi(&models.Message{}),
		}
	}
	return messageService
}

func (m *MessageService) CreateMessage(data *models.Message) (interface{}, error) {
	err := m.DB.Create(data).Error
	return data, err
}
