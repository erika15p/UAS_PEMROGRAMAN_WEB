package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("user") == "" {
			c.Redirect(302, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
