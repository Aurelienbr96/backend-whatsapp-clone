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

func (uRepo *UserRepository) CreateUser(phone_number string) (*ent.User, error) {
	u, err := uRepo.client.User.Create().SetPhoneNumber(phone_number).Save(uRepo.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func (uRepo *UserRepository) GetOneById(id uuid.UUID) (*ent.User, error) {

	u, err := uRepo.client.User.Query().Where(user.IDEQ(id)).First(uRepo.ctx)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uRepo *UserRepository) GetOneByUsername(username string) (*ent.User, error) {
	u, err := uRepo.client.User.Query().Where(user.Username(username)).First(uRepo.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uRepo *UserRepository) GetOneByPhoneNumber(phoneNumber string) (*ent.User, error) {
	u, err := uRepo.client.User.Query().Where(user.PhoneNumber(phoneNumber)).First(uRepo.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uRepo *UserRepository) UpdateOne(id uuid.UUID, username string, phoneNumber string) (*ent.User, error) {
	u, err := uRepo.client.User.UpdateOneID(id).SetUsername(username).SetPhoneNumber(phoneNumber).Save(uRepo.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uRepo *UserRepository) UpdateOneVerifiedStatusById(id uuid.UUID, isVerified bool) (*ent.User, error) {
	u, err := uRepo.client.User.UpdateOneID(id).SetIsVerified(isVerified).Save(uRepo.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uRepo *UserRepository) DeleteOne(id uuid.UUID) (uuid.UUID, error) {

	uRepo.client.User.DeleteOneID(id).Exec(uRepo.ctx)
	return id, nil
}
