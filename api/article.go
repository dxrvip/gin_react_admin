package api

import (
	"encoding/json"
	"fmt"
	"goVueBlog/models"
	"goVueBlog/service"
	"goVueBlog/service/serializer"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleApi struct {
	BaseApi
	Service *service.ArticleService
	Code    int
	Model   []models.Article
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
	var body service.ArticleRequry
	if err := p.BindResquest(BindRequestOtpons{Ctx: c, Ser: body, BindUri: false}); err != nil {
		return
	}
	// 添加到数据库
	params := map[string]any{
		"title":   body.Title,
		"content": body.Content,
		"cid":     body.Cid,
		"picture": body.Img,
	}
	if err := p.Service.Create(&params); err != nil {
		p.Code = errmsg.ERROR_ARTICLE_ADD_FAIL
		p.Fail(utils.Response{Code: p.Code})
		return
	}
	p.Ok(utils.Response{Msg: "success", Data: params})
}

// PostList
// @Summary 文章列表
// @Tags 文章
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
// @Router /posts [get]
func (p *ArticleApi) ArticleList(c *gin.Context) {
	// 对查询参数进行解析
	var querys serializer.CommonQueryOtpones
	p.ResolveQueryParams(BinldQueryOtpons{Ctx: c, Querys: &querys})
	filter := c.DefaultQuery("filter", "{}")
	ranges := c.DefaultQuery("range", "[0,10]")
	var skip, limit int
	var rangeArr []int
	if err := json.Unmarshal([]byte(ranges), &rangeArr); err != nil {
		skip, limit = 0, 10
	} else {
		skip, limit = rangeArr[0], rangeArr[1]-rangeArr[0]+1
	}

	fmt.Println(skip, limit)
	var sortArr []string
	sort := c.DefaultQuery("sort", "['id', 'ASC']")
	if err := json.Unmarshal([]byte(sort), &sortArr); err != nil {
		sortArr = make([]string, 2)
		sortArr[0] = "id"
		sortArr[1] = "ASC"
	}

	fmt.Println(filter, ranges, sort)
	// 将字符串进行切片 "[0,10]"

	datas, total, error := p.Service.List(models.Article{}, p.Model, limit, skip, sortArr)
	if error != nil {
		p.Code = errmsg.ERROR_ARTICLE_CONTENT
		p.Fail(utils.Response{Code: p.Code})
		return
	}
	// 进行类型断言
	rs := fmt.Sprintf("%d-%d/%d", skip, skip+len(datas.([]models.Article)), total)
	utils.Success(c, utils.Response{Code: errmsg.SUCCESS, Data: datas}, rs)
}

// PostCreate
// @Summary 获取文章
// @Tags 文章
// @Param id path uint true "文章ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object PostResponse
// @Router /posts/{id} [post]
func (p *ArticleApi) ArticleInfo(c *gin.Context) {
	var id serializer.CommonIDDTO
	if ok := p.BindResquest(BindRequestOtpons{Ctx: c, Ser: &id, BindUri: true}).GetError(); ok != nil {
		return
	}

	var article *models.Article
	article, err := p.Service.GetArticleById(int(id.ID))
	var code int
	if err != nil {
		code = errmsg.ERROR_ART_NOT_EXIST
		utils.Fails(c, utils.Response{Code: code})
		return
	}
	code = errmsg.SUCCESS
	data := service.ArticleResponse{
		ID:      article.ID,
		Content: article.Content,
		Title:   article.Title,
		Cid:     article.Cid,
		Desc:    article.Desc,
		Picture: article.Picture,
	}
	utils.Success(c, utils.Response{Data: data, Code: code}, "")
}

// PostDelete
// @Summary 删除文章
// @Tags 文章
// @Param id path uint true "文章ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
// @Router /posts/{id} [delete]
func (p *ArticleApi) ArticleDelete(c *gin.Context) {
	var (
		code int
		id   serializer.CommonIDDTO
	)
	if ok := p.BindResquest(BindRequestOtpons{Ctx: c, Ser: &id, BindUri: true}).GetError(); ok != nil {
		return
	}
	if err := p.Service.DeleteArticleByID(int(id.ID)); err != nil {
		code = errmsg.ERROR_ARTICLE_DELETE_FAIL
		c.JSON(http.StatusBadRequest, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}

	utils.Success(c, utils.Response{Code: errmsg.SUCCESS}, "")
}

// PostUpdate
// @Summary 更新 文章
// @Tags 文章
// @Param id path uint true "文章ID"
// @Param Authorization header string true "Bearer token"
// @Param body body models.Article true "修改内容"
// @Success 200 {object} object
// @Router /posts/{id} [put]
func (p *ArticleApi) ArticleUpdate(c *gin.Context) {
	var (
		code int
		id   serializer.CommonIDDTO
	)

	if err := p.BindResquest(BindRequestOtpons{Ctx: c, Ser: &id, BindUri: true}).GetError(); err != nil {
		return
	}
	body := models.Article{}

	if err := p.BindResquest(BindRequestOtpons{Ctx: c, Ser: &body, BindUri: false}); err != nil {
		return
	}

	if err := p.Service.UpdateArticleById(int(id.ID), &body); err != nil {
		code = errmsg.ERROR_ARTICLE_UPDATE_FAIL
		c.JSON(http.StatusBadRequest, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}

	p.Ok(utils.Response{Code: errmsg.SUCCESS, Data: body})
}
