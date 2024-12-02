package model

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ConversationNotCreated = errors.New("message not part of a conversation")
)

type Message struct {
	Id             uuid.UUID `json:"id"`
	ConversationId string    `json:"conversationId"`
	Content        string    `json:"content"`
	SenderId       string    `json:"senderId"`
	Read           bool      `json:"read"`
	SentAt         time.Time `json:"sentAt"`
	DeletedBy      []string  `json:"deleted_by"`
}

func NewMessage(conversationId string, content string, senderId string) (*Message, error) {

	return &Message{
		Id:             uuid.New(),
		ConversationId: conversationId,
		Content:        content,
		SenderId:       senderId,
		Read:           false,
		SentAt:         time.Now(),
		DeletedBy:      []string{},
	}, nil
}

func (m *Message) SetRead(read bool) {
	m.Read = read
}

func (m *Message) SetDeletedBy(deletedBy []string) {
	m.DeletedBy = deletedBy
}
