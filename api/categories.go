package api

import (
	"goVueBlog/models"
	"goVueBlog/service"
	"goVueBlog/service/serializer"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"

	"github.com/gin-gonic/gin"
)

var shopCategoryApi *CategoriesApi

type CategoriesApi struct {
	BaseApi
	Service *service.ShopCategoryService
}

func NewCategoriesServiceApi() *CategoriesApi {
	if shopCategoryApi == nil {
		return &CategoriesApi{
			BaseApi: NewBaseApi(),
			Service: service.NewCategoriesService(),
		}
	}
	return shopCategoryApi
}

// 列表
func (a *CategoriesApi) List(c *gin.Context) {
	var querys serializer.CommonQueryOtpones
	if err := a.ResolveQueryParams(c, BinldQueryOtpons{Querys: &querys}).GetError(); err != nil {
		return
	}

	// 获取分类列表的逻辑
	var datas []models.ProductCategory

	rs, err := a.Service.List(&datas, &querys)
	if err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	// if data, err := a.Service.ListTree(datas); err == nil {
	// 	a.Ok(c, utils.Response{
	// 		Code: errmsg.SUCCESS,
	// 		Data: data,
	// 	}, rs)
	// } else {
	// 	a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
	// }
	a.Ok(c, utils.Response{
		Code: errmsg.SUCCESS,
		Data: datas,
	}, rs)

}

// 添加
func (a *CategoriesApi) Create(c *gin.Context) {
	var category models.ProductCategory
	if err := a.BindResquest(c, BindRequestOtpons{Ser: &category, BindUri: false}).GetError(); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	if err := a.Service.Create(&category); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	a.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: category}, "")
}

// 更新分类
func (a *CategoriesApi) Update(c *gin.Context) {
	var id serializer.CommonIDDTO

	if err := a.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	var category models.ProductCategory
	if err := a.BindResquest(c, BindRequestOtpons{Ser: &category, BindUri: false}).GetError(); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	if err := a.Service.UpdateDataByID(id.ID, &category); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}

	a.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: category}, "")
}

// 删除
func (a *CategoriesApi) Del(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := a.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	if err := a.Service.DeleteByID(id.ID); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	a.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: map[string]interface{}{"id": id.ID}}, "")
}

// 详细
func (a *CategoriesApi) InfoById(c *gin.Context) {
	var id serializer.CommonIDDTO
	if err := a.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	var data models.ProductCategory
	if err := a.Service.GetDataByID(id.ID, &data); err != nil {
		a.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	a.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: data}, "")
}
