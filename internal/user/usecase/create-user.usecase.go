package usecase

import (
	"encoding/json"
	"example.com/boiletplate/ent"
	"example.com/boiletplate/infrastructure/queue"
	"example.com/boiletplate/internal/user/model"
	"example.com/boiletplate/internal/user/repository"
	"log"
)

type CreateUserUseCase struct {
	userRepository *repository.Repository
	publisher      *queue.Publisher
}

func NewCreateUserUseCase(userRepository *repository.Repository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: userRepository}
}

type UserToCreate struct {
	PhoneNumber string
}

type Input struct {
	UserToCreate UserToCreate
}

func (u *CreateUserUseCase) Execute(input Input) (model.User, error) {
	uuid := model.GenerateRandomUUid()
	userToCreate := model.NewUser(uuid, input.UserToCreate.PhoneNumber, "", false, "")
	userToCreate.RemovePhoneNumberWhiteSpace()
	entUser, err := u.userRepository.CreateUser(userToCreate.PhoneNumber)

	if err != nil {
		switch {
		case ent.IsConstraintError(err):
			// c.JSON(http.StatusConflict, gin.H{"message": "Phone number already used"})
			return model.User{}, err
		}
		// c.JSON(http.StatusInternalServerError, err.Error())
		return model.User{}, err
	}

	createdUser := model.NewUser(entUser.ID, entUser.Username, entUser.PhoneNumber, entUser.IsVerified, entUser.Avatar)

	msg := queue.CreatedUserSuccess{
		Type:    "created_user",
		Payload: createdUser,
	}

	messageBytes, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to marshal message: %v", err)
		return model.User{}, err
	}

	u.publisher.PushMessage(messageBytes)
	return model.User{}, nil
}
