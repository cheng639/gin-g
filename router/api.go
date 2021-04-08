package router

import (
	"gin-g/controller/api"
	"gin-g/middleware"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(Router *gin.RouterGroup) {
	//初始化api路由组并使用Verify中间件检查token
	r := Router.Group("api").Use(middleware.Verify())
	{
		r.POST("smsCode", api.SendCode)
		r.POST("users/login", api.NewUserController().Login)

		r.GET("books", api.NewBookController().Index)
		r.POST("books", api.NewBookController().Store)
		r.PUT("books/:id", api.NewBookController().Update)
		r.GET("books/:id", api.NewBookController().Show)
		r.DELETE("books/:id", api.NewBookController().Destroy)
	}
}
