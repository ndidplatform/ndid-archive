package main

import (
  "log"
)

func main() {
  topic := "myTopic_1337"
  msg := "I'm ok"
  targetNSQD := "127.0.0.1:4150"

  sender, err := NewSender(targetNSQD)
  if err != nil {
    log.Panic("Cannot create Receiver")
    return
  }

  sender.Send(topic, msg)
}
