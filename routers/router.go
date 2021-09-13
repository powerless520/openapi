/**
* @Author: lik
* @Date: 2021/3/5 19:12
* @Version 1.0
 */
package routers

import (
	"DTCloudAPI/global/constant"
	"DTCloudAPI/global/variable"
	"DTCloudAPI/handler"
	"DTCloudAPI/handler/chaptcha"
	"DTCloudAPI/routers/middleware/authorization"
	"DTCloudAPI/routers/middleware/cors"
	"DTCloudAPI/validator/core/factory"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func InitApiRouter() *gin.Engine {
	var router *gin.Engine

	if variable.ConfigYml.GetBool("AppDebug") == false {
		//1.将日志写入日志文件
		gin.DisableConsoleColor()
		f, _ := os.Create(variable.BasePath + variable.ConfigYml.GetString("Logs.GinLogName"))
		gin.DefaultWriter = io.MultiWriter(f)
		// 2.如果是有nginx前置做代理，基本不需要gin框架记录访问日志，开启下面一行代码，屏蔽上面的三行代码，性能提升 5%
		//gin.SetMode(gin.ReleaseMode)

		router = gin.Default()
	} else {
		// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
		router = gin.Default()
		//pprof.Register(router)
	}

	//根据配置进行设置跨域
	if variable.ConfigYml.GetBool("HttpServer.AllowCrossDomain") {
		router.Use(cors.Next())
	}

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Api 模块接口 hello word！")
	})

	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/description/favicon.ico")

	//  创建一个门户类接口路由组
	vApi := router.Group("/api/v1/")
	{
		vApi.GET("sms", (&chaptcha.Captcha{}).CheckCode)
		vApi.POST("signup", factory.Create(constant.ValidatorPrefix+"UsersSignup"))
		vApi.POST("login", factory.Create(constant.ValidatorPrefix+"UsersLogin"))

	}

	// 【需要token】中间件验证的路由
	vApi.Use(authorization.CheckAuth())

	new(handler.MemberController).HandlerRouter(vApi)

	return router

}
