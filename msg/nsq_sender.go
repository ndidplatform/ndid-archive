package main

import (
  "log"
  "github.com/bitly/go-nsq"
)

func nsqSend(topic string, msg string) {

  config := nsq.NewConfig()
  w, _ := nsq.NewProducer(connectionString, config)

  err := w.Publish(topic, []byte(msg))
  if err != nil {
      log.Panic("Could not connect")
  }

  w.Stop()
}


