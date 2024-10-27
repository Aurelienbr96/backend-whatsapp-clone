package user

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/boiletplate/infrastructure/queue"
	"example.com/boiletplate/internal/model"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository UserRepository
	publisher      queue.IPublisher
}

func NewUserController(userRepository UserRepository, publisher queue.IPublisher) *UserController {
	return &UserController{userRepository: userRepository, publisher: publisher}
}

// @BasePath /api/v1

// @Summary Create User
// @Schemes
// @Description Create a new user
// @Tags example
// @Accept json
// @Produce json
// @Param data body model.User true "User Data"
// @Success 201 {string} string "User successfully created"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user [post]
func (co *UserController) CreateOne(c *gin.Context) {
	user := model.User{}

	if c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := co.userRepository.CreateUser(user.PhoneNumber, user.Name)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	msg := queue.CreatedUserSuccess{
		Type:    "created_user",
		Payload: u,
	}

	messageBytes, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to marshal message: %v", err)
		return
	}

	co.publisher.PushMessage(messageBytes)

	c.JSON(http.StatusCreated, gin.H{"message": "User successfully created"})
}

// @Summary Get User
// @Schemes
// @Description Get a user by ID
// @Tags example
// @Accept json
// @Produce json
// @Param id path string true "User ID" example("123e4567-e89b-12d3-a456-426614174000")
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/{id} [get]
func (co *UserController) GetOne(c *gin.Context) {
	id := c.Param("id")
	u, err := co.userRepository.GetOneById(id)
	if err != nil {

		log.Fatalf("Failed to fetch user: %v", err)
		return
	}

	c.JSON(200, u)
}
