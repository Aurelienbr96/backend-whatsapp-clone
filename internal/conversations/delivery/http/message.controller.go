package http

import (
	"encoding/json"
	"example.com/boiletplate/internal/conversations/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type MessageController struct {
}

func NewMessageController() *MessageController {
	return &MessageController{}
}

// map conv uuid: Conv model
var currentConversations map[string]model.Conversation

// map conv uuid: Message model
var currentMessages map[string][]model.Message

// @Summary Get conversation message
// @Schemes
// @Description Get conversation message
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} []model.Message
// @Router /message/{conversationId} [get]
func (a *MessageController) GetConversationMessages(c *gin.Context) {
	sub, exist := c.Get("sub")
	if currentConversations == nil {
		c.JSON(http.StatusInternalServerError, "no conversation")
		return
	}
	if currentMessages == nil {
		c.JSON(http.StatusInternalServerError, "no messages")
	}

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	conversationId := c.Param("conversationId")

	var conv model.Conversation

	if _, ok := currentConversations[conversationId]; !ok {
		c.JSON(http.StatusNotFound, "conversation not found")
		return
	} else {
		conv = currentConversations[conversationId]
	}

	if !conv.IsAllowedToSeeConversation(sub.(string)) {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	c.JSON(http.StatusOK, currentMessages[conversationId])
	return
}

// @Summary Get list of conversations
// @Schemes
// @Description Get list of conversations
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} model.Conversation
// @Router /conversations [get]
func (a *MessageController) GetUsersConversations(c *gin.Context) {
	sub, exist := c.Get("sub")
	if currentConversations == nil {
		c.JSON(http.StatusInternalServerError, "no conversation")
		return
	}

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	userConversation := make([]model.Conversation, 0)

	for i, conv := range currentConversations {
		if conv.IsAllowedToSeeConversation(sub.(string)) {
			userConversation = append(userConversation, currentConversations[i])
		}
	}

	c.JSON(http.StatusOK, userConversation)
	return
}

type IncomingMessage struct {
	ConversationId string
	Message        string
	SentTo         string
}

func (a *MessageController) SendMessage(c *gin.Context, userUuid string, im *IncomingMessage) *model.Message {
	_, exist := currentConversations[im.ConversationId]
	if currentConversations == nil || !exist {
		c.JSON(http.StatusInternalServerError, "no conversation")
		return nil
	}

	userConversation := make([]model.Conversation, 0)

	for i, conv := range currentConversations {
		if conv.IsAllowedToSeeConversation(userUuid) {
			userConversation = append(userConversation, currentConversations[i])
		}
	}

	msg, _ := model.NewMessage(im.ConversationId, im.Message, im.SentTo)

	return msg
}

var connectedUsers = make(map[string]*websocket.Conn)

func closeWsConn(conn *websocket.Conn, userUuid string) {
	conn.Close()
	delete(connectedUsers, userUuid)
}

func (a *MessageController) Connect(c *gin.Context) {
	sub, exist := c.Get("sub")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	userUuid := sub.(string)
	conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	connectedUsers[userUuid] = conn
	defer closeWsConn(conn, userUuid)
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var incomingMsg IncomingMessage
		if err := json.Unmarshal(p, &incomingMsg); err != nil {
			log.Println("Invalid message format:", err)
			continue
		}
		msg := a.SendMessage(c, userUuid, &incomingMsg)

		sendTo, exist := connectedUsers[incomingMsg.SentTo]
		if exist {
			jsonMsg, _ := json.Marshal(msg)
			if err := sendTo.WriteMessage(messageType, jsonMsg); err != nil {
				log.Println(err)
				return
			}
		}
	}
}
