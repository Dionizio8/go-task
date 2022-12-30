package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Dionizio8/go-task/app/api/middleware"
	"github.com/Dionizio8/go-task/infra/kafka"
	"github.com/Dionizio8/go-task/usecase/task"
	"github.com/Dionizio8/go-task/usecase/user"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer      *http.Server
	taskService     task.Service
	userService     user.Service
	userMiddler     middleware.UserMiddler
	messageExecutor kafka.KafkaMessageExecutor
}

func NewServer(options ...func(server *Server) error) (*Server, error) {
	server := &Server{}
	for _, option := range options {
		err := option(server)
		if err != nil {
			return nil, err
		}
	}

	r := gin.Default()

	server.router(r)
	server.httpServer = &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := server.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return server, nil
}

func WithUserService(userService user.Service) func(server *Server) error {
	return func(server *Server) error {
		server.userService = userService
		return nil
	}
}

func WithUserMiddler(userMiddler middleware.UserMiddler) func(server *Server) error {
	return func(server *Server) error {
		server.userMiddler = userMiddler
		return nil
	}
}

func WithTaskService(taskService task.Service) func(server *Server) error {
	return func(server *Server) error {
		server.taskService = taskService
		return nil
	}
}

func WithMessageExecutor(messageExecutor kafka.KafkaMessageExecutor) func(server *Server) error {
	return func(server *Server) error {
		server.messageExecutor = messageExecutor
		return nil
	}
}

func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
