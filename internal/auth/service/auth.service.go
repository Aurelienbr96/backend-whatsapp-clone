package service

import (
	"example.com/boiletplate/internal/auth/model"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func SignInAccessToken(uid uuid.UUID, jwtSecret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": uid,
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"exp": time.Now().Add(time.Minute * 15).Unix(), // token expire in 15 minutes
	})

	fmt.Printf("%v", token)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func SignInRefreshToken(uid uuid.UUID, jwtSecret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": uid,
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // token expire in 1 month
	})

	fmt.Printf("%v", token)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidatingJWT(tokenString string, jwtSecret []byte) (*model.AuthPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		authPayload := model.AuthPayload{Sub: claims["sub"].(string)}
		return &authPayload, nil

	} else {
		return nil, err
	}
}

func SetRefreshTokenCookie(c *gin.Context, refreshToken string) {
	setCookie(
		c,
		"refresh-token",
		refreshToken,
		3600*24*30,
		"/",
		"localhost",
		false,
		true,
	)
}

func SetAccessTokenCookie(c *gin.Context, accessToken string) {
	setCookie(
		c,
		"access-token",
		accessToken,
		60*15,
		"/",
		"localhost",
		false,
		true,
	)
}

func setCookie(c *gin.Context, cookieName string, tokenString string, maxAge int, path string, domain string, secure bool, httpOnly bool) {
	c.SetCookie(
		cookieName,  // Cookie name
		tokenString, // Token value
		maxAge,      // MaxAge in seconds (24 hours)
		path,        // Path
		domain,      // Domain (adjust for your environment)
		secure,      // Secure (only send over HTTPS)
		httpOnly,    // HttpOnly (not accessible via JavaScript)
	)
}
