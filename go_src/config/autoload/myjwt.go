package autoload

import (
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"time"
)

// AuthClaims 是 claims struct
type AuthClaims struct {
	UserId uuid.UUID `json:"userId"`
	jwt.StandardClaims
}

// SecretKey 密钥建议单独私有存放
var SecretKey = []byte("sdfasdfasdfasdfaskldfjakls;dfjlak;314u893819asdfasdfas")

// GenerateToken 一般在登录之后使用生成token  能够返回给前端
func GenerateToken(userId uuid.UUID, expireTime time.Time) (string, error) {
	// 创建一个 claim
	claim := AuthClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			//过期时间
			ExpiresAt: expireTime.Unix(),
			// 签名时间
			IssuedAt: time.Now().Unix(),
			// 签名颁发者
			Issuer: "abcnull",
			// 签名主题
			Subject: "gindemo",
		},
	}

	// 使用指定的签名方式创建 toklen， 有 1， 2 段内容，第3段内容没有加上
	noSignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// 使用secretKey 密钥进行加密处后拿到最终的 token string
	return noSignedToken.SignedString(SecretKey)
}
