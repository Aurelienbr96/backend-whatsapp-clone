package queue

type IPublisher interface {
	PushMessage(body []byte) error
}

type IConsumer interface {
	Subscribe()
}
