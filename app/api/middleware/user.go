package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Dionizio8/go-task/entity"
	"github.com/Dionizio8/go-task/usecase/user"
	"github.com/gin-gonic/gin"
)

type UserMiddler struct {
	r user.Repository
}

func NewUserMiddler(r user.Repository) *UserMiddler {
	return &UserMiddler{
		r: r,
	}
}

func (u *UserMiddler) RoleUseMiddler(ctx *gin.Context) {
	userId := ctx.GetHeader("userId")
	errorMessage := "unidentified user"
	if userId == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New(errorMessage))
		return
	}

	user, err := u.r.GetById(userId)
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = u.validRule(user, ctx.Request.Method)
	if err != nil {
		ctx.AbortWithError(http.StatusForbidden, err)
		return
	}

	ctx.Params = append(ctx.Params, gin.Param{Key: "role", Value: user.Role})
}

// TODO: Refatorar o valiRule por URI e não por Method
func (u *UserMiddler) validRule(user entity.User, method string) error {
	if user.Role == entity.GetUserRoleManager() && method != http.MethodGet {
		return fmt.Errorf("the user %v does not have permission to perform this action", user.Id)
	}
	return nil
}
