package service

import (
	"goVueBlog/models"
)

var roleService *RoleService

type RoleService struct {
	BaseService
}

type RoleParams struct {
	Name   string `json:"name" binding:"required,min=2,max=50"`
	Key    string `json:"key" binding:"omitempty,max=50,min=3" label:"权限标识符"`
	Sort   uint   `json:"sort" binding:"omitempty" label:"排序顺序"`
	Active bool   `json:"active" binding:"omitempty" label:"是否启用"`
}
type RoleResponse struct {
	ID    uint   `json:"id" binding:"omitempty" label:"id"`
	Menus string `json:"menus,omitempty" binding:"omitempty" label:"权限菜单"`
	RoleParams
}
type UpdateParams struct {
	ID     uint   `json:"id,omitempty"`
	Name   string `json:"name,omitempty" binding:"omitempty,min=2,max=50"`
	Key    string `json:"key,omitempty" binding:"omitempty,max=50,min=3" validate:"omitempty"`
	Sort   uint   `json:"sort,omitempty" label:"排序顺序"`
	Active bool   `json:"active,omitempty" label:"是否启用"`
	Menus  string `json:"menus,omitempty" binding:"omitempty" label:"权限菜单"`
}

func NewRoleService() *RoleService {
	if roleService == nil {
		return &RoleService{
			BaseService: NewBaseApi(&models.Role{}),
		}
	}
	return roleService
}
