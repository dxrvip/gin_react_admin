package service

import (
	"encoding/json"
	"fmt"
	"goVueBlog/globar"
	"goVueBlog/service/serializer"
	"reflect"
	"strings"
	"time"

	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type CustomDate time.Time

// UnmarshalJSON 反序列化
func (ct *CustomDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	// 无需时间部分格式是 "2025-04-30"

	t, err := time.Parse(time.DateOnly, s) // 支持带毫秒的 ISO 8601
	if err != nil {
		return fmt.Errorf("invalid ISO 8601 time format: %v", err)
	}
	*ct = CustomDate(t.UTC()) // 强制转为 UTC 时间
	return nil
}

// MarshalJSON 序列化（可选）
func (ct CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(ct).Format(time.RFC3339Nano))

}

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
func (m *BaseService) Create(params interface{}) error {
	// const msg string = "创建失败"
	// // 获取参数值的反射对象
	// v := reflect.ValueOf(params)
	// // 处理指针类型（如果传入的是指针。获取其指向的值
	// if v.Kind() == reflect.Ptr {
	// 	v = v.Elem()
	// }
	// // 更具不同的类型执行不同逻辑
	// switch v.Kind() {

	// case reflect.Struct:

	// 	// 结构体方法
	// 	structToMap := structToMap(params)
	// 	result := m.DB.Model(m.Model).Create(structToMap)
	// 	if result.Error != nil {
	// 		return nil, fmt.Errorf("%s : %v", msg, result.Error)
	// 	}
	// 	return result, nil
	// case reflect.Map:
	// 	result := m.DB.Model(m.Model).Create(params)
	// 	if result.Error != nil {
	// 		return nil, fmt.Errorf("%s2: %v", msg, result.Error)
	// 	}
	// 	return result, nil
	// default:
	// 	return nil, fmt.Errorf("数据类型不支持！%T， 仅支持 struct， 或map", params)
	// }
	result := m.DB.Model(m.Model).Create(params)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("重复添加错误: %v", result.Error)
		}
		return result.Error
	}
	return nil

}

// 将结构体转换为 map 的辅助函数
func structToMap(data interface{}) map[string]interface{} {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	dataMap := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		dataMap[field.Name] = v.Field(i).Interface()
	}
	return dataMap
}

// 查询所有数据
const Empty string = "0-0/0"

func (m *BaseService) List(datas interface{}, params *serializer.CommonQueryOtpones) (string, error) {
	// 添加查询条件
	query := m.DB.Model(m.Model)
	// 构建查询条件
	for key, value := range params.Filter {
		if err := applyFilter(query, key, value); err != nil {
			return "", err
		}
	}

	var total int64
	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return Empty, err
	}
	if total <= 0 {
		return Empty, nil
	}
	v := reflect.ValueOf(datas)
	// 检查是否为指针，并解引用
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if params.Sort.Field == "" {

		query = query.Find(datas)
	} else {
		sort := fmt.Sprintf("%s %s", params.Sort.Field, params.Sort.Md)
		query = query.Order(sort).Offset(params.Ranges.Skip).Limit(params.Ranges.Limit).Find(datas)
	}

	rs := fmt.Sprintf("%d-%d/%d", params.Ranges.Skip, params.Ranges.Skip+v.Len(), total)

	return rs, query.Error

}

// applyFilter 用于应用查询条件
func applyFilter(query *gorm.DB, key string, value interface{}) error {
	switch {
	case strings.HasPrefix(key, "q_"):
		fieldStr := key[2:]
		return query.Where(fieldStr+" LIKE ?", "%"+value.(string)+"%").Error
	case key == "id":
		v := reflect.ValueOf(value)
		data := v.Index(0).Interface()
		d := reflect.ValueOf(data)
		if d.Kind() == reflect.Array || d.Kind() == reflect.Slice {
			if d.Len() == 1 {
				return query.Where("id = ?", d.Index(0).Interface()).Error
			} else {
				values := make([]interface{}, d.Len())
				for i := 0; i < d.Len(); i++ {
					values[i] = d.Index(i).Interface()
				}
				return query.Where("id IN (?)", values).Error
			}
		} else {

			return query.Where(key+" IN ?", value).Error
		}
	case key == "id_ne":
		return query.Where("id != ?", value).Error
		// case key == "categories_id":
		// 多对多查询：通过中间表关联商品分类和属性键
		// return query.
		// 	Joins("JOIN category_attributes ON attribute.id = category_attributes.attribute_id").
		// 	Where("category_attributes.category_id = ?", value).
		// 	Error
		return nil

	default:
		return query.Where(key+" = ?", value).Error
	}
}

// 获取数据根据ID
func (m *BaseService) GetDataByID(id uint, datas interface{}) error {
	err := m.DB.Model(m.Model).First(&datas, id).Error
	return err

}

func (m *BaseService) Updates(id uint, params any) (any, error) {
	rs := m.DB.Model(m.Model).Where("id = ?", id).Updates(params)

	if rs.Error != nil {
		return nil, fmt.Errorf("更新失败！%v", params)
	}
	return params, nil
}

// 更新数据根据ID
func (m *BaseService) UpdateDataByID(id uint, data interface{}) error {
	rValue := reflect.ValueOf(data).Elem()
	ids := rValue.FieldByName("ID")
	ids.SetUint(uint64(id))
	// result := m.DB.Model(&m.Model).Save(rValue.Interface())
	result := m.DB.Model(m.Model).Where("id = ?", id).Updates(rValue.Interface())
	return result.Error
}

// 根据ID删除分类
func (m *BaseService) DeleteByID(id uint) error {
	result := m.DB.Delete(m.Model, id)
	return result.Error
}
