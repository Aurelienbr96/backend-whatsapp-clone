package http

import (
	"github.com/gin-gonic/gin"
)

func NewAuthHandlers(group *gin.RouterGroup, authController *AuthController) {
	g := group.Group("/auth")
	{
		g.POST("/login", authController.Login)
		g.POST("/send-code", authController.SendCode)
		g.POST("/logout", authController.Logout)
		g.POST("/refresh", authController.RefreshToken)
	}
}
