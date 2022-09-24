package routers

import "github.com/gin-gonic/gin"
import controllerV1 "go_blog/internal/controller/v1"

func SetApiRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		auth := controllerV1.NewAuthController()
		v1.POST("/login", auth.Login)
		demo := controllerV1.NewDemoController()
		v1.GET("/hello-world", demo.HelloWorld)
		v1.POST("/register", auth.NewRegister)
	}
}
