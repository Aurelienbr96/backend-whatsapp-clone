package http

import (
	"example.com/boiletplate/internal/auth/delivery/http"
	"github.com/gin-gonic/gin"
)

func NewConversationHandlers(v1 *gin.RouterGroup, conversationController *MessageController) {
	user := v1.Group("/conversation")
	user.Use(http.AuthGuard())
	{
		user.GET("/", conversationController.GetUsersConversations)
		user.POST("/connect", conversationController.Connect)
		user.GET("/:conversationId", conversationController.GetConversationMessages)
	}
}
