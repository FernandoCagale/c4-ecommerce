package event

import (
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"os"
)

type EventRepository struct {
	uri string
}

func New() *EventRepository {
	fmt.Println(os.Getenv("AMQP_URI"))
	return &EventRepository{
		uri: os.Getenv("AMQP_URI"),
	}
}

func (o *EventRepository) getConfig() (config amqp.Config) {
	return amqp.NewDurablePubSubConfig(o.uri, nil)
}

func (o *EventRepository) message(payload interface{}) (*message.Message, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return message.NewMessage(watermill.NewUUID(), body), nil
}

func (o *EventRepository) Publish(topic string, payload interface{}) (error) {
	publisher, err := amqp.NewPublisher(o.getConfig(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	msg, err := o.message(payload)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = publisher.Publish(topic, msg)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(msg.Payload))
	return nil
}
