package api

import (
	"encoding/json"
	"fmt"
	"goVueBlog/globar"
	"goVueBlog/service/serializer"
	"goVueBlog/utils"
	"goVueBlog/utils/errmsg"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: globar.Logger,
	}
}

type BindRequestOtpons struct {
	Ctx     *gin.Context
	Ser     any
	BindUri bool
}

func (b *BaseApi) BindResquest(option BindRequestOtpons) *BaseApi {
	var errResult error
	b.Ctx = option.Ctx

	if option.Ser != nil {
		if option.BindUri {
			errResult = utils.AppendError(errResult, b.Ctx.ShouldBindUri(option.Ser))
		} else {
			errResult = utils.AppendError(errResult, b.Ctx.ShouldBind(option.Ser))
		}
		if errResult != nil {
			b.AddError(errResult)
			b.Fail(utils.Response{
				Code: errmsg.ERROR_REQUERY_ARG_WRONG,
				Msg:  b.GetError().Error(),
			})
		}

	}
	return b
}

func (b *BaseApi) ParseValidateErrors(errs error, target any) error {
	var errResult error
	errValidation, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errs
	}

	// 通过反射获取结构体字段
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errValidation {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}
		if errMessage == "" {
			errMessage = fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
		}
		errResult = utils.AppendError(errResult, fmt.Errorf(errMessage))
	}
	return errResult
}

type BinldQueryOtpons struct {
	Ctx    *gin.Context
	Querys *serializer.CommonQueryOtpones
}

// 对query数据进行解析
func (b *BaseApi) ResolveQueryParams(option BinldQueryOtpons) *BaseApi {
	// 将Query取出来
	b.Ctx = option.Ctx
	// 对数据进行转换
	option.Querys.Filter = utils.StringToJson(b.Ctx.DefaultQuery("filter", "{}")) // 字符串转json
	option.Querys.Ranges = b.ParamQuery(b.Ctx.DefaultQuery("ranges", "[0,10]"))   // 对字符串分页进行转换
	option.Querys.Sort = b.ParamQuery(b.Ctx.DefaultQuery("sort", "['id', 'desc']"))

	return b

}

func (b *BaseApi) ParamQuery(text string) []any {
	var parater []any
	_ = json.Unmarshal([]byte(text), &parater)
	return parater
}

func (b *BaseApi) AddError(err error) {
	b.Errors = utils.AppendError(b.Errors, err)
}

func (b *BaseApi) GetError() error {
	return b.Errors
}

func (b *BaseApi) Fail(resp utils.Response) {
	globar.Logger.Fatalf("错误代码：%d,错误信息：%s", resp.Code, resp.Msg)
	utils.Fails(b.Ctx, resp)
}

func (b *BaseApi) Ok(resp utils.Response) {
	utils.Success(b.Ctx, resp, "")
}

func (m *BaseApi) ServerFail(resp utils.Response) {
	utils.ServerFail(m.Ctx, resp)
}
