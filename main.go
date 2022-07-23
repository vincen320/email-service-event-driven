package main

import (
	"email-service-event-driven/controller"
	"email-service-event-driven/listener"
	"email-service-event-driven/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nsqio/go-nsq"
	"gopkg.in/gomail.v2"
)

func main() {

	mailer := gomail.NewMessage()
	dialer := gomail.NewDialer(service.CONFIG_SMTP_HOST,
		service.CONFIG_SMTP_PORT,
		service.CONFIG_AUTH_EMAIL,
		service.CONFIG_AUTH_PASSWORD)

	emailService := service.NewEmailService(mailer, dialer)
	consumer, err := nsq.NewConsumer("product-created-event", "emailService", nsq.NewConfig())
	if err != nil {
		panic(err) //500 Internal Server Error
	}

	emailController := controller.NewEmailController()
	router := gin.New()
	router.GET("/ping", emailController.Ping)

	server := http.Server{
		Addr:           ":8083",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//Start Consume
	go listener.OnCreatedProductEvent(consumer, emailService)

	log.Println("Email Service Start in 8083 port")

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			panic("Cannot Start Server " + err.Error()) //500 Internal Server Error
		}
	}()

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	// Gracefully stop the consumer.
	consumer.Stop()
}
