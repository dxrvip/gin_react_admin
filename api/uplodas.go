package api

import (
	"goVueBlog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
)

type MyPutRet struct {
	Key  string
	Hash string
	Name string
}

// UploadImage
// @Summary 图片上传鉴权
// @Tags 上传
// @Param fileName path string true "文件名称"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} object
// @Router /auto/upload [get]
func UploadImage(c *gin.Context) {
	// 上传到七牛云
	bucket := viper.GetString("qiniu.BUCKET") // 空间名称

	// 获取文件名称参数
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	// 设置token有效时间
	putPolicy.Expires = viper.GetUint64("qiniu.Expires") //有效时间

	mac := auth.New(viper.GetString("qiniu.AK"), viper.GetString("qiniu.SK"))
	upToken := putPolicy.UploadToken(mac)

	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"msg":    errmsg.GetErrMsg(code),
		"domain": viper.GetString("qiniu.DOMAIN"),
		"token":  upToken,
	})
}
