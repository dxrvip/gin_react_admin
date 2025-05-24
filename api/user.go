/* 用户管理 */
package api

import (
	"goVueBlog/models"
	"goVueBlog/service"
	"goVueBlog/service/serializer"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() *UserApi {
	return &UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

// Register
// @Summary 用户注册
// @Tags 用户
// @Accept json
// @Param register body service.RegisterData true "注册信息"
// @Success 200 {string} userInfo
// @Router /user/register [post]
// @Alias: 注册用户
// @Description: 注册一个用户
func (m *UserApi) Register(c *gin.Context) {
	// 拿到表单信息
	var params service.RegisterData

	if err := m.BindResquest(c, BindRequestOtpons{Ser: &params, BindUri: false}).GetError(); err != nil {

		m.Fail(c, utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	if params.Password != params.RePassword {
		m.Fail(c, utils.Response{Msg: "二次密码不正确！"})
	}
	var registerData models.User = models.User{
		Username: params.Username,
		Password: params.Password,
		Email:    params.Email,
		NikeName: params.NikeName,
		Gender:   params.Gender,
		Active:   params.Active,
	}

	token, err := m.Service.CreateUser(&registerData)
	if err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_USERNAME_USED, Msg: err.Error()})
		return
	}
	m.Ok(c, utils.Response{
		Code: errmsg.SUCCESS,
		Data: map[string]any{"token": token, "id": registerData.ID, "username": registerData.Username},
	}, "")
}

// Login
// @Summary 用户登录
// @Tags 用户
// @Accept json
// @Param login body service.LoginRequest true "登录信息"
// @Success 200
// @Router /user/login [post]
func (m *UserApi) Login(c *gin.Context) {
	var params service.LoginRequest

	if err := m.BindResquest(c, BindRequestOtpons{Ser: &params, BindUri: false}).GetError(); err != nil {

		m.Fail(c, utils.Response{Code: errmsg.ERROR, Msg: err.Error()})
		return
	}
	// 获取数据库中密码
	user, _ := m.Service.GetUserByUsername(params.Username)
	if user.Username == "" {
		m.Fail(c, utils.Response{Msg: "用户不存在！"})
		return
	}
	// 谁他妈的总是改密码
	// 验证密码是否正确///$2a$10$vMa2gs0V85bjxxLEMfdjBOjqNRS4xkRN5Kf1rkwPGcfHvCoN8HCOe
	if !utils.CheckPassword(params.Password, user.Password) {
		m.Fail(c, utils.Response{Msg: "密码错误！"})
		return
	}

	// 生成token
	token, err := utils.GenerateToken(params.Username, int(user.ID))
	if err != nil {
		m.Fail(c, utils.Response{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
		"token":   token,
		"data":    user,
	})
}

// List
// @Summary 用户列表
// @Tags 用户
// @Param Authorization header string true "Bearer token"
func (m *UserApi) List(c *gin.Context) {
	var quers serializer.CommonQueryOtpones
	if err := m.ResolveQueryParams(c, BinldQueryOtpons{Querys: &quers}).GetError(); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}

	var datas []service.ResponseUser
	re, err := m.Service.List(&datas, &quers)
	if err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	m.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: datas}, re)
}

// Info
// @Summary 获取用户信息
// @Tags 用户
// @Param Authorization header string true "Bearer token"
// @Success 200
// @Router /user/{id} [get]
func (m *UserApi) Info(c *gin.Context) {
	// 实现获取用户信息的逻辑

	var id serializer.CommonIDDTO

	if error := m.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); error != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: error.Error()})
		return

	}

	var user service.ResponseUser
	if error := m.Service.GetDataByID(id.ID, &user); error != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: error.Error()})
		return
	}
	m.Ok(c, utils.Response{Code: errmsg.SUCCESS, Data: user}, "")
}

// Delete
// @Summary 删除用户
// @Tags 用户
// @Param Authorization header string true "Bearer token"
// @Param id path int true "用户ID"
// @Success 200
// @Router /user/{id} [delete]
func (m *UserApi) Delete(c *gin.Context) {
	// 实现删除用户的逻辑
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	if m.Service.DeleteUser(uint(i)) == nil {
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
// @Param body body object true "用户信息"
// @Success 200
// @Router /user/{id} [put]
func (m *UserApi) Update(c *gin.Context) {
	// 实现修改用户信息的逻辑
	var id serializer.CommonIDDTO
	if err := m.BindResquest(c, BindRequestOtpons{Ser: &id, BindUri: true}).GetError(); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	// 获取修改数据
	var updateData service.UpdateUser
	if err := m.BindResquest(c, BindRequestOtpons{Ser: &updateData, BindUri: false}).GetError(); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	// 修改数据
	if err := m.Service.UpdateUser(id.ID, updateData); err != nil {
		m.Fail(c, utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	// 组织返回数据
	var response service.ResponseUser
	response.ID = id.ID
	response.Username = updateData.Username
	response.NikeName = updateData.NikeName
	response.Email = updateData.Email
	response.Active = updateData.Active

	m.Ok(c, utils.Response{Code: errmsg.SUCCESS, Msg: "修改成功", Data: response}, "")
}
