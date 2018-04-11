package main

import (
	"errors"
	"github.com/bitly/go-nsq"
)

type Receiver struct {
	nsqConsumer		*nsq.Consumer
	topic					string
	channel				string
	callback  		func(string, string)
}

func NewReceiver(topic string, channel string) (*Receiver, error) {
	config := nsq.NewConfig()
	q, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return nil, err
	}

	r := &Receiver{
		nsqConsumer: q,
		topic: topic,
		channel: channel,
	}

	return r, nil
}

// Set callback function when messages are received.
//  	callback	func(string, string)
func (r *Receiver) SetCallback(callback func(string, string)) {
	r.callback = callback
	r.nsqConsumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		msgString := string(message.Body)
		callback(r.topic, msgString)
		return nil
	}))
}

// Connect the receiver to the NSQD.
//  	targetNSQD	string	[ip:port]
// return error if unable to connect or callback is nil.
func (r *Receiver) ConnectToNSQD(targetNSQD string) error {
	if r.callback == nil {
		return errors.New("Callback is not set")
	}

	err := r.nsqConsumer.ConnectToNSQD(targetNSQD)
	if err != nil {
		return err
	}

	return nil
}
