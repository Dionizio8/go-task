package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Dionizio8/go-task/app/api/middleware"
	"github.com/Dionizio8/go-task/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/Dionizio8/go-task/usecase/task"
	mock_task "github.com/Dionizio8/go-task/usecase/task/mock"
	"github.com/Dionizio8/go-task/usecase/user"
	mock_user "github.com/Dionizio8/go-task/usecase/user/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func Test_CreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userTec := entity.NewUser("Gabriel", entity.GetUserRoleTechnician())
	userMgr := entity.NewUser("Dionizio", entity.GetUserRoleManager())

	newTask := entity.NewTask("Title Task", "Desc Task", userTec.Id, userMgr.Id, entity.GetTaskInitialState())

	controller := gomock.NewController(t)
	defer controller.Finish()

	userRepo := mock_user.NewMockRepository(controller)
	userRepo.EXPECT().GetById(userTec.Id.String()).Return(userTec, nil)
	userRepo.EXPECT().GetById(userMgr.Id.String()).Return(userMgr, nil)
	userService := *user.NewService(userRepo)

	taskRepo := mock_task.NewMockRepository(controller)
	taskRepo.EXPECT().Create(gomock.Any()).Return(newTask.Id, nil)
	taskService := *task.NewService(taskRepo)

	middleware := middleware.NewUserMiddler(userRepo)

	rr := httptest.NewRecorder()

	payload := fmt.Sprintf(`
	{
		"title": "Tarefa 5",
		"description": "desc da tarefa 5",
		"managerUserId": "%v"
	}
	`, userMgr.Id.String())

	server := &Server{
		userService: userService,
		taskService: taskService,
		userMiddler: *middleware,
	}

	router := gin.Default()
	server.router(router)

	request, err := http.NewRequest(http.MethodPost, "/task", strings.NewReader(payload))
	request.Header.Set("userId", userTec.Id.String())
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)

	assert.Equal(t, 201, rr.Code)
}

func Test_EditTaskStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userTec := entity.NewUser("Gabriel", entity.GetUserRoleTechnician())
	newTask := entity.NewTask("Title Task", "Desc Task", userTec.Id, uuid.New(), entity.GetTaskInitialState())

	controller := gomock.NewController(t)
	defer controller.Finish()

	userRepo := mock_user.NewMockRepository(controller)
	userRepo.EXPECT().GetById(userTec.Id.String()).Return(userTec, nil)
	userService := *user.NewService(userRepo)

	taskRepo := mock_task.NewMockRepository(controller)
	taskRepo.EXPECT().UpdateStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(newTask, nil)
	taskService := *task.NewService(taskRepo)

	middleware := middleware.NewUserMiddler(userRepo)

	rr := httptest.NewRecorder()

	server := &Server{
		userService: userService,
		taskService: taskService,
		userMiddler: *middleware,
	}

	router := gin.Default()
	server.router(router)

	path := fmt.Sprintf("/task/conclude/%v", newTask.Id.String())
	request, err := http.NewRequest(http.MethodPatch, path, nil)
	request.Header.Set("userId", userTec.Id.String())
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)

	assert.Equal(t, 201, rr.Code)
}

func Test_GetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userMgr := entity.NewUser("Dionizio", entity.GetUserRoleManager())
	newTask := entity.NewTask("Title Task", "Desc Task", uuid.New(), userMgr.Id, entity.GetTaskInitialState())
	tasks := []*entity.Task{newTask}

	controller := gomock.NewController(t)
	defer controller.Finish()

	userRepo := mock_user.NewMockRepository(controller)
	userRepo.EXPECT().GetById(userMgr.Id.String()).Return(userMgr, nil)
	userService := *user.NewService(userRepo)

	taskRepo := mock_task.NewMockRepository(controller)

	taskRepo.EXPECT().GetAll().Return(tasks, nil)
	taskService := *task.NewService(taskRepo)

	middleware := middleware.NewUserMiddler(userRepo)

	rr := httptest.NewRecorder()

	server := &Server{
		userService: userService,
		taskService: taskService,
		userMiddler: *middleware,
	}

	router := gin.Default()
	server.router(router)

	request, err := http.NewRequest(http.MethodGet, "/task", nil)
	request.Header.Set("userId", userMgr.Id.String())
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)

	assert.Equal(t, 200, rr.Code)
}

func Test_GetByUserId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userTec := entity.NewUser("Gabriel", entity.GetUserRoleTechnician())
	newTask := entity.NewTask("Title Task", "Desc Task", userTec.Id, uuid.New(), entity.GetTaskInitialState())
	tasks := []*entity.Task{newTask}

	controller := gomock.NewController(t)
	defer controller.Finish()

	userRepo := mock_user.NewMockRepository(controller)
	userRepo.EXPECT().GetById(userTec.Id.String()).Return(userTec, nil)
	userService := *user.NewService(userRepo)

	taskRepo := mock_task.NewMockRepository(controller)
	taskRepo.EXPECT().GetByUserId(gomock.Any()).Return(tasks, nil)
	taskService := *task.NewService(taskRepo)

	middleware := middleware.NewUserMiddler(userRepo)

	rr := httptest.NewRecorder()

	server := &Server{
		userService: userService,
		taskService: taskService,
		userMiddler: *middleware,
	}

	router := gin.Default()
	server.router(router)

	request, err := http.NewRequest(http.MethodGet, "/task", nil)
	request.Header.Set("userId", userTec.Id.String())
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)

	assert.Equal(t, 200, rr.Code)
}
