package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/Dionizio8/go-task/app/worker"
	"github.com/Dionizio8/go-task/infra/kafka"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("start worker")

	godotenv.Load(path.Join(os.Getenv("HOME"), "/go/src/github.com/Dionizio8/go-task/.env"))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	kafkaRead := kafka.InitKafkaConsumer()
	defer kafkaRead.Close()

	messageConsumer := kafka.NewKafkaMessageConsumer(kafkaRead)

	notification := worker.NewNotificationHandler(*messageConsumer)
	notification.Execute(ctx)

	<-ctx.Done()
	stop()

	err := kafkaRead.Close()
	if err != nil {
		log.Fatal("forced to shutdown: ", err)
	}

	log.Println("exiting")
}
