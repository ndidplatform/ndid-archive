# NSQ Messenging Library

## Node can send message following this instruciton
**connectionString:** Consist of IP:PORT of NSQ Deamon

**Topic:** Receiver node id exp: node_id_12344

**msg:**   String message that want to send to node id
```go
package main

var connectionString = "127.0.0.1:32772"

func main() {
  // topic: topic of NSQ
  // msg: String of payload
  topic := "node_id_12344"
  msg := "I'm ok"

  nsqSend(topic, msg)
}
```

## Node can receive message following this instruciton
**connectionString:** consist of IP:PORT of NSQ Deamon

**Topic:** Receiver node id exp: node_id_12344

**receivedMessageCallback:** callback function, will trigger when new message arrive
```go
package main

import (
  "os"
  "log"
)

var connectionString = "127.0.0.1:32772"

func didReceiveMessage(topic string, msg string) {
	log.Printf("topic: " + topic + ", receiveMessage: " + msg)
}

func main() {
	exitSignal := make(chan os.Signal)

	// topic: topic of NSQ
	// receivedMessageCallback: Call when new message arrive
	topic := "node_id_12344"
	receivedMessageCallback := didReceiveMessage
	go initReceiver(topic, receivedMessageCallback)

	<-exitSignal
}
```
