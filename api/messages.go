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

var messageApi *MessageApi

type MessageApi struct {
	BaseApi
	Service *service.MessageService
}

func NewMessageApi() *MessageApi {
	if messageApi == nil {
		return &MessageApi{
			BaseApi: NewBaseApi(),
			Service: service.NewMessageService(),
		}
	}
	return messageApi
}

// 创建消息
// @Summary 创建消息
// @Tags 消息管理
// @Security ApiKeyAuth
// @Param body body models.Message true "消息信息"
// @Success 200 {object} utils.Response
// @Router /message [post]
func (m *MessageApi) CreateMessage(c *gin.Context) {
	var params models.Message
	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &params, BindUri: false}).GetError(); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	idValue, exists := c.Get("id")
	if !exists {
		m.Fail(utils.Response{Code: errmsg.ERROR, Msg: "获取用户ID失败"})
		return
	}
	// 将interface{}转换为uint

	id, _ := idValue.(int)

	params.Creator = uint(id)

	// 创建消息
	data, err := m.Service.CreateMessage(&params)
	if err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	m.Ok(utils.Response{Data: data, Code: http.StatusOK}, "")
}

// 更新消息
// @Summary 更新消息
// @Tags 消息管理
// @Security ApiKeyAuth
// @Param id path int true "留言ID"
// @Param body body models.Message true "消息信息"
// @Success 200 {object} utils.Response
// @Router /messages/{id} [put]
func (m *MessageApi) UpdateMessage(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	var params models.Message
	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &params, BindUri: false}).GetError(); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	if err := m.Service.UpdateDataByID(id.ID, &params); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	m.Ok(utils.Response{Code: errmsg.SUCCESS, Data: ""}, "")
}

// 删除消息
// @Summary 删除消息
// @Tags 消息管理
// @Security ApiKeyAuth
// @Param id path int true "消息ID"
// @Success 200 {object} utils.Response
// @Router /messages/{id} [delete]
func (m *MessageApi) DeleteMessage(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	if err := m.Service.DeleteByID(id.ID); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	m.Ok(utils.Response{Code: errmsg.SUCCESS, Data: ""}, "")
}

// 获取消息列表
// @Summary 获取消息列表
// @Tags 消息管理
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200 {object} utils.Response
// @Router /messages [get]
func (m *MessageApi) ListMessage(c *gin.Context) {
	var querys serializer.CommonQueryOtpones
	if err := m.ResolveQueryParams(BinldQueryOtpons{Ctx: c, Querys: &querys}).GetError(); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	var datas []models.Message
	re, err := m.Service.List(&datas, &querys)
	if err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	m.Ok(utils.Response{Code: errmsg.SUCCESS, Data: datas}, re)
}

// 获取消息详情
// @Summary 获取消息详情
// @Tags 消息管理
// @Security ApiKeyAuth
// @Param id path int true "消息ID"
// @Success 200 {object} utils.Response
// @Router /messages/{id} [get]
func (m *MessageApi) GetMessageById(c *gin.Context) {
	var id serializer.CommonIDDTO
	if error := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &id, BindUri: true}).GetError(); error != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: error.Error()})
		return
	}
	var data models.Message
	if error := m.Service.GetDataByID(id.ID, &data); error != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: error.Error()})
		return
	}
	m.Ok(utils.Response{Code: errmsg.SUCCESS, Data: data}, "")

}
