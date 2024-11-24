package queue

import (
	"example.com/boiletplate/internal/user/entity"
)

type CreatedUserSuccess struct {
	Type    string       `json:"type"`
	Payload *entity.User `json:"payload"`
}
