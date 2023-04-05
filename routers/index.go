package routers

import (
	"gin/controllers/index"

	"github.com/gin-gonic/gin"
)

func IndexInit(r *gin.Engine) {
	group := r.Group("index")
	group.GET("/", index.Init)
}
