package routers

import "github.com/gin-gonic/gin"
import controllerBlog "go_blog/internal/controller/blog"

func SetBlogRouter(r *gin.Engine) {
	v1 := r.Group("/blog")
	{
		blog := controllerBlog.NewBlogController()
		v1.GET("/query/:id", blog.QueryBlog)
		v1.POST("/upload", blog.InsertOne)
	}
}
