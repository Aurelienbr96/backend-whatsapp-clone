package queue

type IPublisher interface {
	PushMessage(body []byte)
}

type IConsumer interface {
	Subscribe()
}
