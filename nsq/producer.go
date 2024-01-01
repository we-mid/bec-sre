package main

import (
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	// err := w.Publish("write_test", []byte("test"))
	// if err != nil {
	// 	log.Panic("Could not connect")
	// }

	// w.Stop()
	// select {} // block forever
	for {
		err := w.Publish("write_test", []byte("test"))
		if err != nil {
			log.Panic("Could not connect")
		}
		time.Sleep(time.Second * 3)
	}
}
