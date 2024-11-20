package http

import (
	"github.com/gin-gonic/gin"
)

func NewContactHandlers(v1 *gin.RouterGroup, cc *ContactController) {
	contact := v1.Group("/contact")
	{
		contact.GET("/:userId", cc.GetAllUserContacts)
	}
}
