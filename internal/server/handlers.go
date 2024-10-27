package server

import (
	"context"

	"example.com/boiletplate/docs"
	"example.com/boiletplate/infrastructure/queue"
	"example.com/boiletplate/internal/user"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHandlers(s *Server, publisher *queue.Publisher) {
	userRepository := user.NewUserRepository(context.Background(), s.entClient)
	userController := user.NewUserController(*userRepository, publisher)

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := s.gin.Group("/api/v1")
	{
		eg := v1.Group("/user")
		{
			eg.POST("/", userController.CreateOne)
			eg.GET("/:id", userController.GetOne)
		}
	}
	s.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
