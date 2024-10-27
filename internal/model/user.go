package model

type User struct {
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Name        string `json:"name" binding:"required"`
}

type UserWithId struct {
	User
	Id string `json:"id"`
}
