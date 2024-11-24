package model

type Contact struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	Avatar      string `json:"avatar"`
}
