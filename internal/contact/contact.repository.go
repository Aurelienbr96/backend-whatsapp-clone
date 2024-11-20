package contact

import (
	"context"
	"fmt"

	"example.com/boiletplate/ent"
	"example.com/boiletplate/ent/contact"
	"example.com/boiletplate/ent/user"
	"github.com/google/uuid"
)

type ContactRepository struct {
	ctx    context.Context
	client *ent.Client
}

func NewContactRepository(ctx context.Context, client *ent.Client) *ContactRepository {
	return &ContactRepository{ctx: ctx, client: client}
}

func (r *ContactRepository) GetContactsByOwnerId(ownerId uuid.UUID) ([]*ent.Contact, error) {
	u, err := r.client.Contact.Query().Where(contact.HasOwnerWith(user.IDEQ(ownerId))).WithContactUser(func(q *ent.UserQuery) {
		q.Select(user.FieldUsername, user.FieldPhoneNumber, user.FieldAvatar, user.FieldUsername)
	}).All(r.ctx)

	return u, err
}

func (r *ContactRepository) CreateMany(contactIds []uuid.UUID, ownerId uuid.UUID) error {
	if len(contactIds) == 0 {
		return fmt.Errorf("no contact IDs provided")
	}
	var entContactCreate []*ent.ContactCreate
	for _, c := range contactIds {
		entContactCreate = append(entContactCreate, r.client.Contact.Create().SetContactUserID(c).SetOwnerID(ownerId))
	}
	if _, err := r.client.Contact.CreateBulk(entContactCreate...).Save(r.ctx); err != nil {
		return fmt.Errorf("failed to create contacts: %w", err)
	}

	return nil
}
