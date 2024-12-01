package usecase

import (
	"encoding/json"
	"example.com/boiletplate/ent"
	"example.com/boiletplate/infrastructure/queue"
	"example.com/boiletplate/internal/user/entity"
	"github.com/pkg/errors"
	"log"
)

type CreateUserUseCase struct {
	userRepository IUserRepository
	publisher      queue.IPublisher
}

func NewCreateUserUseCase(userRepository IUserRepository, publisher queue.IPublisher) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: userRepository, publisher: publisher}
}

type UserToCreate struct {
	PhoneNumber string
}

type Input struct {
	UserToCreate UserToCreate
}

var (
	ErrPhoneNumberAlreadyUsed = errors.New("phone number already used")
	ErrInternalServer         = errors.New("internal server error")
)

func (u *CreateUserUseCase) Execute(input Input) (*entity.User, error) {
	randomUUid := entity.GenerateRandomUUid()
	userToCreate := entity.NewUser(randomUUid, "", input.UserToCreate.PhoneNumber, false, "")
	userToCreate.RemovePhoneNumberWhiteSpace()
	entUser, err := u.userRepository.CreateUser(userToCreate.PhoneNumber)

	if err != nil {
		switch {
		case ent.IsConstraintError(err):
			return &entity.User{}, ErrPhoneNumberAlreadyUsed
		default:
			return &entity.User{}, ErrInternalServer
		}
	}

	createdUser := entity.NewUser(entUser.ID, entUser.Username, entUser.PhoneNumber, entUser.IsVerified, entUser.Avatar)

	msg := queue.NewCreatedUserSuccessMessage(createdUser.PhoneNumber)

	messageBytes, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to marshal message: %v", err)
		return &entity.User{}, ErrInternalServer
	}

	_ = u.publisher.PushMessage(messageBytes)
	return createdUser, nil
}
