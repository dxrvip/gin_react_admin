package api

import (
	"goVueBlog/models"
	"goVueBlog/service"
	"goVueBlog/service/serializer"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"

	"github.com/gin-gonic/gin"
)

var secondHandSkuApi *SecondHandSkuApi

type SecondHandSkuApi struct {
	BaseApi
	Service *service.SecondHandSkuService
}

func NewSecondHandSkuApi() *SecondHandSkuApi {
	if secondHandSkuApi == nil {
		return &SecondHandSkuApi{
			BaseApi: NewBaseApi(),
			Service: service.NewSecondHandSkuService(),
		}
	}
	return secondHandSkuApi
}

// CreateSecondHandSku 创建二手商品SKU
// @Summary 创建二手商品SKU
// @Tags 二手商品
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body service.CreateSecondHandSkuRequest true "创建二手商品SKU请求参数"
// @Success 200 {object} utils.Response
// @Router /api/v1/secondHandSkus [post]
func (sa *SecondHandSkuApi) CreateSecondHandSku(c *gin.Context) {
	var req service.CreateSecondHandSkuRequest
	if err := sa.BindResquest(c, BindRequestOtpons{Ser: &req, BindUri: false}).GetError(); err != nil {
		sa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	sku, err := sa.Service.CreateSecondHandSku(&req)
	if err != nil {
		sa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	sa.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: sku}, "")
}

// UpdateSecondHandSku 更新二手商品SKU
// @Summary 更新二手商品SKU
// @Tags 二手商品
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "SKU ID"
// @Param object body service.UpdateSecondHandSkuRequest true "更新二手商品SKU请求参数"
// @Success 200 {object} utils.Response
// @Router /api/v1/secondHandSkus/{id} [put]
func (sa *SecondHandSkuApi) UpdateSecondHandSku(c *gin.Context) {
	var id serializer.CommonIDDTO
	var req service.UpdateSecondHandSkuRequest

	if err := sa.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		sa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	if err := sa.BindResquest(c, BindRequestOtpons{Ser: &req, BindUri: false}).GetError(); err != nil {
		sa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	if err := sa.Service.UpdateSecondHandSku(id.ID, &req); err != nil {
		sa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	sa.Ok(c, utils.Response{Code: errmsg.SUCCESS}, "")
}

// GetSecondHandSku 获取二手商品SKU详情
// @Summary 获取二手商品SKU详情
// @Tags 二手商品
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "SKU ID"
// @Success 200 {object} utils.Response
// @Router /api/v1/secondHandSkus/{id} [get]
func (sa *SecondHandSkuApi) GetSecondHandSku(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := sa.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		sa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	sku, err := sa.Service.GetSecondHandSkuByID(id.ID)
	if err != nil {
		sa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	sa.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: sku}, "")
}

// GetProductSecondHandSkus 获取商品的二手SKU列表
// @Summary 获取商品的二手SKU列表
// @Tags 二手商品
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param productId query int true "商品ID"
// @Success 200 {object} utils.Response
// @Router /api/v1/products/{productId}/secondHandSkus [get]
func (sa *SecondHandSkuApi) GetProductSecondHandSkus(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := sa.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		sa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	skus, err := sa.Service.GetProductSecondHandSkus(id.ID)
	if err != nil {
		sa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	sa.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: skus}, "")
}

// DeleteSecondHandSku 删除二手商品SKU
// @Summary 删除二手商品SKU
// @Tags 二手商品
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path int true "SKU ID"
// @Success 200 {object} utils.Response
// @Router /api/v1/secondHandSkus/{id} [delete]
func (sa *SecondHandSkuApi) DeleteSecondHandSku(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := sa.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		sa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	if err := sa.Service.DeleteSecondHandSku(id.ID); err != nil {
		sa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	sa.Ok(c, utils.Response{Code: errmsg.SUCCESS}, "")
}

// ListSecondHandSkus 获取二手商品SKU列表
// @Summary 获取二手商品SKU列表
// @Tags 二手商品
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200 {object} utils.Response
// @Router /api/v1/secondHandSkus [get]
func (sa *SecondHandSkuApi) ListSecondHandSkus(c *gin.Context) {
	var querys serializer.CommonQueryOtpones
	if err := sa.ResolveQueryParams(c, BinldQueryOtpons{Querys: &querys}).GetError(); err != nil {
		sa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	var skus []models.SecondHandSku
	rs, err := sa.Service.List(&skus, &querys)
	if err != nil {
		sa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	sa.Ok(c, utils.Response{
		Code: errmsg.SUCCESS,
		Data: skus,
	}, rs)
}
