package main

import (
	"errors"
	"github.com/bitly/go-nsq"
)

type Receiver struct {
	nsqConsumer *nsq.Consumer
	topic       string
	channel     string
	handler     func(string, string)
}

func NewReceiver(topic string, channel string) (*Receiver, error) {
	config := nsq.NewConfig()
	q, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return nil, err
	}

	r := &Receiver{
		nsqConsumer: q,
		topic:       topic,
		channel:     channel,
	}

	return r, nil
}

// Set handler function when messages are received.
//    handler  func(string, string)
func (r *Receiver) SetHandler(handler func(string, string)) {
	r.handler = handler
	r.nsqConsumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		msgString := string(message.Body)
		handler(r.topic, msgString)
		return nil
	}))
}

// Connect the receiver to the NSQD.
//    addr  string  [ip:port]
// return error if unable to connect or handler is nil.
func (r *Receiver) ConnectToNSQD(addr string) error {
	if r.handler == nil {
		return errors.New("Callback is not set")
	}

	err := r.nsqConsumer.ConnectToNSQD(addr)
	if err != nil {
		return err
	}

	return nil
}
