package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Dionizio8/go-task/app/api/presenter"
	"github.com/Dionizio8/go-task/entity"
	"github.com/Dionizio8/go-task/infra/kafka"
	"github.com/Dionizio8/go-task/usecase/task"
	"github.com/Dionizio8/go-task/usecase/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskHandler struct {
	service     task.Service
	userService user.Service
	msg         kafka.KafkaMessageExecutor
}

func NewTaskHandler(service task.Service, userService user.Service, msg kafka.KafkaMessageExecutor) *TaskHandler {
	return &TaskHandler{service: service, userService: userService, msg: msg}
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
		return
	}

	userId := uuid.MustParse(ctx.GetHeader("userId"))

	managerUserId, err := h.userService.FindUserById(input.ManagerUserId.String())
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, errorMessage)
		return
	}

	if managerUserId.Role != entity.GetUserRoleManager() {
		log.Println("managerUserId not mnager")
		ctx.AbortWithError(http.StatusForbidden, errorMessage)
		return
	}

	taskId, err := h.service.CreateTask(input.Title, input.Description, userId, input.ManagerUserId, entity.GetTaskInitialState())
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, errorMessage)
		return
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
		return
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

func (h *TaskHandler) EditStatus(ctx *gin.Context) {
	userId := ctx.GetHeader("userId")
	taskId := ctx.Param("taskId")
	if taskId == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("taskId not found"))
		return
	}

	//TODO: Criar responsta quando usuário não é o mesmo criador da task
	errorMessage := errors.New("error find task")
	task, err := h.service.EditTaskStatus(taskId, userId, entity.GetTaskFinalState())
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, errorMessage)
		return
	}

	updateAt := time.Now()

	msg := kafka.TaskMessage{
		UserId:    task.UserId.String(),
		ManagerId: task.ManagerUserId.String(),
		TaskId:    task.Id.String(),
		Status:    task.Status,
		Date:      updateAt,
	}

	msgString, _ := json.Marshal(msg)

	if task.Status == entity.GetTaskFinalState() {
		go h.msg.Push(ctx, uuid.New().String(), string(msgString))
	}

	response := &presenter.Task{
		Id:          task.Id.String(),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}

	ctx.JSON(http.StatusCreated, response)

}

func (h *TaskHandler) getRoleList(role string, userId string) ([]entity.Task, error) {
	if role == entity.GetUserRoleManager() {
		return h.service.List()
	}
	return h.service.FindTaskByUserId(userId)
}
