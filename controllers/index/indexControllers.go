package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(ctx *gin.Context) {
	// 模板渲染
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"message": "你好golang，我是渲染的模板",
	})
}
