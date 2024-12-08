package usecase

import (
	"example.com/boiletplate/ent"
	otphandler "example.com/boiletplate/infrastructure/OTPHandler"
	"example.com/boiletplate/internal/auth/model"
	"example.com/boiletplate/internal/auth/service"
	"example.com/boiletplate/internal/user/adapter"
	"example.com/boiletplate/internal/user/entity"
	"example.com/boiletplate/internal/user/repository"
	"github.com/google/uuid"

	"fmt"
)

var ACCESS_TOKEN_SECRET = []byte("your_secret_key")
var REFRESH_TOKEN_SECRET = []byte("your_secret_key")

type LoginUserUseCase struct {
	uRepo      *repository.Repository
	otpHandler otphandler.OTPHandler
}

func NewLoginUserUseCase(uRepo *repository.Repository, otpHandler otphandler.OTPHandler) *LoginUserUseCase {
	return &LoginUserUseCase{uRepo: uRepo, otpHandler: otpHandler}
}

type AuthWithUser struct {
	Auth *model.Auth
	User *entity.User
}

func (l *LoginUserUseCase) Execute(phoneNumber string, code string) (*AuthWithUser, error) {

	entityUser := entity.NewUser(uuid.New(), "", phoneNumber, false, "")
	entityUser.RemovePhoneNumberWhiteSpace()

	u, err := l.uRepo.GetOneByPhoneNumber(entityUser.PhoneNumber)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	userToLogin := adapter.EntUserAdapter(u)

	err = l.otpHandler.VerifyOTP(phoneNumber, code)
	if err != nil {
		return nil, fmt.Errorf("could not verify code")
	}

	if !userToLogin.HasVerifyAccount() {
		_, err := l.uRepo.UpdateOneVerifiedStatusById(u.ID, true)
		if err != nil {
			return nil, fmt.Errorf("could not update user status")
		}

	}

	accessToken, err := service.SignInAccessToken(u.ID, ACCESS_TOKEN_SECRET)
	if err != nil {
		return nil, fmt.Errorf("could not generate token")
	}

	refreshToken, err := service.SignInRefreshToken(u.ID, REFRESH_TOKEN_SECRET)
	if err != nil {
		return nil, fmt.Errorf("could not generate token")
	}

	authPayload := &model.Auth{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}

	r := &AuthWithUser{
		Auth: authPayload,
		User: userToLogin,
	}

	return r, nil
}
