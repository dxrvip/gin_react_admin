package api

import (
	"goVueBlog/models"
	"goVueBlog/service"
	"goVueBlog/service/serializer"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"

	"github.com/gin-gonic/gin"
)

var productApi *ProductApi

type ProductApi struct {
	BaseApi
	Service *service.ProductService
}

func NewProductApi() *ProductApi {
	if productApi == nil {
		return &ProductApi{
			BaseApi: NewBaseApi(),
			Service: service.NewProductService(),
		}
	}
	return productApi
}

func (pa *ProductApi) ProductList(c *gin.Context) {
	var querys serializer.CommonQueryOtpones
	if err := pa.ResolveQueryParams(c, BinldQueryOtpons{Querys: &querys}).GetError(); err != nil {
		pa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	var datas []models.Product

	rs, err := pa.Service.List(&datas, &querys)
	if err != nil {
		pa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	// 转换响应结构
	var responseList []service.ProductResponse
	for _, p := range datas {
		responseList = append(responseList, service.ProductResponse{
			ID:          p.ID,
			Title:       p.Title,
			Price:       p.Price,
			Stock:       p.Stock,
			Description: p.Description,
			// Attributes:        p.AttributesParsed,
			Images:            p.Images,
			Status:            p.Status,
			BrandID:           p.BrandID,
			ProductCategoryID: p.ProductCategoryID,
			CreatedAt:         p.CreatedAt,
			SecondHandSku:     p.SecondHandSku,
		})
	}

	pa.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: responseList}, rs)
}

func (pa *ProductApi) ProductCreate(c *gin.Context) {
	var params service.ProductRequry
	if err := pa.BindResquest(c, BindRequestOtpons{Ser: &params, BindUri: false}).GetError(); err != nil {
		pa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	var savaData models.Product = models.Product{
		Title:             params.Title,
		Price:             params.Price,
		Stock:             params.Stock,
		Description:       params.Description,
		Images:            params.Images,
		ProductCategoryID: params.ProductCategoryID,
		Status:            params.Status,
		BrandID:           params.BrandID,
		Attributes:        params.Attributes,
	}
	if err := pa.Service.Create(&savaData); err != nil {
		pa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	pa.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: service.ProductResponse{
		ID:                savaData.ID,
		Title:             savaData.Title,
		Price:             savaData.Price,
		Stock:             savaData.Stock,
		Description:       savaData.Description,
		Images:            savaData.Images,
		ProductCategoryID: savaData.ProductCategoryID,
		Status:            savaData.Status,
		BrandID:           savaData.BrandID,
		Attributes:        params.Attributes,
	}}, "")
}

func (pa *ProductApi) ProductUpdate(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := pa.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		pa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	var params service.ProductRequry
	if err := pa.BindResquest(c, BindRequestOtpons{Ser: &params, BindUri: false}).GetError(); err != nil {
		pa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	// 手动序列化attributes
	// attrsJSON, err := json.Marshal(params.Attributes)
	// if err != nil {
	// 	pa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: "属性数据格式错误"})
	// 	return
	// }
	var savaData models.Product = models.Product{

		Title:             params.Title,
		Price:             params.Price,
		Stock:             params.Stock,
		Description:       params.Description,
		Images:            params.Images,
		ProductCategoryID: params.ProductCategoryID,
		Status:            params.Status,
		BrandID:           params.BrandID,
		// Attributes:        attrsJSON,
	}
	if updata, err := pa.Service.Updates(id.ID, savaData); err != nil {
		pa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	} else {
		pa.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: updata}, "")
	}

}

func (pa *ProductApi) ProductDelete(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := pa.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		pa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	if err := pa.Service.DeleteByID(id.ID); err != nil {
		pa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	pa.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: map[string]interface{}{
		"id": id.ID,
	}}, "")
}

func (pa *ProductApi) ProductInfo(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := pa.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		pa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	var data models.Product
	if err := pa.Service.GetDataByID(id.ID, &data); err != nil {
		pa.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	// 反序列化
	res := service.ProductResponse{
		ID:                data.ID,
		Title:             data.Title,
		Price:             data.Price,
		Stock:             data.Stock,
		Description:       data.Description,
		Images:            data.Images,
		ProductCategoryID: data.ProductCategoryID,
		Status:            data.Status,
		BrandID:           data.BrandID,
		Attributes:        data.Attributes,
		CreatedAt:         data.CreatedAt,
	}
	pa.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: res}, "")
}
