package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	BaseModel

	Creator uint   `gorm:"not null" json:"creator"` // 创建人ID
	Content string `gorm:"type:text;not null;comment:消息内容"`
}

// UserMessage 结构体，用于关联用户消息
type UserMessage struct {
	BaseModel
	UserID    uint      `gorm:"not null;comment:用户ID" json:"user_id"`
	MessageID uint      `gorm:"not null;comment:消息ID" json:"message_id"`
	IsRead    bool      `gorm:"default:false;comment:是否已读" json:"is_read"`
	ReadAt    time.Time `gorm:"comment:阅读时间" json:"read_at,omitempty"`
}

// MarkAsRead 标记消息为已读
func (um *UserMessage) MarkAsRead(tx *gorm.DB) error {
	um.IsRead = true
	um.ReadAt = time.Now()
	return tx.Save(um).Error
}
