package api

import (
	"goVueBlog/models"
	"goVueBlog/service"
	"goVueBlog/service/serializer"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type BrandApi struct {
	BaseApi
	Service *service.BrandService
}

func NewBrandApi() BrandApi {
	return BrandApi{
		BaseApi: NewBaseApi(),
		Service: service.NewBrandService(),
	}

}

func (m *BrandApi) Create(c *gin.Context) {
	var params service.BrandCreateData
	if err := m.BindResquest(c, BindRequestOtpons{Ser: &params, BindUri: false}).GetError(); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	var savaData models.Brand = models.Brand{
		Name:        params.Name,
		Logo:        params.Logo,
		Description: params.Description,
	}
	// 将数据保存到数据库
	err := m.Service.Create(&savaData)
	if err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	} else {

		m.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: savaData}, "")
		return
	}

}

func (m *BrandApi) List(c *gin.Context) {
	var querys serializer.CommonQueryOtpones
	if err := m.ResolveQueryParams(c, BinldQueryOtpons{Querys: &querys}).GetError(); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	var datas []service.BrandResponse
	rs, err := m.Service.List(&datas, &querys)
	if err != nil {
		m.Fail(c, utils.Response{Msg: err.Error(), Code: errmsg.ERROR})
		return
	}
	m.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: datas}, rs)
}

func (m *BrandApi) InfoById(c *gin.Context) {
	var id serializer.CommonIDDTO
	// 获取id
	if err := m.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})

		return
	}
	// 根据id获取数据
	var brand models.Brand
	if err := m.Service.GetDataByID(id.ID, &brand); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	m.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: brand}, "")
}

func (m *BrandApi) Update(c *gin.Context) {
	var id serializer.CommonIDDTO

	if err := m.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	var upData service.BrandCreateData
	if err := m.BindResquest(c, BindRequestOtpons{Ser: &upData, BindUri: false}).GetError(); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	// 更新数据
	_, err := m.Service.Updates(id.ID, upData)
	if err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	//更新成功
	m.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: models.Brand{
		Name:        upData.Name,
		BaseModel:   models.BaseModel{ID: id.ID},
		Logo:        upData.Logo,
		Description: upData.Description,
	}}, "")
}

// 删除
func (m *BrandApi) Delete(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := m.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return

	}
	if err := m.Service.DeleteByID(id.ID); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return

	}

	m.Ok(c, utils.Response{Code: errmsg.SUCCESS}, "")
}
