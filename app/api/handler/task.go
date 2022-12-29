package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Dionizio8/go-task/app/api/presenter"
	"github.com/Dionizio8/go-task/entity"
	"github.com/Dionizio8/go-task/usecase/task"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskHandler struct {
	service task.Service
}

func NewTaskHandler(service task.Service) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) Create(ctx *gin.Context) {
	errorMessage := errors.New("error adding task")

	var input struct {
		Title         string    `json:"title"`
		Description   string    `json:"description"`
		ManagerUserId uuid.UUID `json:"managerUserId"`
	}

	err := json.NewDecoder(ctx.Request.Body).Decode(&input)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, errorMessage)
	}

	userId := uuid.MustParse(ctx.GetHeader("userId"))

	taskId, err := h.service.CreateTask(input.Title, input.Description, userId, input.ManagerUserId, entity.GetTaskInitialState())
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, errorMessage)
	}

	response := &presenter.Task{
		Id:          taskId.String(),
		Title:       input.Title,
		Description: input.Description,
		Status:      entity.GetTaskInitialState(),
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *TaskHandler) List(ctx *gin.Context) {
	userId := ctx.GetHeader("userId")
	errorMessage := errors.New("error returning task list")

	tasks, err := h.getRoleList(ctx.Param("role"), userId)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, errorMessage)
	}

	var response []*presenter.Task
	for _, task := range tasks {
		response = append(response, &presenter.Task{
			Id:          task.Id.String(),
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
		})
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *TaskHandler) getRoleList(role string, userId string) ([]entity.Task, error) {
	if role == entity.GetUserRoleManager() {
		return h.service.List()
	}
	return h.service.FindTaskByUserId(userId)
}
