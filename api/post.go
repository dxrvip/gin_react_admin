package api

import (
	"encoding/json"
	"fmt"
	"goVueBlog/models"
	h "goVueBlog/modules"
	"goVueBlog/serializer"
	"goVueBlog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ResponsePosts 响应文章列表
type ResponsePosts []struct {
	models.Article
}

// PostCreate
// @Summary 创建文章
// @Tags 文章
// @Param data body PostCreate true "文章"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
// @Router /posts [post]
func CreatePost(c *gin.Context) {
	body := serializer.PostRequry{}
	var code int
	if err := c.ShouldBindJSON(&body); err != nil {
		code = errmsg.ERROR_ARTICLE_ERROR
		h.Fails(c, h.Response{Code: code})
	}
	fmt.Println(body)
	// 添加到数据库
	if err := models.CreatePost(&body); err != nil {
		code = errmsg.ERROR_ARTICLE_ADD_FAIL
		h.Fails(c, h.Response{Code: code})
	}
	h.Success(c, h.Response{Msg: "success", Data: body})
}

// PostList
// @Summary 文章列表
// @Tags 文章
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
// @Router /posts [get]
func PostList(c *gin.Context) {
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
	// 将字符串进行切片 "[0,10]"
	var code int
	article, total, error := models.PostList(limit, skip, sortArr)
	if error != nil {
		code = errmsg.ERROR_ARTICLE_CONTENT
		c.JSON(http.StatusBadRequest, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}
	// 添加一个返回协议头
	c.Header("Content-Range", fmt.Sprintf("%d-%d/%d", skip, skip+len(*article), total))
	c.Header("Content-Type", "application/json")
	h.Success(c, h.Response{Msg: "success", Data: article, Total: total})
}

// PostCreate
// @Summary 获取文章
// @Tags 文章
// @Param id path uint true "文章ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
// @Router /posts/{id} [post]
func PostInfo(c *gin.Context) {
	var (
		code int
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error(), "code": code})
		return
	}
	article, err := models.GetPostById(id)
	if err != nil {
		code = errmsg.ERROR_ART_NOT_EXIST
		c.JSON(http.StatusBadRequest, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}
	code = errmsg.SUCCESS
	h.Success(c, h.Response{Data: article, Code: code})
}

// PostDelete
// @Summary 删除文章
// @Tags 文章
// @Param id path uint true "文章ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
// @Router /posts/{id} [delete]
func PostDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var code int
	if err != nil {
		code = errmsg.ERROR_ARTICLE_ERROR
		c.JSON(http.StatusBadRequest, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}
	if err := models.DeletePost(id); err != nil {
		code = errmsg.ERROR_ARTICLE_DELETE_FAIL
		c.JSON(http.StatusBadRequest, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}

	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
}

// PostUpdate
// @Summary 更新 文章
// @Tags 文章
// @Param id path uint true "文章ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
// @Router /posts/{id} [put]
func PostUpdate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var code int
	if err != nil {
		code = errmsg.ERROR_ARTICLE_ERROR
		c.JSON(http.StatusBadRequest, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}
	data := models.Article{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error(), "code": code})
		return
	}

	if err := models.UpdatePost(id, &data); err != nil {
		code = errmsg.ERROR_ARTICLE_UPDATE_FAIL
		c.JSON(http.StatusBadRequest, gin.H{"msg": errmsg.GetErrMsg(code), "code": code})
		return
	}

	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, data)

}
