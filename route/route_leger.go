package route

import (
	"github.com/gin-gonic/gin"
	"ledger/auth"
)

var legerPrefix = "/leger"

func RegisterLegerGroup(r *gin.RouterGroup) {
	r.Group(legerPrefix, auth.AuthMiddleware())
	//r.GET(legerPrefix+"/:id", GetLege)
	//r.POST(legerPrefix, CreateLege)
}
