package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct{}

func (con Api) MyJson(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	v, _ := username.(string)
	fmt.Println(v + "：Api-MyJson") //获取中间件设置的值
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "我是json数据",
	})
}

func (con Api) MyJsonp(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, gin.H{
		"success": true,
		"message": "我是jsonp数据",
	})
}

func (con Api) MyString(ctx *gin.Context) {
	ctx.String(http.StatusOK, "我是string数据")
}
