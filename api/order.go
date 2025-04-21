package api

import (
	"goVueBlog/models"
	"goVueBlog/service"
	"goVueBlog/service/serializer"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"
	"time"

	"github.com/gin-gonic/gin"
)

var orderApi *OrderApi

type OrderApi struct {
	BaseApi
	Service *service.OrderService
}

func NewOrderApi() *OrderApi {
	if orderApi == nil {
		return &OrderApi{
			BaseApi: NewBaseApi(),
			Service: service.NewOrderService(),
		}
	}
	return orderApi
}

// CreateOrder 创建订单
func (oa *OrderApi) CreateOrder(c *gin.Context) {
	var req service.CreateOrderRequest
	if err := oa.BindResquest(c, BindRequestOtpons{Ser: &req, BindUri: false}).GetError(); err != nil {
		oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	// 判断用户是否存在

	// userID, err := utils.GetUserId(c)
	// if err != nil {
	// 	oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
	// 	return
	// }

	response, err := oa.Service.CreateOrder(&req)
	if err != nil {
		oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	oa.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: response}, "")
}

// update 更新
func (oa *OrderApi) UpdateOrder(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := oa.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	var req service.UpdateOrderRequest
	if err := oa.BindResquest(c, BindRequestOtpons{Ser: &req, BindUri: false}).GetError(); err != nil {
		oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	//组织更新数据
	var updateData models.Order = models.Order{
		BaseModel: models.BaseModel{
			ID:        id.ID,
			CreatedAt: time.Time(req.ClienTime),
		},
		Status:     req.Status,
		Address:    req.Address,
		Note:       req.Note,
		CostPrice:  req.CostPrice,
		OrderItems: req.Items, // 注意：这里假设 OrderItems 是一个切片
		Weight:     req.Weight,
		UserID:     req.UserId, // 确保 UserID 也被更新
	}
	if response, err := oa.Service.Updates(id.ID, &updateData); err != nil {
		oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	} else {

		oa.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: response}, "")
	}
}

// UpdateOrderStatus 更新订单状态
func (oa *OrderApi) UpdateOrderStatus(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := oa.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	status := c.PostForm("status")
	if status == "" {
		oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: "状态不能为空"})
		return
	}

	if err := oa.Service.UpdateOrderStatus(id.ID, status); err != nil {
		oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	oa.Ok(c, utils.Response{Code: errmsg.SUCCESS}, "")
}

// GetOrderDetail 获取订单详情
func (oa *OrderApi) GetOrderDetail(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := oa.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	response, err := oa.Service.GetOrderByID(id.ID)
	if err != nil {
		oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	oa.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: response}, "")
}

// GetUserOrders 获取用户订单列表
func (oa *OrderApi) GetUserOrders(c *gin.Context) {
	// 从上下文获取用户ID
	var querys serializer.CommonQueryOtpones
	if err := oa.ResolveQueryParams(c, BinldQueryOtpons{Querys: &querys}).GetError(); err != nil {
		oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	var orders []models.Order
	rs, err := oa.Service.List(&orders, &querys)
	if err != nil {
		oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	oa.Ok(c, utils.Response{
		Code: errmsg.SUCCESS,
		Data: orders,
	}, rs)
}

// 删除订单
func (oa *OrderApi) DelteOrder(c *gin.Context) {
	var id serializer.CommonIDDTO

	if err := oa.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	if err := oa.Service.DeleteByID(id.ID); err != nil {
		oa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	oa.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: map[string]interface{}{"id": id.ID}}, "")
}
