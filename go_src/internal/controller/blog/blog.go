package blog

import (
	"github.com/gin-gonic/gin"
	"go_blog/internal/controller"
	log "go_blog/internal/pkg/logger"
	"go_blog/internal/service"
	"go_blog/internal/validator"
	"go_blog/internal/validator/form"
	"strconv"
)

type BlogController struct {
	controller.Api
}
type Data struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

func NewBlogController() *BlogController {
	return &BlogController{}
}

func (blog *BlogController) QueryBlog(c *gin.Context) {
	idStr := c.Param("id")
	id, berr := strconv.ParseUint(idStr, 10, 64)
	if berr != nil {
		blog.Err(c, berr)
		return
	}
	data, berr := service.NewBlogServer().QueryById(uint(id))
	if berr != nil {
		log.Logger.Sugar().Error("查询失败！err:", berr)
		blog.Err(c, berr)
		return
	}
	blog.Success(c, data)
}

func (blog *BlogController) InsertOne(c *gin.Context) {
	blogForm := form.NewBlogForm()
	if berr := validator.CheckPostParams(c, &blogForm); berr != nil {
		return
	}
	id, berr := service.NewBlogServer().CreateBlog(blogForm.Title, []byte(blogForm.Content), blogForm.Author, blogForm.Visits)
	if berr != nil {
		blog.Err(c, berr)
		return
	}
	blog.Success(c, id)
}

func (blog *BlogController) GetList(c *gin.Context) {
	Page := validator.Page{}
	if berr := validator.CheckQueryParams(c, &Page); berr != nil {
		return
	}
	data := service.NewBlogServer().GetList(int(Page.Page), int(Page.Limit))
	blog.Success(c, data)
}
