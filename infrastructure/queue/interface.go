package queue

type IPublisher interface {
	PushMessage(body []byte)
}
