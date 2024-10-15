package utils

import (
	"goVueBlog/utils/errmsg"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code  int         `json:"code"`
	Data  any         `json:"data,omitempty"`
	Msg   string      `json:"message,omitempty"`
	Err   interface{} `json:"err,omitempty"`
	Range string      `json:"range,omitempty"`
}

func (r *Response) IsEmpty() bool {
	return reflect.DeepEqual(r, &Response{})
}

func HttpResopnse(ctx *gin.Context, status int, resp *Response) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(status)
		return
	}

	resp.Msg = errmsg.GetErrMsg(resp.Code)
	ctx.AbortWithStatusJSON(status, resp)
}

func Success(ctx *gin.Context, resp Response, rs string) {
	if rs != "" {
		// 添加一个返回协议头
		ctx.Header("Content-Range", rs)
		ctx.Header("Content-Type", "application/json")
	}
	HttpResopnse(ctx, http.StatusOK, &resp)
}

func Fails(ctx *gin.Context, resp Response) {
	HttpResopnse(ctx, http.StatusBadRequest, &resp)
}

func ServerFail(ctx *gin.Context, resp Response) {
	HttpResopnse(ctx, http.StatusInternalServerError, &resp)
}
