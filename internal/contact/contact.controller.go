package contact

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ContactController struct {
	contactRepository *ContactRepository
}

func NewContactController(contactRepository *ContactRepository) *ContactController {
	return &ContactController{contactRepository: contactRepository}
}

// @Summary Get User's contacts
// @Schemes
// @Description Get a user contacts by ID
// @Tags example
// @Accept json
// @Produce json
// @Param userId path string true "User ID" example("123e4567-e89b-12d3-a456-426614174000")
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /contact/{userId} [get]
func (co *ContactController) GetAllUserContacts(c *gin.Context) {
	id := c.Param("userId")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	userContacts, err := co.contactRepository.GetContactsByOwnerId(uuidID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())

	}
	c.JSON(200, userContacts)
}
