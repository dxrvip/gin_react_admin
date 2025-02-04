/* 权限管理 */
package api

import (
	"goVueBlog/service"
	"goVueBlog/service/serializer"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"

	"github.com/gin-gonic/gin"
)

var roleApi *RoleApi

type RoleApi struct {
	BaseApi
	Service *service.RoleService
}

func NewRoleApi() *RoleApi {
	if roleApi == nil {
		return &RoleApi{
			BaseApi: NewBaseApi(),
			Service: service.NewRoleService(),
		}
	}
	return roleApi
}

// AddCategory
// @Summary 添加权限
// @Tags 权限管理
// @Accept json
// @Param Authorization header string true "Bearer token"
// @Param data body CategoryRequest true "权限名称"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /role [post]
func (m *RoleApi) CreateRole(c *gin.Context) {
	var params service.RoleParams
	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &params, BindUri: false}).GetError(); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	// 假如数据库
	datas, err := m.Service.Create(params)
	if err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	m.Ok(utils.Response{Code: errmsg.SUCCESS, Data: &datas}, "")
}

func (m *RoleApi) GetRoleById(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(utils.Response{Msg: err.Error()})
		return
	}

	datas, err := m.Service.GetDataByID(uint(id.ID))
	if err != nil {
		m.Fail(utils.Response{Msg: err.Error()})
		return
	}
	m.Ok(utils.Response{Data: datas}, "")

}
func (m *RoleApi) ListRole(c *gin.Context) {
	var querys serializer.CommonQueryOtpones
	if err := m.ResolveQueryParams(BinldQueryOtpons{Ctx: c, Querys: &querys}).GetError(); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	var datas []service.RoleResponse
	rs, err := m.Service.List(&datas, &querys)
	if err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	m.Ok(utils.Response{Code: errmsg.SUCCESS, Data: datas}, rs)
}

func (m *RoleApi) UpdateRole(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(utils.Response{Msg: err.Error()})
		return
	}

	// 获取更新数据
	var params service.UpdateParams
	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &params, BindUri: false}).GetError(); err != nil {
		m.Fail(utils.Response{Msg: err.Error()})
		return
	}
	// 组织更新内容
	activeUint := 0
	if params.Active {
		activeUint = 1
	}
	// 判断是否有用户数据 更新用户数据
	updateData := make(map[string]interface{})

	updateData["active"] = activeUint
	updateData["key"] = params.Key
	updateData["name"] = params.Name
	updateData["menus"] = params.Menus
	updateData["sort"] = params.Sort
	// 先更新数据

	if err := m.Service.UpdateRoleDataByID(id.ID, &updateData); err != nil {
		m.Fail(utils.Response{Msg: "数据更新失败！"})
		return
	}

	updateData["id"] = id.ID

	m.Ok(utils.Response{Data: updateData}, "")
}

// 更新用户权限
func (m *RoleApi) UpdateRoleUsers(c *gin.Context) {
	// 获取更新数据
	var id serializer.CommonIDDTO
	// 获取id
	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(utils.Response{Msg: err.Error()})
		return
	}

	var params struct {
		User []uint `json:"user" binding:"required"`
	}

	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &params, BindUri: false}).GetError(); err != nil {
		m.Fail(utils.Response{Msg: err.Error()})
		return
	}

	if err := m.Service.UpdateRoleUsers(id.ID, params.User); err != nil {
		m.Fail(utils.Response{Msg: "数据更新失败！"})
		return
	}
	m.Ok(utils.Response{Data: map[string]interface{}{"id": id.ID}}, "")

}

// 更新用户权限
func (m *RoleApi) UpdateRoleMenus(c *gin.Context) {
	// 获取更新数据
	var id serializer.CommonIDDTO
	// 获取id
	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(utils.Response{Msg: err.Error()})
		return
	}

	var params service.UpdateParams
	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &params, BindUri: false}).GetError(); err != nil {
		m.Fail(utils.Response{Msg: err.Error()})
		return
	}

	if err := m.Service.UpdateRoleMenus(id.ID, &params); err != nil {
		m.Fail(utils.Response{Msg: "数据更新失败！"})
		return
	}
	m.Ok(utils.Response{Data: map[string]interface{}{"id": id.ID}}, "")

}

func (m *RoleApi) DelRole(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(utils.Response{Msg: err.Error()})
		return
	}

	if err := m.Service.DeleteByID(id.ID); err != nil {
		m.Fail(utils.Response{Msg: err.Error()})
		return
	}

	m.Ok(utils.Response{Data: id.ID}, "")
}
