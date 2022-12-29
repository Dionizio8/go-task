package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RoleUseMiddler(c *gin.Context) {
	fmt.Println("Middler roles user ! ...")
}
