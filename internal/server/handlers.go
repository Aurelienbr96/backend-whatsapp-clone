package server

import (
	authHttp "example.com/boiletplate/internal/auth/delivery/http"
	"example.com/boiletplate/internal/auth/usecase"
	contactHttp "example.com/boiletplate/internal/contact/delivery/http"
	"example.com/boiletplate/internal/contact/repository"
	userHttp "example.com/boiletplate/internal/user/delivery/http"

	"example.com/boiletplate/docs"
	otphandler "example.com/boiletplate/infrastructure/OTPHandler"
	"example.com/boiletplate/infrastructure/queue"
	"example.com/boiletplate/internal/user"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHandlers(s *Server, publisher *queue.Publisher, otpHandlers otphandler.OTPHandler) {
	userRepository := user.NewUserRepository(s.entClient)
	contactRepository := repository.NewContactRepository(s.entClient)

	loginUseCase := usecase.NewLoginUserUseCase(userRepository, otpHandlers)

	userController := userHttp.NewUserController(userRepository, publisher, contactRepository)
	authController := authHttp.NewAuthController(userRepository, otpHandlers, loginUseCase)
	contactController := contactHttp.NewContactController(contactRepository)

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := s.gin.Group("/api/v1")

	contactHttp.NewContactHandlers(v1, contactController)
	userHttp.NewUserHandlers(v1, userController)
	authHttp.NewAuthHandlers(v1, authController)

	s.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
