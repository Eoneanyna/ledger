package route

import (
	"github.com/gin-gonic/gin"
	"ledger/auth"
)

var PublicPrefix = "/api/v1"

func RegisterRoute(r *gin.Engine) {
	rootRoute := r.Group(PublicPrefix, auth.AuthMiddleware())
	RegisterLegerGroup(rootRoute)
	RegisterUserGroup(rootRoute)
}
