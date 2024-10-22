package api

import (
	"fmt"
	"goVueBlog/models"
	"goVueBlog/service"
	"goVueBlog/service/serializer"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"
	"reflect"

	"github.com/gin-gonic/gin"
)

type ArticleApi struct {
	BaseApi
	Service *service.ArticleService
}

func NewArticleApi() ArticleApi {
	return ArticleApi{
		BaseApi: NewBaseApi(),
		Service: service.NewArticleService(),
	}
}

// @Summary 创建文章
// @Tags 文章
// @Param data body PostCreate true "文章"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
// @Router /article [post]
func (p *ArticleApi) CreateArticle(c *gin.Context) {
	var params service.ArticleRequry
	if err := p.BindResquest(BindRequestOtpons{Ctx: c, Ser: &params, BindUri: false}).GetError(); err != nil {
		return
	}
	// 创建一个映射以存储字段名和对应的值
	datas := map[string]any{}

	// 使用反射遍历结构体字段
	v := reflect.ValueOf(params)
	t := reflect.TypeOf(params)

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// 将字段名和对应的值存入 map
		datas[field.Name] = value.Interface()
	}
	if err := p.Service.Create(&datas); err != nil {
		p.Fail(utils.Response{Code: errmsg.ERROR_ARTICLE_ADD_FAIL})
		return
	}
	p.Ok(utils.Response{Code: errmsg.SUCCESS, Data: datas}, "")
}

// PostList
// @Summary 文章列表
// @Tags 文章
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
// @Router /posts [get]
func (m *ArticleApi) ArticleList(c *gin.Context) {
	// 对查询参数进行解析
	var querys serializer.CommonQueryOtpones
	if err := m.ResolveQueryParams(BinldQueryOtpons{Ctx: c, Querys: &querys}).GetError(); err != nil {
		return
	}

	// 将字符串进行切片 "[0,10]"
	var datas []models.Article
	total, error := m.Service.List(&datas, &querys)
	if error != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_ARTICLE_CONTENT})
		return
	}
	// 进行类型断言
	rs := fmt.Sprintf("%d-%d/%d", querys.Ranges.Skip, querys.Ranges.Skip+len(datas), total)
	utils.Success(c, utils.Response{Code: errmsg.SUCCESS, Data: datas}, rs)
}

// ArticleCreate
// @Summary 获取文章
// @Tags 文章
// @Param id path uint true "文章ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object ArticleResponse
// @Router /article/{id} [post]
func (m *ArticleApi) ArticleInfo(c *gin.Context) {
	var id serializer.CommonIDDTO
	if ok := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &id, BindUri: true}).GetError(); ok != nil {
		return
	}
	var datas service.ArticleResponse
	if err := m.Service.GetDataByID(id.ID, &datas); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_ARTICLE_ADD_FAIL})
		return
	}

	m.Ok(utils.Response{Data: datas, Code: errmsg.SUCCESS}, "")
}

// ArticleDelete
// @Summary 删除文章
// @Tags 文章
// @Param id path uint true "文章ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
// @Router /article/{id} [delete]
func (m *ArticleApi) ArticleDelete(c *gin.Context) {
	var id serializer.CommonIDDTO

	if ok := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &id, BindUri: true}).GetError(); ok != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_ARTICLE_DELETE_FAIL, Msg: ok.Error()})
		return
	}
	if err := m.Service.DeleteByID(id.ID); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_ARTICLE_DELETE_FAIL, Msg: err.Error()})
		return
	}

	m.Ok(utils.Response{Code: errmsg.SUCCESS}, "")
}

// ArticleUpdate
// @Summary 更新 文章
// @Tags 文章
// @Param id path uint true "文章ID"
// @Param Authorization header string true "Bearer token"
// @Param body body models.Article true "修改内容"
// @Success 200 {object} object
// @Router /article/{id} [put]
func (m *ArticleApi) ArticleUpdate(c *gin.Context) {
	var id serializer.CommonIDDTO

	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_ARTICLE_UPDATE_FAIL, Msg: err.Error()})
		return
	}
	var params models.Article

	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &params, BindUri: false}).GetError(); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_ARTICLE_UPDATE_FAIL, Msg: err.Error()})
		return
	}

	if err := m.Service.UpdateDataByID(id.ID, &params); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_ARTICLE_UPDATE_FAIL, Msg: err.Error()})
		return
	}

	m.Ok(utils.Response{Code: errmsg.SUCCESS, Data: params}, "")
}
