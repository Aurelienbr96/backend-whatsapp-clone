package http

import (
	"encoding/json"
	"example.com/boiletplate/internal/contact/repository"
	http2 "example.com/boiletplate/internal/user"
	"example.com/boiletplate/internal/user/adapter"
	"log"
	"net/http"

	"example.com/boiletplate/ent"
	"example.com/boiletplate/infrastructure/queue"
	"example.com/boiletplate/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userRepository    *http2.Repository
	contactRepository *repository.Repository
	publisher         queue.IPublisher
}

func NewUserController(userRepository *http2.Repository, publisher queue.IPublisher, contactRepository *repository.Repository) *UserController {
	return &UserController{userRepository: userRepository, publisher: publisher, contactRepository: contactRepository}
}

type UserToCreate struct {
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

// @BasePath /api/v1

// @Summary Create User
// @Schemes
// @Description Create a new user
// @Tags example
// @Accept json
// @Produce json
// @Param data body UserToCreate true "User Data"
// @Success 201 {string} string "User successfully created"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user [post]
func (co *UserController) CreateOne(c *gin.Context) {

	user := UserToCreate{}

	if c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := co.userRepository.CreateUser(utils.RemoveWhiteSpace(user.PhoneNumber))

	if err != nil {
		switch {
		case ent.IsConstraintError(err):
			c.JSON(http.StatusConflict, gin.H{"message": "Phone number already used"})
			return
		}
		c.JSON(http.StatusInternalServerError, err.Error())
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

	c.JSON(http.StatusCreated, u)
}

type ContactsToSync struct {
	ContactIds []string `json:"contactIds" binding:"required"`
	OwnerId    string   `json:"ownerId" binding:"required"`
}

// @Summary Sync contact
// @Schemes
// @Description Sync contact
// @Tags example
// @Accept json
// @Produce json
// @Param data body ContactsToSync true "Contact Data"
// @Success 201 {string} string "contacts successfully synced"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/sync-contacts [post]
func (co *UserController) SyncContact(c *gin.Context) {

	// TODO get list of actual contact, delete/add new contacts

	contacts := ContactsToSync{}

	if c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&contacts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var contactsUuid []uuid.UUID

	ownedId, _ := uuid.Parse(contacts.OwnerId)

	for _, id := range contacts.ContactIds {
		uuidID, err := uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		contactsUuid = append(contactsUuid, uuidID)
	}

	err := co.contactRepository.CreateMany(contactsUuid, ownedId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "contacts synced successfully"})
}

// @Summary Get User
// @Schemes
// @Description Get user's connected information
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/me [get]
func (co *UserController) GetMe(c *gin.Context) {
	sub, exist := c.Get("sub")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	uuidID, err := uuid.Parse(sub.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	u, err := co.userRepository.GetOneById(uuidID)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			c.JSON(http.StatusNotFound, "Could not found this user")
		}
		return
	}

	userLoggedIn := adapter.EntUserAdapter(u)

	c.JSON(200, userLoggedIn)
}

// @Summary Get User
// @Schemes
// @Description Get a user by ID
// @Tags example
// @Accept json
// @Produce json
// @Param id path string true "User ID" example("123e4567-e89b-12d3-a456-426614174000")
// @Success 200 {object} model.UserWithId
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/{id} [get]
func (co *UserController) GetOneById(c *gin.Context) {
	id := c.Param("id")

	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	u, err := co.userRepository.GetOneById(uuidID)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			c.JSON(http.StatusNotFound, "Could not found this user")
		}
		return
	}

	c.JSON(200, u)
}

// @Summary Get User
// @Schemes
// @Description Get a user by PhoneNumber
// @Tags example
// @Accept json
// @Produce json
// @Param id path string true "User ID" example("+33602222632")
// @Success 200 {object} model.UserWithId
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/by-phone/{phoneNumber} [get]
func (co *UserController) GetOneByPhoneNumber(c *gin.Context) {
	phoneNumber := c.Param("phoneNumber")

	entUser, err := co.userRepository.GetOneByPhoneNumber(utils.RemoveWhiteSpace(phoneNumber))
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			c.String(http.StatusNotFound, "Could not found this user")
		}
		return
	}
	u := adapter.EntUserAdapter(entUser)

	c.JSON(200, u)
}

type UserToUpdate struct {
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	UserName    string `json:"userName" binding:"required"`
}

// @Summary Update User
// @Schemes
// @Description Update a new user
// @Tags example
// @Accept json
// @Produce json
// @Param data body UserToUpdate true "User Data"
// @Param id path string true "User ID" example("123e4567-e89b-12d3-a456-426614174000")
// @Success 201 {string} string "User successfully updated"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/{id} [put]
func (co *UserController) UpdateOne(c *gin.Context) {
	id := c.Param("id")
	userToUpdate := UserToUpdate{}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse uuid"})
		return
	}

	if c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&userToUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entUser, err := co.userRepository.GetOneByPhoneNumber(userToUpdate.PhoneNumber)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			c.JSON(http.StatusNotFound, "could not found this user")
		}
	}
	user := adapter.EntUserAdapter(entUser)
	user.RemovePhoneNumberWhiteSpace()

	_, err = co.userRepository.UpdateOne(uuidID, user.UserName, user.PhoneNumber)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			c.String(http.StatusNotFound, "Could not found this user")
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User successfully updated"})
}

// @Summary Delete User
// @Schemes
// @Description Delete a user
// @Tags example
// @Accept json
// @Produce json
// @Param id path string true "User ID" example("123e4567-e89b-12d3-a456-426614174000")
// @Success 201 {string} string "User successfully deleted"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "Could not find this user"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/{id} [delete]
func (co *UserController) DeleteOne(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse uuid"})
		return
	}
	_, err = co.userRepository.DeleteOne(uuidID)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			c.String(http.StatusNotFound, "Could not find this user")
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User successfully deleted"})
}
