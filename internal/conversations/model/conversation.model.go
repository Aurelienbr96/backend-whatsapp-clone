package model

import "github.com/google/uuid"

type Conversation struct {
	Id        string   `json:"id"`
	UserIds   []string `json:"user_id"`
	DeletedBy []string `json:"deleted_by"`
	Messages  Message  `json:"last_msg"`
}

func NewConversation(userIds []string) *Conversation {
	return &Conversation{
		Id:      uuid.New().String(),
		UserIds: userIds,
	}
}

func (c *Conversation) IsAllowedToSeeConversation(userUuid string) bool {
	isAllowedToSeeConversation := false
	for _, convUserId := range c.UserIds {
		if convUserId == userUuid {
			isAllowedToSeeConversation = true
			break
		}
	}
	return isAllowedToSeeConversation
}
