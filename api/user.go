package api

import (
	"goVueBlog/models"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required,len=6"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

type UserUpdateRequest struct {
	Username string `json:"username" binding:"required,len=6"`
}

type RegisterRequest struct {
	LoginRequest
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type UserResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// Register
// @Summary 用户注册
// @Tags 用户
// @Accept json
// @Param register body RegisterRequest true "注册信息"
// @Success 200 {string} userInfo
// @Router /user/register [post]
func Register(c *gin.Context) {
	// 拿到表单信息
	var (
		users models.User

		json RegisterRequest
	)
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Println(json)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 处理表单验证通过
	// 验证用户名是否存在
	user, _ := models.GetUserByUsername(json.Username)
	if user.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户名已存在",
		})
		return
	}

	// 存入数据库

	users.Username = json.Username
	//$2a$10$p29.OzqJkkjLyt4O6gqGYOOo2zSrp.dzPdOz/39SPPwIIMKTxPAhK
	// if hashPassword, err = utils.EncryptPassword(json.Password); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "加密密码失败",
	// 	})
	// 	return
	// }
	users.Password = json.Password
	models.CreateUser(&users)
	// 生成token
	token, err := utils.GenerateToken(users.Username, int(users.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "生成token失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})

}

// Login
// @Summary 用户登录
// @Tags 用户
// @Accept json
// @Param login body LoginRequest true "登录信息"
// @Success 200 {object} UserResponse
// @Router /user/login [post]
func Login(c *gin.Context) {
	var (
		json LoginRequest
		data UserResponse
		code int
	)

	if err := c.ShouldBindJSON(&json); err != nil {
		code = errmsg.ERROR_USERNAME_OR_PASSWORD_WRONG
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"data": data,
		})
		return
	}
	// 获取数据库中密码
	user, err := models.GetUserByUsername(json.Username)
	if err != nil || user == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"data": data,
		})
		return
	}
	// 验证密码是否正确
	if !utils.CheckPassword(json.Password, user.Password) {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"data": data,
		})
		return
	}

	// 生成token
	token, err := utils.GenerateToken(json.Username, int(user.ID))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"data": data,
		})
		return
	}
	data.Username = user.Username
	data.UserID = int(user.ID)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"token":  token,
		"data":   data,
	})
}

// Info
// @Summary 获取用户信息
// @Tags 用户
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} UserResponse
// @Router /user/info [get]
func Info(g *gin.Context) {
	username := g.MustGet("username")
	// fmt.Println(claims)
	if user, _ := models.GetUserByUsername(username.(string)); user.Username != "" {
		g.JSON(http.StatusBadRequest, gin.H{
			"data":    user,
			"message": "获取成功",
		})

	} else {
		g.JSON(http.StatusOK, gin.H{
			"data":    nil,
			"message": "获取失败",
		})
	}
}

// Delete
// @Summary 删除用户
// @Tags 用户
// @Param Authorization header string true "Bearer token"
// @Param id path int true "用户ID"
// @Success 200 {object} UserResponse
// @Router /user/{id} [delete]
func Delete(c *gin.Context) {
	// 实现删除用户的逻辑
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	if models.DeleteUser(uint(i)) == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "删除成功",
		})
		return
	}
}

// Update
// @Summary 修改用户信息
// @Tags 用户
// @Param Authorization header string true "Bearer token"
// @Param id path int true "用户ID"
// @Param body body UserUpdateRequest true "用户信息"
// @Success 200 {object} UserResponse
// @Router /user/{id} [put]
func Update(c *gin.Context) {
	// 实现修改用户信息的逻辑
	id, _ := strconv.Atoi(c.Param("id"))
	var json UserUpdateRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if models.UpdateUser(uint(id), json.Username) == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "修改成功",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "修改失败",
		})
		return
	}
}
