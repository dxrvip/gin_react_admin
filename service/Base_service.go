package service

import (
	"errors"
	"goVueBlog/globar"

	"gorm.io/gorm"
)

type BaseService struct {
	DB *gorm.DB
}

func NewBaseApi() BaseService {
	return BaseService{
		DB: globar.DB,
	}
}

// 添加数据
func (m *BaseService) Create(params *map[string]interface{}) (err error) {
	err = m.DB.Create(params).Error
	return
}

// 查询所有数据
func (m *BaseService) List(model interface{}, datas interface{}, limit int, skip int, stroArr []string) (interface{}, int64, error) {
	var total int64
	// 计算总数
	if err := m.DB.Model(model).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 验证排序参数
	if len(stroArr) != 2 {
		return nil, 0, errors.New("invalid sort parameters")
	}
	// 分页和排序查询
	if err := m.DB.Model(model).Order(stroArr[0] + " " + stroArr[1]).Offset(skip).Limit(limit).Find(&datas).Error; err != nil {
		return nil, 0, err
	}

	return datas, total, nil
}
