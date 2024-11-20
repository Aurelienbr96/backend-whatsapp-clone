package adapter

import (
	"example.com/boiletplate/ent"
	"example.com/boiletplate/internal/user/model"
)

func EntUserAdapter(entUser *ent.User) *model.User {
	return model.NewUser(entUser.ID, entUser.Username, entUser.PhoneNumber, entUser.IsVerified, entUser.Avatar)
}
