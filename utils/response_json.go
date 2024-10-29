package utils

import (
	"goVueBlog/utils/errmsg"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data any         `json:"data,omitempty"`
	Msg  string      `json:"message,omitempty"`
	Err  interface{} `json:"err,omitempty"`
}

func (r *Response) IsEmpty() bool {
	return reflect.DeepEqual(r, &Response{})
}

func HttpResopnse(ctx *gin.Context, status int, resp *Response) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(status)
		return
	}

	if resp.Msg == "" {
		resp.Msg = errmsg.GetErrMsg(resp.Code)
	}
	ctx.AbortWithStatusJSON(status, resp)
}

func Success(ctx *gin.Context, resp Response, rs any) {

	// 检查 Data 是否为切片类型
	respDataValue := reflect.ValueOf(resp.Data)
	if respDataValue.Kind() == reflect.Slice {
		// 添加一个返回协议头
		ctx.Header("Content-Range", rs.(string))
		ctx.Header("Content-Type", "application/json")
		if respDataValue.Len() <= 0 {
			resp.Data = []any{}
		}
	}
	HttpResopnse(ctx, http.StatusOK, &resp)
}

func Fails(ctx *gin.Context, resp Response) {
	HttpResopnse(ctx, http.StatusBadRequest, &resp)
}

func ServerFail(ctx *gin.Context, resp Response) {
	HttpResopnse(ctx, http.StatusInternalServerError, &resp)
}
