package usecase

import (
	"example.com/boiletplate/ent"
	"example.com/boiletplate/internal/user/adapter"
	"example.com/boiletplate/internal/user/model"
	"example.com/boiletplate/internal/user/repository"
	"github.com/google/uuid"
)

type UpdateUserUseCase struct {
	userRepository *repository.Repository
}

func NewUpdateUserUseCase(userRepository *repository.Repository) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepository: userRepository}
}

type UserToUpdate struct {
	Id          uuid.UUID
	PhoneNumber string
}

type UpdateUserUseCaseInput struct {
	UserToUpdate UserToUpdate
}

func (u *UpdateUserUseCase) Execute(input UpdateUserUseCaseInput) (*model.User, error) {
	entUser, err := u.userRepository.GetOneByPhoneNumber(input.UserToUpdate.PhoneNumber)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, err
		}
	}
	user := adapter.EntUserAdapter(entUser)
	user.RemovePhoneNumberWhiteSpace()

	_, err = u.userRepository.UpdateOne(user.Id, user.UserName, user.PhoneNumber)
	return user, err
}
