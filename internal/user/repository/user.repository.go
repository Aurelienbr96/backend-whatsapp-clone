package repository

import (
	"context"
	"fmt"
	"log"

	"example.com/boiletplate/ent"
	"example.com/boiletplate/ent/user"
	"github.com/google/uuid"
)

type Repository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (uRepo *Repository) CreateUser(phoneNumber string) (*ent.User, error) {
	u, err := uRepo.client.User.Create().SetPhoneNumber(phoneNumber).Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func (uRepo *Repository) GetOneById(id uuid.UUID) (*ent.User, error) {

	u, err := uRepo.client.User.Query().Where(user.IDEQ(id)).First(context.Background())

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uRepo *Repository) GetOneByUsername(username string) (*ent.User, error) {
	u, err := uRepo.client.User.Query().Where(user.Username(username)).First(context.Background())
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uRepo *Repository) GetOneByPhoneNumber(phoneNumber string) (*ent.User, error) {
	u, err := uRepo.client.User.Query().Where(user.PhoneNumber(phoneNumber)).First(context.Background())
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uRepo *Repository) UpdateOne(id uuid.UUID, username string, phoneNumber string) (*ent.User, error) {
	u, err := uRepo.client.User.UpdateOneID(id).SetUsername(username).SetPhoneNumber(phoneNumber).Save(context.Background())
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uRepo *Repository) UpdateOneVerifiedStatusById(id uuid.UUID, isVerified bool) (*ent.User, error) {
	u, err := uRepo.client.User.UpdateOneID(id).SetIsVerified(isVerified).Save(context.Background())
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uRepo *Repository) DeleteOne(id uuid.UUID) (uuid.UUID, error) {

	err := uRepo.client.User.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return [16]byte{}, err
	}
	return id, nil
}
