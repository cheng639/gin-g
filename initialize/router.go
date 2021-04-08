package initialize

import (
	"gin-g/global"
	"gin-g/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.CONFIG.Local.Path, http.Dir(global.CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.LOG.Info("use middleware logger")
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	//global.LOG.Info("use middleware cors")

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitApiRouter(PublicGroup)                   // 注册功能api路由
	}

	global.LOG.Info("router register success")
	return Router
}
