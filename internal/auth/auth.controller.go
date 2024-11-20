package auth

import (
	"fmt"
	"net/http"

	"example.com/boiletplate/ent"
	otphandler "example.com/boiletplate/infrastructure/OTPHandler"
	"example.com/boiletplate/internal/user"
	"example.com/boiletplate/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var ACCESS_TOKEN_SECRET = []byte("your_secret_key")
var REFRESH_TOKEN_SECRET = []byte("your_secret_key")

type AuthController struct {
	uRepo      *user.UserRepository
	otpHandler otphandler.OTPHandler
}

func NewAuthController(uRepo *user.UserRepository, otpHandler otphandler.OTPHandler) *AuthController {
	return &AuthController{uRepo: uRepo, otpHandler: otpHandler}
}

type LoginDTO struct {
	Code        string `json:"code" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

func createGinError(msg string) gin.H {
	return gin.H{"message": msg}
}

// @Summary Login User
// @Schemes
// @Description Login a user
// @Tags example
// @Accept json
// @Produce json
// @Param data body LoginDTO true "Login data"
// @Success 201 {string} string "You are logged in"
// @Router /auth/login [post]
func (a *AuthController) Login(c *gin.Context) {

	l := LoginDTO{}

	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	phoneNumber := utils.RemoveWhiteSpace(l.PhoneNumber)

	fmt.Printf("phone number: %s", phoneNumber)

	u, err := a.uRepo.GetOneByPhoneNumber(phoneNumber)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			c.JSON(http.StatusNotFound, createGinError("user not found"))
		}
		return
	}

	err = a.otpHandler.VerififyOTP(phoneNumber, l.Code)
	if err != nil {
		c.JSON(http.StatusForbidden, createGinError("Incorrect code"))
		return
	}

	if !u.IsVerified {
		_, err := a.uRepo.UpdateOneVerifiedStatusById(u.ID, true)
		if err != nil {
			c.JSON(http.StatusInternalServerError, createGinError("could not update user status"))
		}
		return
	}

	accessToken, err := SignInAccessToken(u.ID, ACCESS_TOKEN_SECRET)
	if err != nil {
		c.JSON(http.StatusInternalServerError, createGinError("could not generate token"))
		return
	}

	refreshToken, err := SignInRefreshToken(u.ID, ACCESS_TOKEN_SECRET)
	if err != nil {
		c.JSON(http.StatusInternalServerError, createGinError("could not generate token"))
		return
	}

	SetAccessTokenCookie(c, accessToken)
	SetRefreshTokenCookie(c, refreshToken)

	c.JSON(http.StatusCreated, u)
}

// @Summary Logout User
// @Schemes
// @Description Logout a user
// @Tags example
// @Accept json
// @Produce json
// @Success 201 {string} string "You are logged out"
// @Router /auth/logout [post]
func (a *AuthController) Logout(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{"status": "you are logged out"})
}

type SendCodeDTO struct {
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

// @Summary Login User
// @Schemes
// @Description Login a user
// @Tags example
// @Accept json
// @Produce json
// @Param data body SendCodeDTO true "Send code body"
// @Success 201 {string} string "Code sent s"
// @Router /auth/send-code [post]
func (a *AuthController) SendCode(c *gin.Context) {
	l := SendCodeDTO{}

	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	/* phoneNumber, err := phonenumbers.Parse(l.PhoneNumber, "FR")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isValid := phonenumbers.IsValidNumberForRegion(phoneNumber, "FR")
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number"})
		return
	} */
	err := a.otpHandler.SendOTP(l.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, createGinError(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "code sent successfully"})
}

// @Summary Refresh Access token
// @Schemes
// @Description Refresh Access token
// @Tags example
// @Accept json
// @Produce json
// @Success 201 {string} string ""
// @Router /auth/refresh [post]
func (a *AuthController) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh-token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}

	AuthPayload, err := ValidatingJWT(refreshToken, REFRESH_TOKEN_SECRET)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not validate refresh token"})
		return
	}

	uuidID, err := uuid.Parse(AuthPayload.Sub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse uuid"})
		return
	}

	accessToken, err := SignInAccessToken(uuidID, ACCESS_TOKEN_SECRET)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not sign in access token"})
		return
	}

	SetAccessTokenCookie(c, accessToken)
	c.JSON(http.StatusCreated, gin.H{"message": "Refresh token successfully created"})
}
