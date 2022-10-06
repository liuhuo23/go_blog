package service

import (
	"go_blog/config/autoload"
	"go_blog/internal/model"
	err "go_blog/internal/pkg/errors"
	log "go_blog/internal/pkg/logger"
	"time"
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
	// 设置过期时间为 1分钟
	var expireTime = time.Now().Add(5 * time.Minute)
	log.Logger.Sugar().Info(time.Now(), expireTime)
	token, tokenErr := autoload.GenerateToken(user.ID, expireTime)
	if tokenErr != nil {
		berr := err.NewBusinessError(err.AuthorizationError)
		return nil, berr
	}
	data := Data{
		Token: token,
	}
	return &data, nil
}
func (auth *AuthService) Register(username, password, mobile string) error {
	adminUsers, berr := model.NewAdminUsers()
	if berr != nil {
		log.Logger.Sugar().Info(berr.Error())
	}
	adminUsers.Password = password
	adminUsers.Username = username
	adminUsers.Mobile = mobile
	berr = adminUsers.Register()
	if berr != nil {
		berr := err.NewBusinessError(err.UserISExist)
		return berr
	}
	return nil
}
