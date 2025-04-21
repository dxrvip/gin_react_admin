package api

import (
	"goVueBlog/models"
	"goVueBlog/service"
	"goVueBlog/service/serializer"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

var departmentApi *DepartmentApi

type DepartmentApi struct {
	BaseApi
	Service *service.DepartmentService
}

func NewDepartmentApi() *DepartmentApi {
	if departmentApi == nil {
		return &DepartmentApi{
			BaseApi: NewBaseApi(),
			Service: service.NewDepartmentService(),
		}
	}
	return departmentApi
}

// 创建部门
// @Summary 创建部门
// @Tags 部门管理
// @Security ApiKeyAuth
// @Param body body models.Department true "部门信息"
// @Success 200 {object} utils.Response
// @Router /departments [post]
func (a *DepartmentApi) CreateDepartment(c *gin.Context) {
	var department models.Department
	if err := a.BindResquest(c, BindRequestOtpons{Ser: &department, BindUri: false}).GetError(); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	// 获取用户id

	userId, err := utils.GetUserId(c)
	if err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	department.Creator = userId

	if err := a.Service.CreateDepartment(&department); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	a.Ok(c, utils.Response{Data: department, Code: 200}, "")

}

// 更新部门
// @Summary 更新部门
// @Tags 部门管理
// @Security ApiKeyAuth
// @Param id path int true "部门ID"
// @Param body body models.Department true "部门信息"
// @Success 200 {object} utils.Response
// @Router /departments/{id} [put]
func (a *DepartmentApi) UpdateDepartment(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := a.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	var department models.Department
	if err := a.BindResquest(c, BindRequestOtpons{Ser: &department, BindUri: false}).GetError(); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	var updateData = map[string]interface{}{
		"name":      department.Name,
		"parent_id": department.ParentID,
	}
	if err := a.Service.UpdateDepartment(uint(id.ID), updateData); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	a.Ok(c, utils.Response{Data: department, Code: 200}, "")

}

// 删除部门
// @Summary 删除部门
// @Tags 部门管理
// @Security ApiKeyAuth
// @Param id path int true "部门ID"
// @Success 200 {object} utils.Response
// @Router /departments/{id} [delete]
func (a *DepartmentApi) DeleteDepartment(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := a.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	if err := a.Service.DeleteDepartment(uint(id.ID)); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	a.Ok(c, utils.Response{Code: 200}, "")

}

// 获取部门列表
// @Summary 获取部门列表
// @Tags 部门管理
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200 {object} utils.Response
// @Router /department [get]
func (a *DepartmentApi) ListDepartment(c *gin.Context) {
	var querys serializer.CommonQueryOtpones
	if err := a.ResolveQueryParams(c, BinldQueryOtpons{Querys: &querys}).GetError(); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	var departments []models.Department

	rs, err := a.Service.List(&departments, &querys)

	if err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	// a.Ok(utils.Response{Data: departments, Code: http.StatusOK}, rs)
	c.Header("Content-Range", rs)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "okk",
		"data": departments,
	})

}

// 获取部门详情
// @Summary 获取部门详情
// @Tags 部门管理
// @Security ApiKeyAuth
// @Param id path int true "部门ID"
// @Success 200 {object} utils.Response
// @Router /departments/{id} [get]
func (a *DepartmentApi) GetDepartmentById(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := a.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	var department models.Department
	if err := a.Service.GetDataByID(uint(id.ID), &department); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	// a.Ok(utils.Response{Data: department, Code: 200}, "")
	c.JSON(http.StatusOK, map[string]any{
		"code": 200,
		"msg":  "ok",
		"data": department,
	})
}
