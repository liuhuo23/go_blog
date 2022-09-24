package middleware

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"

func CorsHandler() gin.HandlerFunc {
	return cors.Default()
}
