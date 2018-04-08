package main

import (
	"log"
	"sync"
	//   "strconv"

	"github.com/bitly/go-nsq"
)

func receive(topic string, wg *sync.WaitGroup, callback func(string, string)) {
	// wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer(topic, "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		msgString := string(message.Body)
		log.Printf("Got a message: %s", msgString)

		// Callback to provided function in init
		callback(topic, msgString)

		wg.Done()
		return nil
	}))
	err := q.ConnectToNSQD(connectionString)
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()
}

func initReceiver(topic string, callback func(string, string)) {
	wg := &sync.WaitGroup{}
	for {
		receive(topic, wg, callback)
		defer wg.Done()
		// log.Printf("seq: " + strconv.Itoa(seq))
	}
}
