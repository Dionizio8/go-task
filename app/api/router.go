package api

import (
	"net/http"

	"github.com/Dionizio8/go-task/app/api/handler"
	"github.com/gin-gonic/gin"
)

func (s *Server) router(r gin.IRouter) {
	r.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "OK"}) })

	taskHandler := handler.NewTaskHandler(s.taskService)
	r.POST("/task", taskHandler.Create)
}
