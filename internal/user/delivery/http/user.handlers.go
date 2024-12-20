package http

import (
	"example.com/boiletplate/internal/auth/delivery/http"
	"github.com/gin-gonic/gin"
)

func NewUserHandlers(v1 *gin.RouterGroup, userController *UserController) {
	user := v1.Group("/user")
	{
		user.POST("/", userController.CreateOne)
		user.POST("/sync-contacts", userController.SyncContact)
		user.GET("/by-phone/:phoneNumber", userController.GetOneByPhoneNumber)
	}
	user.Use(http.AuthGuard())
	{
		user.GET("/me", userController.GetMe)
		user.GET("/:id", userController.GetOneById)
		user.PUT("/:id", userController.UpdateOne)
		user.DELETE("/:id", userController.DeleteOne)
	}
}
