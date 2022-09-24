package v1

import (
	"github.com/gin-gonic/gin"
	"go_blog/internal/controller"
	"go_blog/internal/service"
	"go_blog/internal/validator"
	"go_blog/internal/validator/form"
)

type AuthController struct {
	controller.Api
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// login 一个炮筒去全流程的示例， 业务代码未补充完整
func (api *AuthController) Login(c *gin.Context) {
	// 初始化参数结构体
	loginForm := form.LoginForm()

	//绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &loginForm); err != nil {
		return
	}

	result, err := service.NewAuthService().Login(loginForm.UserName, loginForm.Password)

	//根据业务返回值判断业务成功 OR 失败
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c, result)
}
