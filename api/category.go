package api

import (
	"encoding/json"
	"fmt"
	"goVueBlog/models"
	"goVueBlog/service"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryApi struct {
	BaseApi
	Service *service.CategoryService
}

type CategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func NewCategoryApi() *CategoryApi {
	return &CategoryApi{
		BaseApi: NewBaseApi(),
		Service: service.NewCateGoryService(),
	}
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
	var (
		req  CategoryRequest
		code int
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		code = errmsg.ERROR_CATENAME_FORMAT
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error(), "code": code})
		return
	}
	var category models.Category
	category.Name = req.Name
	if err := m.Service.CreateCategory(&category); err != nil {
		code = errmsg.REEOR_CATE_ADD_FAIL
		c.JSON(http.StatusOK, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}

	// 添加成功返回信息
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{"msg": errmsg.GetErrMsg(code), "code": code, "id": category.ID})

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

	// 获取查询参数
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
	// 获取分类列表的逻辑
	var (
		categories []models.Category
		code       int
	)

	categoriesPointer, totalCount, err := m.Service.GetCategoryList(skip, limit, sortArr)
	if err != nil {
		code = errmsg.ERROR_ARTICLE_CONTENT
		c.JSON(http.StatusOK, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}
	categories = *categoriesPointer
	rs := fmt.Sprintf("%d-%d/%d", skip, skip+len(categories), totalCount)
	utils.Success(c, utils.Response{
		Code: code,
		Data: categories,
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
	var (
		code     int
		jsonData models.Category
	)
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		code = errmsg.ERROR_CATENAME_FORMAT
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error(), "code": code})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code = errmsg.ERROR_CATEGORY_BY_ID_NOT_EXIST
		c.JSON(http.StatusBadRequest, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}
	// 修改
	if err := m.Service.UpdateCategoryByID(uint(id), &jsonData); err != nil {
		code = errmsg.ERROR_CATEGORY_UPDATE_FAIL
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}
	c.JSON(http.StatusOK, jsonData)
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
	id := c.Param("id")
	var (
		code int
	)
	// 将字符串id转成uint
	idInt, err := strconv.Atoi(id)
	if err != nil {
		code = errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}
	category, err := m.Service.GetCategoryByID(uint(idInt))
	if err != nil {
		code = errmsg.ERROR_CATENAME_EXITS
		c.JSON(http.StatusOK, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}
	c.JSON(http.StatusOK, category)
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
	var (
		code int
	)
	id, err := strconv.Atoi(c.Param("id"))
	// 根据id删除分类的逻辑
	if err != nil {
		code = errmsg.ERROR_CATENAME_EXITS
		c.JSON(http.StatusBadRequest, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}
	if !m.Service.IsCategoryExistByID(uint(id)) {
		code = errmsg.ERROR_CATENAME_EXITS
		c.JSON(http.StatusBadRequest, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return

	}
	if err := m.Service.DeleteCategoryByID(uint(id)); err != nil {
		code = errmsg.ERROR
		c.JSON(http.StatusBadRequest, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return

	}
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})

	// 删除分类的逻辑
}
