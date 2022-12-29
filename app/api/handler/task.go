package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Dionizio8/go-task/app/api/presenter"
	"github.com/Dionizio8/go-task/usecase/task"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const InitialState = "IN_PROGRESS"

type TaskHandler struct {
	service task.Service
}

func NewTaskHandler(service task.Service) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) Create(ctx *gin.Context) {
	errorMessage := errors.New("Error adding task")

	var input struct {
		Title         string    `json:"title"`
		Description   string    `json:"description"`
		UserId        uuid.UUID `json:"userId"`
		ManagerUserId uuid.UUID `json:"managerUserId"`
	}

	err := json.NewDecoder(ctx.Request.Body).Decode(&input)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, errorMessage)
	}

	err = h.service.CreateTask(input.Title, input.Description, input.UserId, input.ManagerUserId, InitialState)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, errorMessage)
	}

	response := &presenter.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      InitialState,
	}

	ctx.JSON(http.StatusCreated, response)
}
