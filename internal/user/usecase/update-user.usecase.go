package usecase

import (
	"example.com/boiletplate/ent"
	"example.com/boiletplate/infrastructure/upload-blob/port"
	"example.com/boiletplate/internal/user/adapter"
	"example.com/boiletplate/internal/user/entity"
	"example.com/boiletplate/internal/user/repository"
	"github.com/google/uuid"
)

type UpdateUserUseCase struct {
	userRepository *repository.Repository
	blobHandler    port.HandleBlobPort
}

func NewUpdateUserUseCase(userRepository *repository.Repository, blobHandler port.HandleBlobPort) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepository: userRepository, blobHandler: blobHandler}
}

type UserToUpdate struct {
	Id          uuid.UUID
	PhoneNumber string
}

type UpdateUserUseCaseInput struct {
	UserToUpdate UserToUpdate
}

func (u *UpdateUserUseCase) Execute(input UpdateUserUseCaseInput) (*entity.User, error) {
	entUser, err := u.userRepository.GetOneByPhoneNumber(input.UserToUpdate.PhoneNumber)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, err
		}
	}
	userToUpdate := adapter.EntUserAdapter(entUser)
	userToUpdate.RemovePhoneNumberWhiteSpace()

	_, err = u.userRepository.UpdateOne(userToUpdate.Id, userToUpdate.UserName, userToUpdate.PhoneNumber)
	return userToUpdate, err
}
