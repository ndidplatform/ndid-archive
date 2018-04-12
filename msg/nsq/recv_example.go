package main

import (
  "os"
  "log"
)

func recvAndPrint(topic string, msg string) {
	log.Printf("topic: " + topic + ", receiveMessage: " + msg)
}

func main() {
	exitSignal := make(chan os.Signal)
	topic := "myTopic_1337"
  channel := "myChannel_1337"
	receivedMessageCallback := recvAndPrint
  targetNSQD := "127.0.0.1:4150"
  
  recv, err := NewReceiver(topic, channel)
  if err != nil {
    log.Panic("Cannot create Receiver")
    return
  }

  recv.SetHandler(receivedMessageCallback)
  err = recv.ConnectToNSQD(targetNSQD)
  if err != nil {
    log.Panic("Cannot connect to NSQD")
    return
  }

	<-exitSignal
}
