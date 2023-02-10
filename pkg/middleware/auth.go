package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinisaos/gin-vue-starter/pkg/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		err := utils.ValidateToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Authentication required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
