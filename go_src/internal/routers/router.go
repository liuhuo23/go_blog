package routers

import (
	"github.com/gin-gonic/gin"
	"go_blog/config"
	"go_blog/internal/middleware"
	"go_blog/internal/pkg/errors"
	response2 "go_blog/internal/pkg/response"
	"io/ioutil"
	"net/http"
)

func SetRouters() *gin.Engine {
	var r *gin.Engine

	if config.Config.AppConfig.Debug == false {
		//生产模式
		r = ReleaseRouter()
		r.Use(
			middleware.CustomLogger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
			middleware.RequestCostHandler(),
		)
	} else {
		//开发调试模式
		r = gin.New()
		r.Use(
			middleware.RequestCostHandler(),
			gin.Logger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)
	}
	//设置API路由
	SetApiRouter(r)
	// 设置blog 路由
	SetBlogRouter(r)
	r.NoRoute(func(c *gin.Context) {
		response2.Resp().SetHttpCode(http.StatusNotFound).FailCode(c, errors.NotFound)
	})
	return r
}

// ReleaseRouter 生产模型使用官方建议设置未release 模式
func ReleaseRouter() *gin.Engine {
	//切换到生产模式
	gin.SetMode(gin.ReleaseMode)
	//禁用 gin 输出接口访问日志
	gin.DefaultWriter = ioutil.Discard

	engin := gin.New()

	return engin
}
