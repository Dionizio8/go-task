package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/Dionizio8/go-task/app/api"
	"github.com/Dionizio8/go-task/entity"
	"github.com/Dionizio8/go-task/infra/db"
	"github.com/Dionizio8/go-task/infra/repository"
	"github.com/Dionizio8/go-task/usecase/task"
	"github.com/Dionizio8/go-task/usecase/user"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := db.InitDb()

	if err := db.AutoMigrate(&entity.User{}, &entity.Task{}); err != nil {
		log.Fatalln(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := user.NewService(userRepo)

	taskRepo := repository.NewTaskRepository(db)
	taskService := task.NewService(taskRepo)

	srv, err := api.NewServer(
		api.WithTaskService(*taskService),
		api.WithUserService(*userService),
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
