package service

import "goVueBlog/models"

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
