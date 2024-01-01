package main

import (
	"log"
	_ "sync"

	"github.com/nsqio/go-nsq"
)

func main() {

	// wg := &sync.WaitGroup{}
	// wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		// log.Printf("Got a message: %v", message)
		msgID := string(message.ID[:])
		msgBody := string(message.Body)
		log.Printf("Got a message: id:%v body:%v\n", msgID, msgBody)
		// wg.Done()
		return nil
	}))
	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	// wg.Wait()

	// q.Stop()
	select {} // block forever
}
