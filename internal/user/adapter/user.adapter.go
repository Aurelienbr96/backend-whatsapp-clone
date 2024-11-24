package adapter

import (
	"example.com/boiletplate/ent"
	"example.com/boiletplate/internal/user/entity"
)

func EntUserAdapter(entUser *ent.User) *entity.User {
	return entity.NewUser(entUser.ID, entUser.Username, entUser.PhoneNumber, entUser.IsVerified, entUser.Avatar)
}
