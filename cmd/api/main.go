package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/Dionizio8/go-task/app/api"
	"github.com/Dionizio8/go-task/app/api/middleware"
	"github.com/Dionizio8/go-task/entity"
	"github.com/Dionizio8/go-task/infra/db"
	"github.com/Dionizio8/go-task/infra/kafka"
	"github.com/Dionizio8/go-task/infra/repository"
	"github.com/Dionizio8/go-task/infra/seed"
	"github.com/Dionizio8/go-task/usecase/task"
	"github.com/Dionizio8/go-task/usecase/user"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("start api")

	godotenv.Load(path.Join(os.Getenv("HOME"), "/go/src/github.com/Dionizio8/go-task/.env"))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db := db.InitDb()

	if err := db.AutoMigrate(&entity.User{}, &entity.Task{}); err != nil {
		log.Fatalln(err)
	}

	seed.NewSeedUser(db).Load()

	kafkaWriter := kafka.InitKafkaProducer()
	defer kafkaWriter.Close()

	messageExecutor := kafka.NewKafkaMessageExecutor(kafkaWriter)

	userRepo := repository.NewUserRepository(db)
	userService := user.NewService(userRepo)

	userMiddler := middleware.NewUserMiddler(userRepo)

	taskRepo := repository.NewTaskRepository(db)
	taskService := task.NewService(taskRepo)

	srv, err := api.NewServer(
		api.WithUserService(*userService),
		api.WithUserMiddler(*userMiddler),
		api.WithTaskService(*taskService),
		api.WithMessageExecutor(*messageExecutor),
	)
	if err != nil {
		log.Fatal("error start server: ", err)
	}

	<-ctx.Done()

	stop()

	err = srv.Close()
	if err != nil {
		log.Fatal("forced to shutdown: ", err)
	}

	log.Println("exiting")
}
