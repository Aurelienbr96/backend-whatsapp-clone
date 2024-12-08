package server

import (
	"example.com/boiletplate/infrastructure/upload-blob/port"
	authHttp "example.com/boiletplate/internal/auth/delivery/http"
	authUseCase "example.com/boiletplate/internal/auth/usecase"
	contactHttp "example.com/boiletplate/internal/contact/delivery/http"
	cRepository "example.com/boiletplate/internal/contact/repository"
	conversationHttp "example.com/boiletplate/internal/conversations/delivery/http"
	userHttp "example.com/boiletplate/internal/user/delivery/http"
	"example.com/boiletplate/internal/user/repository"
	userUseCase "example.com/boiletplate/internal/user/usecase"

	"example.com/boiletplate/docs"
	otphandler "example.com/boiletplate/infrastructure/OTPHandler"
	"example.com/boiletplate/infrastructure/queue"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHandlers(s *Server, publisher queue.IPublisher, otpHandlers otphandler.OTPHandler, blobHandler port.HandleBlobPort) {
	// repository
	userRepository := repository.NewUserRepository(s.entClient)
	contactRepository := cRepository.NewContactRepository(s.entClient)

	// use cases
	loginUseCase := authUseCase.NewLoginUserUseCase(userRepository, otpHandlers)
	createUserUseCase := userUseCase.NewCreateUserUseCase(userRepository, publisher)
	updateUserUseCase := userUseCase.NewUpdateUserUseCase(userRepository, blobHandler)
	syncContactsUseCase := userUseCase.NewSyncContactUseCase(contactRepository, userRepository)

	// controllers
	userController := userHttp.NewUserController(
		userRepository,
		contactRepository,
		createUserUseCase,
		updateUserUseCase,
		syncContactsUseCase,
	)
	authController := authHttp.NewAuthController(userRepository, otpHandlers, loginUseCase)
	contactController := contactHttp.NewContactController(contactRepository)
	conversationController := conversationHttp.NewMessageController()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := s.Gin.Group("/api/v1")

	contactHttp.NewContactHandlers(v1, contactController)
	userHttp.NewUserHandlers(v1, userController)
	authHttp.NewAuthHandlers(v1, authController)
	conversationHttp.NewConversationHandlers(v1, conversationController)

	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Title = "Boilerplate API"

	s.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
