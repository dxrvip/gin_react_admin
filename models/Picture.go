package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Picture struct {
	Src   string `json:"src" binding:"omitempty,url"`
	Title string `json:"title" binding:"max=100"`
}

// PictureList 用于处理图片数组的序列化和反序列化
type PictureList []Picture

// Value 实现数据库序列化的 driver.Valuer 接口
func (pl PictureList) Value() (driver.Value, error) {
	return json.Marshal(pl)
}

// Scan 实现数据库反序列化的 sql.Scanner 接口
func (pl *PictureList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	// 尝试直接解析为新格式
	if err := json.Unmarshal(bytes, pl); err != nil {
		// 如果解析失败，尝试解析为旧格式
		var oldFormat []struct {
			Src   string `json:"src"`
			Title string `json:"title"`
		}
		if err := json.Unmarshal(bytes, &oldFormat); err != nil {
			return err
		}
		// // 转换为新格式
		// *pl = make(PictureList, len(oldFormat))
		// for i, old := range oldFormat {
		// 	(*pl)[i] = Picture{
		// 		Image: ImageData{
		// 			Src:   old.Src,
		// 			Title: old.Title,
		// 		},
		// 	}
		// }
	}
	return nil
}

// Value 实现数据库序列化的 driver.Valuer 接口
func (p Picture) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan 实现数据库反序列化的 sql.Scanner 接口
func (p *Picture) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	// 尝试直接解析为 Picture 结构体
	if err := json.Unmarshal(bytes, p); err != nil {
		// 如果解析失败，可能是旧格式，尝试兼容处理
		var oldFormat struct {
			Src   string `json:"src"`
			Title string `json:"title"`
		}
		if err := json.Unmarshal(bytes, &oldFormat); err != nil {
			return err
		}
		// // 转换为新格式
		// p.Image = ImageData{
		// 	Src:   oldFormat.Src,
		// 	Title: oldFormat.Title,
		// }
	}
	return nil
}
