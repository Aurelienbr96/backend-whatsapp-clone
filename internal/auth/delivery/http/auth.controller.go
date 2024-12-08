package http

import (
	"example.com/boiletplate/ent"
	otphandler "example.com/boiletplate/infrastructure/OTPHandler"
	"example.com/boiletplate/internal/auth/service"
	"example.com/boiletplate/internal/auth/usecase"
	"example.com/boiletplate/internal/user/adapter"
	"example.com/boiletplate/internal/user/entity"
	"example.com/boiletplate/internal/user/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

var ACCESS_TOKEN_SECRET = []byte("your_secret_key")
var REFRESH_TOKEN_SECRET = []byte("your_secret_key")

type AuthController struct {
	uRepo        *repository.Repository
	otpHandler   otphandler.OTPHandler
	loginUseCase *usecase.LoginUserUseCase
}

func NewAuthController(uRepo *repository.Repository, otpHandler otphandler.OTPHandler, loginUseCase *usecase.LoginUserUseCase) *AuthController {
	return &AuthController{uRepo: uRepo, otpHandler: otpHandler, loginUseCase: loginUseCase}
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
// @Success 201 {object} entity.User
// @Router /auth/login [post]
func (a *AuthController) Login(c *gin.Context) {

	l := LoginDTO{}

	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r, err := a.loginUseCase.Execute(l.PhoneNumber, l.Code)
	if err != nil {
		fmt.Printf("error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	service.SetAccessTokenCookie(c, r.Auth.AccessToken)
	service.SetRefreshTokenCookie(c, r.Auth.RefreshToken)

	c.JSON(http.StatusCreated, r.User)
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

// @Summary Send a code to a User
// @Schemes
// @Description Send an OTP by sms
// @Tags example
// @Accept json
// @Produce json
// @Param data body SendCodeDTO true "Send code body"
// @Success 201 {string} string "Code sent"
// @Router /auth/send-code [post]
func (a *AuthController) SendCode(c *gin.Context) {
	l := SendCodeDTO{}

	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entityUser := entity.NewUser(uuid.New(), "", l.PhoneNumber, false, "")
	entityUser.RemovePhoneNumberWhiteSpace()
	_, err := a.uRepo.GetOneByPhoneNumber(entityUser.PhoneNumber)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		default:
			// Handle other errors
			fmt.Printf("Unexpected error: %v\n", err)
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
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
	err = a.otpHandler.SendOTP(l.PhoneNumber)
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
// @Success 201 {object} entity.User
// @Router /auth/refresh [post]
func (a *AuthController) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh-token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}

	AuthPayload, err := service.ValidatingJWT(refreshToken, REFRESH_TOKEN_SECRET)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not validate refresh token"})
		return
	}

	uuidID, err := uuid.Parse(AuthPayload.Sub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse uuid"})
		return
	}

	u, err := a.uRepo.GetOneById(uuidID)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		default:
			// Handle other errors
			fmt.Printf("Unexpected error: %v\n", err)
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
	}

	user := adapter.EntUserAdapter(u)

	accessToken, err := service.SignInAccessToken(uuidID, ACCESS_TOKEN_SECRET)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not sign in access token"})
		return
	}

	service.SetAccessTokenCookie(c, accessToken)
	c.JSON(http.StatusCreated, user)
}
