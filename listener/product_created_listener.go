package listener

import (
	"email-service-event-driven/event"
	"email-service-event-driven/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

func OnCreatedProductEvent(consumer *nsq.Consumer, service service.EmailService) {
	consumer.AddHandler(event.ProductCreatedEvent{
		Service: service,
	})
	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	err := consumer.ConnectToNSQLookupd("localhost:4161")
	if err != nil {
		panic(err) //500 Internal Server Error
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	// Gracefully stop the consumer.
	consumer.Stop() //enaknya di main atau di sini (?)

}
