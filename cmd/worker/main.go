package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/Dionizio8/go-task/app/worker"
	"github.com/Dionizio8/go-task/infra/kafka"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("start worker")

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	kafkaRead := kafka.InitKafkaConsumer()
	defer kafkaRead.Close()

	messageConsumer := kafka.NewKafkaMessageConsumer(kafkaRead)

	notification := worker.NewNotificationHandler(*messageConsumer)
	notification.Execute(ctx)

	<-ctx.Done()
	stop()

	kafkaRead.Close()
	if err != nil {
		log.Fatal("forced to shutdown: ", err)
	}

	log.Println("exiting")
}
