package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Auth struct{}

func (con Auth) CheckAuthAll(ctx *gin.Context) {
	// ctx.Set("username", "赵四") //中间件中设置的值，在其他中间件或控制器中可以获取到该值
	cpy := ctx.Copy() //中间件中的ctx必须复制才能在协程中使用
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("url:" + cpy.Request.URL.Path)
	}()
	fmt.Println("鉴权All----start")
	ctx.Next() //跳过以下内容，执行其他的代码，最后倒叙输出
	fmt.Println("鉴权All----end")
}

func (con Auth) CheckAuth1(ctx *gin.Context) {
	fmt.Println("鉴权1----start")
	ctx.Next() //跳过以下内容，执行其他的代码
	fmt.Println("鉴权1----end")
}

func (con Auth) CheckAuth2(ctx *gin.Context) {
	fmt.Println("鉴权2----start")
	ctx.Next() //跳过以下内容，执行其他的代码
	fmt.Println("鉴权2----end")
}

// admin分组中间件
func (con Auth) CheckAdmin(ctx *gin.Context) {
	ctx.Set("username", "赵四") //中间件中设置的值，在其他中间件或控制器中可以获取到该值
	fmt.Println("鉴权admin----start")
	ctx.Next() //跳过以下内容，执行其他的代码
	fmt.Println("鉴权admin----end")
}
