package user

import (
	"context"
	"fmt"
	"log"

	"example.com/boiletplate/ent"
	"example.com/boiletplate/ent/user"
	"github.com/google/uuid"
)

type UserRepository struct {
	ctx    context.Context
	client *ent.Client
}

func NewUserRepository(ctx context.Context, client *ent.Client) *UserRepository {
	return &UserRepository{
		ctx:    ctx,
		client: client,
	}
}

func (uRepo *UserRepository) CreateUser(phone_number string, name string) (*ent.User, error) {
	u, err := uRepo.client.User.Create().SetPhoneNumber(phone_number).SetName(name).Save(uRepo.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func (uRepo *UserRepository) GetOneById(id string) (*ent.User, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	u, err := uRepo.client.User.Query().Where(user.IDEQ(uuidID)).First(uRepo.ctx)

	if err != nil {
		return nil, err
	}

	return u, nil
}
