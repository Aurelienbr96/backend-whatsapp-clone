package repository

import (
	"context"
	"example.com/boiletplate/internal/contact/model"
	"fmt"

	"example.com/boiletplate/ent"
	"example.com/boiletplate/ent/contact"
	"example.com/boiletplate/ent/user"
	"github.com/google/uuid"
)

type Repository struct {
	client *ent.Client
}

func NewContactRepository(client *ent.Client) *Repository {
	return &Repository{client: client}
}

func (r *Repository) GetContactsByOwnerId(ownerId uuid.UUID) ([]*model.Contact, error) {
	u, err := r.client.Contact.Query().Where(contact.HasOwnerWith(user.IDEQ(ownerId))).WithContactUser(func(q *ent.UserQuery) {
		q.Select(user.FieldUsername, user.FieldPhoneNumber, user.FieldAvatar, user.FieldUsername)
	}).All(context.Background())
	if err != nil {
		return nil, err
	}
	var contacts []*model.Contact
	for _, c := range u {
		contactModel := &model.Contact{
			Avatar:   c.Edges.ContactUser.Avatar,
			Username: c.Edges.ContactUser.Username,
			ID:       c.Edges.ContactUser.ID.String(),
		}
		contacts = append(contacts, contactModel)
	}
	return contacts, err
}

func (r *Repository) CreateMany(contactIds []uuid.UUID, ownerId uuid.UUID) error {
	if len(contactIds) == 0 {
		return fmt.Errorf("no contact IDs provided")
	}
	var entContactCreate []*ent.ContactCreate
	for _, c := range contactIds {
		entContactCreate = append(entContactCreate, r.client.Contact.Create().SetContactUserID(c).SetOwnerID(ownerId))
	}
	if _, err := r.client.Contact.CreateBulk(entContactCreate...).Save(context.Background()); err != nil {
		return fmt.Errorf("failed to create contacts: %w", err)
	}

	return nil
}
