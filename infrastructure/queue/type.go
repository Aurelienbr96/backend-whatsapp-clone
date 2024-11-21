package queue

import (
	"example.com/boiletplate/internal/user/model"
)

type CreatedUserSuccess struct {
	Type    string      `json:"type"`
	Payload *model.User `json:"payload"`
}
