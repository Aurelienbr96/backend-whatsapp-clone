package model

import (
	"github.com/google/uuid"
	"regexp"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	UserName    string    `json:"userName"`
	PhoneNumber string    `json:"phoneNumber"`
	IsVerified  bool      `json:"isVerified"`
	Avatar      string    `json:"avatar"`
}

func NewUser(id uuid.UUID, userName string, phoneNumber string, isVerified bool, avatar string) *User {
	return &User{Id: id, UserName: userName, PhoneNumber: phoneNumber, IsVerified: isVerified, Avatar: avatar}
}

func (u *User) HasVerifyAccount() bool {
	return u.IsVerified
}

func (u *User) RemovePhoneNumberWhiteSpace() {
	var nonNumericRegex = regexp.MustCompile(`[^\d+]`)
	u.PhoneNumber = nonNumericRegex.ReplaceAllString(u.PhoneNumber, "")
}
