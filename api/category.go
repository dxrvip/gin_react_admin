/* 分类管理 */
package api

import (
	"goVueBlog/models"
	"goVueBlog/service"
	"goVueBlog/service/serializer"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type CategoryApi struct {
	BaseApi
	Service *service.CategoryService
	Code    int
}

func NewCategoryApi() CategoryApi {
	return CategoryApi{
		BaseApi: NewBaseApi(),
		Service: service.NewCateGoryService(),
	}
}

type CategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

// AddCategory
// @Summary 添加分类
// @Tags 分类管理
// @Accept json
// @Param Authorization header string true "Bearer token"
// @Param data body CategoryRequest true "分类名称"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /category [post]
func (m CategoryApi) AddCategory(c *gin.Context) {
	// 获取请求参数
	var req CategoryRequest
	if err := m.BindResquest(c, BindRequestOtpons{Ser: &req, BindUri: false}).GetError(); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_CATENAME_FORMAT})
		return
	}
	var saveData models.Category = models.Category{
		Name: req.Name,
	}
	err := m.Service.Create(&saveData)
	if err != nil {
		m.Fail(c, utils.Response{Code: errmsg.REEOR_CATE_ADD_FAIL})
		return
	}

	// 添加成功返回信息
	m.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: saveData}, "")

}

// GetCategoryList
// @Summary 添加列表
// @Tags 分类管理
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Success 200 {string} json [{}]
// @Router /category [get]
func (m *CategoryApi) GetCategoryList(c *gin.Context) {
	// 获取分类列表
	var querys serializer.CommonQueryOtpones
	if err := m.ResolveQueryParams(c, BinldQueryOtpons{Querys: &querys}).GetError(); err != nil {
		return
	}

	// 获取分类列表的逻辑
	var datas []models.Category

	rs, err := m.Service.List(&datas, &querys)
	if err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_ARTICLE_CONTENT})
		return
	}

	m.Ok(c, utils.Response{
		Code: m.Code,
		Data: datas,
	}, rs)
}

// UpdateCategory
// @Summary 修改
// @Tags 分类管理
// @Param id path int true "分类id"
// @Param Authorization header string true "Bearer token"
// @Param data body CategoryRequest true "分类名称"
// @Accept json
// @Success 200 {string} models.Category
// @Router /category/{id} [put]
func (m *CategoryApi) UpdateCategory(c *gin.Context) {
	// 修改分类
	var id serializer.CommonIDDTO
	if err := m.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_CATEGORY_BY_ID_NOT_EXIST})
		return
	}
	var jsonData models.Category
	if err := m.BindResquest(c, BindRequestOtpons{Ser: &jsonData, BindUri: false}).GetError(); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_CATENAME_FORMAT})
		return
	}

	// 修改
	if err := m.Service.UpdateDataByID(uint(id.ID), &jsonData); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_CATEGORY_UPDATE_FAIL})
		return
	}
	m.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: jsonData}, "")
}

// GetCategoryById
// @Summary 分类详情
// @Tags 分类管理
// @Param id path int true "分类id"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Success 200 {string} models.Category
// @Router /category/{id} [get]
func (m *CategoryApi) GetCategoryById(c *gin.Context) {
	// 根据分类id获取分类详情的逻辑
	var id serializer.CommonIDDTO

	if err := m.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR})
		return
	}
	var datas models.Category
	if err := m.Service.GetDataByID(id.ID, &datas); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_CATENAME_EXITS})
		return
	}
	m.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: datas}, "")
}

// DeleteCategory
// @Summary 删除分类
// @Tags 分类管理
// @Param id path int true "分类id"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Success 200 {string} json [{}]
// @Router /category/{id} [delete]
func (m *CategoryApi) DeleteCategory(c *gin.Context) {

	var id serializer.CommonIDDTO
	// 根据id删除分类的逻辑
	if err := m.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_CATENAME_EXITS, Msg: err.Error()})
		return
	}

	if err := m.Service.DeleteByID(uint(id.ID)); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return

	}
	m.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: nil}, "")
}
