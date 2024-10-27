package queue

import "example.com/boiletplate/ent"

type CreatedUserSuccess struct {
	Type    string    `json:"type"`
	Payload *ent.User `json:"payload"`
}
