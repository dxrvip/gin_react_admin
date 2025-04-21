package api

import (
	"goVueBlog/models"
	"goVueBlog/service"
	"goVueBlog/service/serializer"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"
	"goVueBlog/validators"

	"github.com/gin-gonic/gin"
)

// 属性api
var attributekeyApi *AttributeApi

// 修正 AttributeKeyApi 使用正确的模型
type AttributeApi struct {
	BaseApi
	Service *service.AttributeService
}

func NewAttributeApi() *AttributeApi {
	if attributekeyApi == nil {
		return &AttributeApi{
			BaseApi: NewBaseApi(),
			Service: service.NewAttributeKeyService(),
		}
	}
	return attributekeyApi
}

// 属性列表
func (a *AttributeApi) ListAttribute(c *gin.Context) {
	var querys serializer.CommonQueryOtpones
	if err := a.ResolveQueryParams(c, BinldQueryOtpons{Querys: &querys}).GetError(); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	var datas []models.Attribute // 修正为正确模型

	rs, err := a.Service.List(&datas, &querys)

	// 映射到 DTO
	response := make([]service.ListAttribute, len(datas))
	for i, attr := range datas {
		response[i] = service.ListAttribute{
			ID: attr.ID,
			BaseAttribute: service.BaseAttribute{
				Name:         attr.Name,
				Type:         attr.Type,
				IsRequired:   attr.IsRequired,
				DefaultValue: attr.DefaultValue,
				Options:      attr.Options,
				MinValue:     attr.MinValue,
				MaxValue:     attr.MaxValue,
				Unit:         attr.Unit,
			},
			Categories: attr.Categories,
			// CategoryIDs: attr.CategoryIDs,
		}
	}
	if err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	a.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: response}, rs)

}

// 添加属性key
func (a *AttributeApi) CreateAttribute(c *gin.Context) {
	var req service.Request
	err := a.BindResquest(c, BindRequestOtpons{Ser: &req, BindUri: false}).GetError()
	if err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	// 数据验证
	if req.Type == models.TypeEnum {
		if len(req.Options) == 0 {
			a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: "枚举类型必须包含选项"})
			return
		}
		if validators.ContainsDuplicates(req.Options) {
			a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: "枚举选项不能重复"})
			return
		}
	}
	// 转换数据
	attribute := models.Attribute{
		Name:         req.Name,
		Type:         req.Type,
		IsRequired:   req.IsRequired,
		DefaultValue: req.DefaultValue,
		MinValue:     req.MinValue,
		MaxValue:     req.MaxValue,
		Options:      req.Options, // 自动触发 BeforeSave 钩子
		CategoryIDs:  req.CategoryIDs,
	}
	if err := a.Service.Create(&attribute); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	a.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: attribute}, "")

}

// 修改属性key
func (a *AttributeApi) UpdateAttribute(c *gin.Context) {
	id, err := a.GetParamsId(c)
	if err != nil {
		return
	}
	var req service.Request // 修正为正确模型
	if err := a.BindResquest(c, BindRequestOtpons{Ser: &req, BindUri: false}).GetError(); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	params := models.Attribute{
		Name:         req.Name,
		Type:         req.Type,
		IsRequired:   req.IsRequired,
		DefaultValue: req.DefaultValue,
		MinValue:     req.MinValue,
		MaxValue:     req.MaxValue,
		Options:      req.Options, // 自动触发 BeforeSave 钩子
		CategoryIDs:  req.CategoryIDs,
		Unit:         req.Unit,
	}
	if err := a.Service.Updates(id, &params); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	} else {

		a.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: params}, "")
	}
}

// 删除属性key
func (a *AttributeApi) DelAttribute(c *gin.Context) {
	id, err := a.GetParamsId(c)
	if err != nil {
		return
	}
	if err := a.Service.DeleteByID(id); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	a.Ok(c, utils.Response{Data: map[string]interface{}{"id": id}}, "")
}

// 查看属性详细Key
func (a *AttributeApi) InfoAttribute(c *gin.Context) {
	id, err := a.GetParamsId(c)
	if err != nil {
		return
	}
	var attribute service.InfoAttribute
	if err := a.Service.GetAttributeDetails(id, &attribute); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	a.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: attribute}, "")
}
