package utils

import (
	"goVueBlog/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, ResponseData{
		Code: errmsg.SUCCESS,
		Data: data,
		Msg:  errmsg.GetErrMsg(errmsg.SUCCESS),
	})
}

func ResponseError(c *gin.Context, code int, data interface{}) {
	c.JSON(code, ResponseData{
		Code: 200,
		Data: data,
		Msg:  errmsg.GetErrMsg(code),
	})
}

func ResponseAuthError(c *gin.Context, code int, data interface{}) {
	c.JSON(401, ResponseData{
		Code: code,
		Data: data,
		Msg:  errmsg.GetErrMsg(code),
	})
}
