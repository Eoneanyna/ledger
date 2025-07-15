package route

import (
	"github.com/gin-gonic/gin"
	"ledger/route/handler"
)

var userPrefix = "/user"

func RegisterUserGroup(r *gin.RouterGroup) {
	userR := r.Group(userPrefix)
	userR.POST("/login", handler.LoginUser)
	userR.POST("/register", handler.RegisterUser)
}
