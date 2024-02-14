package nats

import (
	"encoding/json"
	"fmt"
	"nats_example/config"
	"time"

	"github.com/nats-io/nats.go"
)

type Nats struct {
	cfg *config.Config
	//sc  stan.Conn
	ns *nats.Conn
}

func NewNats(cfg *config.Config) (*Nats, error) {

	url := fmt.Sprintf("docker-nats://%s:%s", cfg.Nats.Host, cfg.Nats.Port)
	ns, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &Nats{cfg: cfg, ns: ns}, nil
}

func (n *Nats) PublishMessage(subject string, message string) error {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return n.ns.Publish(subject, messageJSON)
}

func (n *Nats) SubscribeAndReceiveMessage(subject string) (string, error) {

	var receivedMessage string

	ch := make(chan string)
	_, err := n.ns.Subscribe(subject, func(msg *nats.Msg) {
		err := json.Unmarshal(msg.Data, &receivedMessage)
		if err != nil {
			fmt.Println("Error unmarshalling message:", err)
			return
		}

		ch <- receivedMessage
	})
	if err != nil {
		fmt.Println("Error in subscribing on subject:", err)
		return receivedMessage, err
	}

	select {
	case receivedMessage = <-ch:
		return receivedMessage, nil
	case <-time.After(60 * time.Second):
		return receivedMessage, nats.ErrTimeout
	}
}
