package usecase

import (
	"example.com/boiletplate/ent"
	"example.com/boiletplate/internal/contact/model"
	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateUser(phoneNumber string) (*ent.User, error)
	GetOneById(id uuid.UUID) (*ent.User, error)
	GetOneByUsername(username string) (*ent.User, error)
	GetOneByPhoneNumber(phoneNumber string) (*ent.User, error)
	UpdateOne(id uuid.UUID, username string, phoneNumber string) (*ent.User, error)
	UpdateOneVerifiedStatusById(id uuid.UUID, isVerified bool) (*ent.User, error)
	DeleteOne(id uuid.UUID) (uuid.UUID, error)
	FindManyByPhoneNumbers(phoneNumbers []string) ([]*ent.User, error)
}

type IContactRepository interface {
	GetContactsByOwnerId(ownerId uuid.UUID) ([]*model.Contact, error)
	CreateMany(contactIds []uuid.UUID, ownerId uuid.UUID) error
}
