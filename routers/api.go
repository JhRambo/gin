package routers

import (
	"gin/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiInit(r *gin.Engine) {
	//分组
	group := r.Group("api")
	group.GET("myjson", api.Api{}.MyJson)
	group.GET("myjsonp", api.Api{}.MyJsonp)
	group.GET("mystring", api.Api{}.MyString)
}
