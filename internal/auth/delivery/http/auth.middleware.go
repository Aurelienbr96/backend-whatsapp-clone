package http

import (
	"example.com/boiletplate/internal/auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString, err := c.Cookie("access-token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
			c.Abort()
			return
		}

		authPayload, err := service.ValidatingJWT(tokenString, ACCESS_TOKEN_SECRET)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("sub", authPayload.Sub)

		c.Next()
	}
}
