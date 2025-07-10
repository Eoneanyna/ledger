package route

import (
	"github.com/gin-gonic/gin"
	"ledger/handler"
)

var userPrefix = "/user"

func RegisterUserGroup(r *gin.RouterGroup) {
	r.GET(legerPrefix+"/:id", handler.GetUser)
	r.POST(legerPrefix+"/login", handler.LoginUser)
	r.POST(legerPrefix+"/register", handler.RegisterUser)
}
