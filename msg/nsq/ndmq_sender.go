package main

import (
	"github.com/bitly/go-nsq"
)

type Sender struct {
	nsqProducer		*nsq.Producer
}

func NewSender(addr string) (*Sender, error) {
	config := nsq.NewConfig()
	q, err := nsq.NewProducer(addr, config)
	if err != nil {
		return nil, err
	}

  // Verify newly created producer
  err = q.Ping()
  if err != nil {
    return nil, err
  }

	r := &Sender{
		nsqProducer: q,
	}

	return r, nil
}

// Send message with specified topic to the configured NSQD.
//  	topic    string
//    msg      string
// return err if msg could not be sent
func (s *Sender) Send(topic string, msg string) error {
  err := s.nsqProducer.Publish(topic, []byte(msg))
  if err != nil {
      return err
  }

  return nil
}
