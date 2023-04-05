package routers

import (
	"gin/controllers/admin"
	"gin/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminInit(r *gin.Engine) {
	//分组
	group := r.Group("admin", middlewares.Auth{}.CheckAdmin) //分组中间件
	group.GET("user/getOne", admin.UserController{}.GetOne)
	group.GET("user/getById/:id", admin.UserController{}.GetById) //动态路由
	group.GET("user/list", admin.UserController{}.List)
	group.POST("user/add", admin.UserController{}.Add)
	group.POST("user/del", admin.UserController{}.Del)
	group.POST("user/upd", admin.UserController{}.Upd)
	group.POST("user/getShouldBind", admin.UserController{}.GetShouldBind) //数据绑定
}
