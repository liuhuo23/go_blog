package service

import (
	"go_blog/internal/model"
	err "go_blog/internal/pkg/errors"
	log "go_blog/internal/pkg/logger"
)

type AuthService struct {
}
type Data struct {
	Token string `json:"token"`
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (auth *AuthService) Login(username, password string) (*Data, error) {
	// 查询用户是否存在
	adminUsersModel, _ := model.NewAdminUsers()
	user := adminUsersModel.GetUserInfo(username)

	if user == nil {
		berr := err.NewBusinessError(err.UserDoesNotExist)
		return nil, berr
	}

	// 校验密码
	if !adminUsersModel.ComparePasswords(password) {
		return nil, err.NewBusinessError(err.FAILURE, "用户密码错误")
	}

	/* TODO 生成 token 等业务逻辑，此处不再演示，直接返回用户信息 */
	// ...
	data := Data{
		Token: "hello",
	}
	return &data, nil
}
func (auth *AuthService) Register(username, password string) error {
	adminUsers, berr := model.NewAdminUsers()
	if berr != nil {
		log.Logger.Sugar().Info(berr.Error())
	}
	adminUsers.Password = password
	adminUsers.Username = username
	berr = adminUsers.Register()
	if berr != nil {
		berr := err.NewBusinessError(err.UserISExist)
		return berr
	}
	return nil
}
