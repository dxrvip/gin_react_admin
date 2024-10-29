/* 文章管理 */
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
// @Param data body service.ArticleRequry true "文章"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
// @Router /article [post]
// @Description: 创建一片文章
func (p *ArticleApi) CreateArticle(c *gin.Context) {
	var params service.ArticleRequry
	if err := p.BindResquest(BindRequestOtpons{Ctx: c, Ser: &params, BindUri: false}).GetError(); err != nil {
		return
	}
	params.CreatedAt = time.Now()
	result, err := p.Service.Create(params)
	if err != nil {
		p.Fail(utils.Response{Code: errmsg.ERROR_ARTICLE_ADD_FAIL})
		return
	}
	p.Ok(utils.Response{Code: errmsg.SUCCESS, Data: result}, "")
}

// ArticleList
// @Summary 文章列表
// @Tags 文章
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
// @Router /article [get]
func (m *ArticleApi) ArticleList(c *gin.Context) {
	// 对查询参数进行解析
	var querys serializer.CommonQueryOtpones
	if err := m.ResolveQueryParams(BinldQueryOtpons{Ctx: c, Querys: &querys}).GetError(); err != nil {
		return
	}

	// 将字符串进行切片 "[0,10]"
	var datas []models.Article
	rs, error := m.Service.List(&datas, &querys)
	if error != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_ARTICLE_CONTENT})
		return
	}

	m.Ok(utils.Response{Code: errmsg.SUCCESS, Data: datas}, rs)
}

// ArticleCreate
// @Summary 获取文章
// @Tags 文章
// @Param id path uint true "文章ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
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
