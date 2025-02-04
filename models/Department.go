package models

type Department struct {
	BaseModel
	Name     string       `gorm:"type:varchar(100);not null;unique" json:"name" validate:"required" label:"部门名称"`
	Creator  uint         `gorm:"not null" json:"creator"`                       // 创建人ID
	ParentID *uint        `gorm:"index" json:"parent_id"`                        // 父部门ID
	Parent   *Department  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`   // 父部门
	Children []Department `gorm:"foreignKey:ParentID" json:"children,omitempty"` // 子部门
	Users    []User       `gorm:"foreignKey:DepartmentID" json:"users,omitempty"`
	//Messages []Message    `gorm:"foreignKey:DepartmentID" json:"messages,omitempty"`
}
