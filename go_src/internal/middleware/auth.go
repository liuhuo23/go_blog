package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go_blog/config/autoload"
	"go_blog/internal/pkg/errors"
	log "go_blog/internal/pkg/logger"
	r "go_blog/internal/pkg/response"
	"time"
)

// JWTAuth 鉴权中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := r.Resp()
		//获取请求头中 token ，实际是一个完整签名过的token； a  complete, signed token
		tokenStr, _ := c.Cookie("Authorization")
		log.Logger.Sugar().Info("tokenStr,", tokenStr)
		if tokenStr == "" {
			response.FailCode(c, errors.NOTOKEN)
			c.Abort()
			return
		}

		//解析拿到完整有效的token， 其中包含解析后的 3 segment
		token, err := ParseToken(tokenStr)
		log.Logger.Sugar().Info("token:", token)
		claims, ok := token.Claims.(*autoload.AuthClaims)
		log.Logger.Sugar().Info("claims", claims.ExpiresAt, time.Unix(claims.ExpiresAt, 0).Format("2006/01/02 15:04:05"))
		if err != nil {
			response.FailCode(c, errors.NOTOKEN)
			c.Abort()
			return
		}
		// 获取token中的claims
		claims, ok = token.Claims.(*autoload.AuthClaims)
		if !ok {
			response.FailCode(c, errors.TOKENERROR)
			c.Abort()
			return
		}
		// 将claims 中的用户信息存储在 context 中
		c.Set("userId", claims.UserId)

		// 这里执行路由 HandlerFunc
		c.Next()
	}
}

// ParseToken 解析 token string =》*jwt.Token
func ParseToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &autoload.AuthClaims{}, func(tk *jwt.Token) (interface{}, error) {
		return autoload.SecretKey, nil
	})
}
