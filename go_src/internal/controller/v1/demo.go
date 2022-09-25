package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_blog/internal/controller"
)

type DemoController struct {
	controller.Api
}

func NewDemoController() *DemoController {
	return &DemoController{}
}

// HelloWorld hello world
func (api *DemoController) HelloWorld(c *gin.Context) {
	//获取参数
	str, ok := c.GetQuery("name")
	if !ok {
		str = "失败"
	}
	api.Success(c, fmt.Sprintf("hello %s", str))
}
