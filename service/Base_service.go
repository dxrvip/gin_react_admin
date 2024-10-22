package service

import (
	"fmt"
	"goVueBlog/globar"
	"goVueBlog/service/serializer"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type BaseService struct {
	DB    *gorm.DB
	Model interface{}
}

func NewBaseApi(model interface{}) BaseService {
	return BaseService{
		DB:    globar.DB,
		Model: model,
	}
}

// 添加数据
func (m *BaseService) Create(params interface{}) (mapData map[string]any, err error) {
	mapData = map[string]any{}

	v := reflect.ValueOf(params)
	t := reflect.TypeOf(params)
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		mapData[field.Name] = value.Interface()
	}

	err = m.DB.Model(m.Model).Create(&mapData).Error
	return
}

// 查询所有数据
func (m *BaseService) List(datas interface{}, params *serializer.CommonQueryOtpones) (int64, error) {
	// 添加查询条件
	fmt.Println(params)
	query := m.DB.Model(&m.Model)
	for key, value := range params.Filter {
		fmt.Printf("%s: %v\n", key, value)
		if strings.HasPrefix(key, "q_") {
			fieldStr := key[2:]
			query = query.Where(fieldStr+" LIKE ?", "%"+value.(string)+"%")
		} else if key == "id" {
			query = query.Where(key+" in ?", value)
		} else {
			query = query.Where(key+" = ?", value)
		}
	}

	var total int64
	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}
	if total <= 0 {
		return total, nil
	} else {
		if params.Sort.Field == "" {

			query = query.Find(datas)
		} else {
			sort := fmt.Sprintf("%s %s", params.Sort.Field, params.Sort.Md)
			query = query.Order(sort).Offset(params.Ranges.Skip).Limit(params.Ranges.Limit).Find(datas)
		}
		return total, query.Error
	}

}

// 获取数据根据ID
func (m *BaseService) GetDataByID(id uint, datas interface{}) error {
	err := m.DB.Model(&m.Model).First(&datas, id).Error
	return err

}

// 更新数据根据ID
func (m *BaseService) UpdateDataByID(id uint, data interface{}) error {
	rValue := reflect.ValueOf(data).Elem()
	ids := rValue.FieldByName("ID")
	ids.SetUint(uint64(id))
	result := m.DB.Model(&m.Model).Where("id = ?", id).Updates(rValue.Interface())
	return result.Error
}

// 根据ID删除分类
func (m *BaseService) DeleteByID(id uint) error {
	result := m.DB.Delete(&m.Model, id)
	return result.Error
}
