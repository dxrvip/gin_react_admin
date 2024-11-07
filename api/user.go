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

	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &params, BindUri: false}).GetError(); err != nil {

		m.Fail(utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	if params.Password != params.RePassword {
		m.Fail(utils.Response{Msg: "二次密码不正确！"})
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
		m.Fail(utils.Response{Code: errmsg.ERROR_USERNAME_USED, Msg: err.Error()})
		return
	}

	m.Ok(utils.Response{
		Code: errmsg.SUCCESS,
		Data: token,
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

	if err := m.BindResquest(BindRequestOtpons{Ctx: c, Ser: &params, BindUri: false}).GetError(); err != nil {

		m.Fail(utils.Response{Code: errmsg.ERROR_USERNAME_OR_PASSWORD_WRONG, Msg: err.Error()})
		return
	}
	// 获取数据库中密码
	user, _ := m.Service.GetUserByUsername(params.Username)
	if user.Username == "" {
		m.Fail(utils.Response{Msg: "用户不存在！"})
		return
	}
	// 验证密码是否正确
	if !utils.CheckPassword(params.Password, user.Password) {
		m.Fail(utils.Response{Msg: "密码错误！"})
		return
	}

	// 生成token
	token, err := utils.GenerateToken(params.Username, int(user.ID))
	if err != nil {
		m.Fail(utils.Response{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "ok",
		"token":  token,
		"data":   user,
	})
}

// List
// @Summary 用户列表
// @Tags 用户
// @Param Authorization header string true "Bearer token"
func (m *UserApi) List(c *gin.Context) {
	var quers serializer.CommonQueryOtpones
	if err := m.ResolveQueryParams(BinldQueryOtpons{Ctx: c, Querys: &quers}).GetError(); err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}

	var datas []service.ResponseUser
	re, err := m.Service.List(&datas, &quers)
	if err != nil {
		m.Fail(utils.Response{Code: errmsg.ERROR_USER_NOT_EXIST, Msg: err.Error()})
		return
	}
	m.Ok(utils.Response{Code: errmsg.SUCCESS, Data: datas}, re)
}

// Info
// @Summary 获取用户信息
// @Tags 用户
// @Param Authorization header string true "Bearer token"
// @Success 200
// @Router /user/info [get]
func (m *UserApi) Info(g *gin.Context) {
	username := g.MustGet("username")
	// fmt.Println(claims)
	if user, _ := m.Service.GetUserByUsername(username.(string)); user.Username != "" {
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
	id, _ := strconv.Atoi(c.Param("id"))
	var json service.RegisterData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if m.Service.UpdateUser(uint(id), json.Username) == nil {
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
