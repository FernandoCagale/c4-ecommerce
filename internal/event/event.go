package event

type Event interface {
	Publish(topic string, payload interface{}) (error)
}
