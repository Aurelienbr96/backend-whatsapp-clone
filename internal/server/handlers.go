package server

import (
	"context"

	"example.com/boiletplate/docs"
	otphandler "example.com/boiletplate/infrastructure/OTPHandler"
	"example.com/boiletplate/infrastructure/queue"
	"example.com/boiletplate/internal/auth"
	"example.com/boiletplate/internal/contact"
	"example.com/boiletplate/internal/user"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHandlers(s *Server, publisher *queue.Publisher, otpHandlers otphandler.OTPHandler) {
	userRepository := user.NewUserRepository(context.Background(), s.entClient)
	contactRepository := contact.NewContactRepository(context.Background(), s.entClient)

	userController := user.NewUserController(userRepository, publisher, contactRepository)
	authController := auth.NewAuthController(userRepository, otpHandlers)
	contactController := contact.NewContactController(contactRepository)

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := s.gin.Group("/api/v1")
	{
		contact := v1.Group("/contact")
		{
			contact.GET("/:userId", contactController.GetAllUserContacts)
		}
		user := v1.Group("/user")
		{
			user.POST("/", userController.CreateOne)
			user.POST("/sync-contacts", userController.SyncContact)
		}
		user.Use(auth.AuthGuard())
		{
			user.GET("/:id", userController.GetOneById)
			user.GET("/me", userController.GetMe)
			user.GET("/by-phone/:phoneNumber", userController.GetOneById)
			user.PUT("/:id", userController.UpdateOne)
			user.DELETE("/:id", userController.DeleteOne)
		}
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/send-code", authController.SendCode)
			auth.POST("/logout", authController.Logout)
			auth.POST("/refresh", authController.RefreshToken)
		}
	}
	s.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
