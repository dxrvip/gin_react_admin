package api

import (
	"fmt"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var systemMenuApi *SystemMenuApi

type SystemMenuApi struct {
	BaseApi
}

func NewSystemMenuApi() *SystemMenuApi {
	if systemMenuApi == nil {
		return &SystemMenuApi{
			BaseApi: NewBaseApi(),
		}
	}

	return systemMenuApi
}

const One int = 1

// SystemMenuList
// @Summary 菜单列表
// @Tags 菜单
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
// @Router /systemMenu [get]
func (m *SystemMenuApi) SystemMenuList(c *gin.Context) {
	m.SetCtx(c)
	currentDir, _ := os.Getwd()
	dir := filepath.Join(currentDir, "/api")
	funcs, err := utils.ParseApiFiles(dir)
	if err != nil {
		return
	}
	tatol := len(funcs)
	rs := fmt.Sprintf("%d-%d/%d", One, One+tatol, One)
	m.Ok(utils.Response{Code: errmsg.SUCCESS, Data: funcs}, rs)
}
